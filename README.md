# gunner

A simple, flexible configuration loader for Go applications that supports multiple file formats and environment variable overrides.

## Features

- Load configuration from YAML, JSON, and .env files
- Environment variable overrides with custom prefix
- Automatic config directory resolution (follows OS conventions)
- Merges multiple configuration files
- Type-safe configuration loading

## Installation

```bash
go get github.com/abs3ntdev/gunner
```

## Usage

### Basic Example

Define your configuration struct:

```go
package main

import (
    "log"
    "github.com/abs3ntdev/gunner"
)

type Config struct {
    Database struct {
        Host     string `default:"localhost"`
        Port     int    `default:"5432"`
        User     string `env:"DB_USER"`
        Password string `env:"DB_PASSWORD"`
    }
    Server struct {
        Port int    `default:"8080"`
        Host string `default:"0.0.0.0"`
    }
}

func main() {
    var cfg Config

    if err := gunner.LoadApp(&cfg, "myapp"); err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Use your configuration
    log.Printf("Server will run on %s:%d", cfg.Server.Host, cfg.Server.Port)
    log.Printf("Database: %s:%d", cfg.Database.Host, cfg.Database.Port)
}
```

### Configuration Files

Gunner looks for configuration files in the following locations:
- `~/.config/myapp/myapp.yml`
- `~/.config/myapp/myapp.yaml`
- `~/.config/myapp/myapp.json`
- `~/.config/myapp/.env`

All files are merged together if multiple exist.

### Example Configuration Files

#### YAML (myapp.yaml)

```yaml
database:
  host: localhost
  port: 5432
  user: dbuser

server:
  port: 8080
  host: 0.0.0.0
```

#### JSON (myapp.json)

```json
{
  "database": {
    "host": "localhost",
    "port": 5432,
    "user": "dbuser"
  },
  "server": {
    "port": 8080,
    "host": "0.0.0.0"
  }
}
```

#### .env

```env
MYAPP_DATABASE_HOST=localhost
MYAPP_DATABASE_PORT=5432
MYAPP_DATABASE_USER=dbuser
MYAPP_SERVER_PORT=8080
MYAPP_SERVER_HOST=0.0.0.0
```

### Environment Variables

Environment variables are automatically loaded with the app name as prefix (uppercase):

- `MYAPP_DATABASE_HOST` overrides `database.host`
- `MYAPP_DATABASE_PORT` overrides `database.port`
- etc.

## Supported Struct Tags

Gunner uses [aconfig](https://github.com/cristalhq/aconfig) under the hood, supporting tags like:

- `default:"value"` - default value
- `env:"VAR_NAME"` - specific environment variable name
- `required:"true"` - mark field as required

## License

This project follows the license of its dependencies.

## Credits

Built using:
- [cristalhq/aconfig](https://github.com/cristalhq/aconfig) - configuration loader
- [sigs.k8s.io/yaml](https://github.com/kubernetes-sigs/yaml) - YAML parsing
