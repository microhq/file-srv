# File Service

The file service is a simple RPC based static file server.

## Prerequisites

We need service discovery

### Zero Dependency

Use multicast DNS locally by passing `--registry=mdns` to the client and server

```
go run main.go --registry=mdns
```

### Consul

```
brew install consul
```

```
consul agent -dev
```

## Usage

### Run Service

```
go run main.go --file_dir=/tmp
```

### Client

A client can be found at [micro/go-file](https://github.com/micro/go-file)

```go
import (
	"github.com/micro/go-micro"
	"github.com/micro/go-file"
)

service := micro.NewService()
service.Init()

client := file.NewClient("go.micro.srv.file", service.Client())
client.Download("remote.file", "local.file")
```

## API

- File.Open
- File.Stat
- File.Read
- File.Close
