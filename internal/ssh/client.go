package ssh

import (
	"fmt"
	"net"
	"os"

	"github.com/cristianperen/servercheck/internal/models"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/crypto/ssh/knownhosts"
)

func getHostKeyCallback() (ssh.HostKeyCallback, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	knownHostsPath := home + "/.ssh/known_hosts"

	callback, err := knownhosts.New(knownHostsPath)
	if err != nil {
		return nil, err
	}

	return callback, nil
}

func getAgent() (ssh.AuthMethod, error) {
	socket := os.Getenv("SSH_AUTH_SOCK")

	if socket == "" {
		return nil, fmt.Errorf("SSH_AUTH_SOCK no está configurado")
	}

	conn, err := net.Dial("unix", socket)
	if err != nil {
		return nil, fmt.Errorf("No se pudo conectar al SSH agent: %w", err)
	}

	agentClient := agent.NewClient(conn)

	return ssh.PublicKeysCallback(agentClient.Signers), nil
}

func Connect(server models.Server) (*ssh.Client, error) {
	authMethod, err := getAgent()
	if err != nil {
		return nil, err
	}

	hostKeyCallback, err := getHostKeyCallback()
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: server.User,
		Auth: []ssh.AuthMethod{
			authMethod,
		},
		HostKeyCallback: hostKeyCallback,
	}

	address := net.JoinHostPort(
		server.Host,
		fmt.Sprintf("%d", server.Port),
	)

	client, err := ssh.Dial("tcp", address, config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func RunCommand(client *ssh.Client, command string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", err
	}

	return string(output), nil
}
