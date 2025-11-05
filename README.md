# nmod sample modules

This repository demonstrates using two independent Go modules: a simple IP validation library and an HTTP API that consumes it.

## Getting started

The library and API live in separate directories and have individual `go.mod` files:

- `iprange`: exposes `iprange.IsValid` to check if a string looks like a valid IPv4 or IPv6 address.
- `api`: provides an HTTP server that calls the library.

## Multi-module versioning

This repository contains **two independent modules** that must be tagged separately:

- **iprange module**: `github.com/WebBeds/nmod/iprange`
- **api module**: `github.com/WebBeds/nmod/api`

### Tagging strategy

Each module uses **subdirectory-prefixed tags**:

```bash
# Tag the IP validation library
git tag iprange/v1.0.0
git push origin iprange/v1.0.0

# Tag the API separately
git tag api/v1.0.0
git push origin api/v1.0.0
```

### Why subdirectory prefixes?

Go's module system requires tags to include the subdirectory path when modules aren't at the repository root. This allows:

- Independent versioning of each module
- Proper resolution by `go get`
- Clear semantic versioning per component

### Consuming tagged versions

Other projects can depend on specific versions:

```bash
# Use a specific version of iprange
go get github.com/WebBeds/nmod/iprange@v1.0.0

# Use a specific version of api
go get github.com/WebBeds/nmod/api@v1.0.0
```

## Run the API

```bash
PORT=8080 go run api/main.go
```

If you leave `PORT` unset, the server defaults to 8080.

## Sample requests

With the server running:

- Valid IP example: `curl -i "http://localhost:8080/validate?ip=192.168.0.11"`
- Invalid IP example: `curl -i "http://localhost:8080/validate?ip=192.168.0.256"`
- Bad input example: `curl -i "http://localhost:8080/validate"`

All responses are JSON with the shape:

```json
{ "ip": "<value you sent>", "valid": <true|false> }
```
