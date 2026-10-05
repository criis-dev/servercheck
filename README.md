# ServerCheck

ServerCheck es una pequeña herramienta de línea de comandos para realizar comprobaciones de salud en servidores SSH remotos. Se conecta a cada servidor utilizando el agente SSH local, ejecuta comprobaciones básicas (uso de disco e inodos) y muestra una tabla resumida con el estado general y las métricas.

Versión: 0.1.0

---

## Características

- Se conecta a servidores remotos vía SSH (usa el agente SSH / autenticación por clave pública)
- Ejecuta comprobaciones de uso de disco e inodos (mediante `df -hl` y `df -il`)
- Agrega resultados y reporta el estado general (OK / Warning / Critical / DOWN)
- Configuración simple en YAML para listar servidores

---

## Requisitos

- Go (el módulo requiere `go 1.27.1`, tal como aparece en `go.mod`)
- Un agente SSH con las claves adecuadas cargadas (`SSH_AUTH_SOCK` configurado)
- Las claves host de los servidores remotos presentes en `~/.ssh/known_hosts` (el programa usa `known_hosts` para verificar la clave del host)

---

## Instalación

Hay varias formas de compilar e instalar ServerCheck.

Desde el código fuente (flujo de trabajo para desarrollo):

```bash
# Compilar el binario en la raíz del repositorio
make build

# Ejecutarlo directamente
make run --silent
# o
./servercheck <comando>
```

Crear una versión de lanzamiento para Linux amd64:

```bash
make release
# Esto creará dist/servercheck-linux-amd64
```

Instalar de forma global (el objetivo del Makefile usa sudo):

```bash
make release
make install
```

El objetivo `install` hará lo siguiente:
- crear `/etc/servercheck`
- copiar `dist/servercheck-linux-amd64` a `/usr/local/bin/servercheck`
- copiar `configs/servers.yaml` a `/etc/servercheck/servers.yaml`

> Nota: `make install` usa `sudo` y sobrescribe rutas del sistema. Revisa los objetivos antes de ejecutarlo.

---

## Configuración

ServerCheck lee un archivo YAML con la lista de servidores a comprobar. La ruta de configuración predeterminada es `/etc/servercheck/servers.yaml`, pero se puede especificar otra ruta con la opción `--config`.

Ejemplo de configuración (`configs/servers.yaml` incluido en el repositorio):

```yaml
servers:
  - name: servir-server2
    host: servir-server2.servir-vpn
    port: 22
    user: administrator

  - name: ubuntu-db-server
    host: ubuntu-db-server.servir-vpn
    port: 22
    user: administrator

  - name: ubuntu-cloud-server
    host: ubuntu-cloud-server.servir-vpn
    port: 22
    user: administrator
```

Campos:
- `name`: nombre amigable que se usa en la salida
- `host`: nombre de host o dirección IP
- `port`: puerto SSH (entero)
- `user`: usuario SSH que se usará para la conexión

---

## Uso

Comandos:

- `servercheck version` — muestra la versión de la herramienta
- `servercheck check` — ejecuta comprobaciones para todos los servidores de la configuración
- `servercheck check --server <name>` — ejecuta comprobaciones solo para el servidor indicado
- `servercheck check --verbose` — muestra resultados detallados de cada comprobación
- `servercheck check --config <path>` — especifica un archivo de configuración alternativo

Ejemplos:

```bash
# Revisar todos los servidores (usa la configuración predeterminada /etc/servercheck/servers.yaml)
servercheck check

# Revisar un solo servidor por nombre
servercheck check --server ubuntu-db-server

# Ejecutar con salida detallada
servercheck check --verbose

# Usar un archivo de configuración personalizado ubicado en el repositorio
servercheck check --config configs/servers.yaml
```

---

## Cómo funciona

- La CLI carga la lista de servidores desde el archivo YAML de configuración.
- Para cada servidor intenta conectarse usando el agente SSH local (`SSH_AUTH_SOCK`) y las claves del host almacenadas en `~/.ssh/known_hosts`.
- Si la sesión SSH se establece, ServerCheck ejecuta `df -hl` para recopilar el uso del disco y `df -il` para recopilar el uso de inodos.
- La herramienta analiza las salidas, evalúa el estado por sistema de archivos según umbrales y agrega un estado para cada servidor.
- La salida se imprime en una tabla legible; el modo detallado muestra los resultados completos y analizados.

---

## Desarrollo

Ejecutar pruebas:

```bash
make test
# o
go test ./...
```

Formateo:

```bash
make fmt
```

Compilar el binario para desarrollo:

```bash
make build
```

Estructura del proyecto (nivel alto):

- `cmd/servercheck` - punto de entrada principal de la CLI
- `internal/config` - carga de la configuración
- `internal/ssh` - cliente SSH (autenticación con agente y verificación de clave del host)
- `internal/checks` - lógica para ejecutar comprobaciones en un servidor
- `internal/output` - utilidades para formatear e imprimir resultados
- `configs/servers.yaml` - configuración de ejemplo

---

## Solución de problemas

- "SSH_AUTH_SOCK no está configurado" — asegúrate de que un agente SSH esté en ejecución y de que tu clave esté cargada (por ejemplo, `eval "$(ssh-agent -s)"` y `ssh-add ~/.ssh/id_rsa`).
- Errores de verificación de clave del host — agrega la clave del host remoto a `~/.ssh/known_hosts` conectándote una vez con `ssh user@host`.
- Permisos denegados / fallos de autenticación — asegúrate de que la clave SSH cargada en el agente tenga acceso a la cuenta remota.
- Si no se puede leer el archivo de configuración, verifica la ruta y los permisos (predeterminado: `/etc/servercheck/servers.yaml`).

---

## Contribuciones

Las contribuciones son bienvenidas. Abre issues o PRs en el repositorio.

---

## Licencia

No se incluye un archivo de licencia en el repositorio. Agrega un archivo LICENSE si deseas publicar este proyecto bajo una licencia de código abierto.

---

Mantenedores: consulta la ruta del módulo en `go.mod` (github.com/cristianperen/servercheck)
