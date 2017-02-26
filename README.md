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

## Hand Wavy Bench

Local hand wavy benchmarks for rough estimates on transfer speed

size	|	time taken
----	|	----------
1mb	|	15.590542ms
8mb	|	75.184788ms
64mb	|	516.236417ms
128mb	|	1.141906576s
1024mb	|	9.794891634s

Using connection pooling and caching selector

size    |       time taken
----    |       ----------
1mb     |       13.521179ms
8mb     |       53.160487ms
64mb    |       415.388025ms
128mb   |       889.409332ms
512mb   |       4.177052391s
1024mb  |       8.347038098s
