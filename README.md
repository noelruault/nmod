# nmod sample modules

This repository demonstrates using two independent Go modules: a simple IP validation library and an HTTP API that consumes it.

## Getting started

The library and API live in separate directories and have individual `go.mod` files:

- `iprange`: exposes `iprange.IsValid` to check if a string looks like a valid IPv4 or IPv6 address.
- `api`: provides an HTTP server that calls the library.

## Multi-module versioning

This repository contains **two independent modules** that must be tagged separately:

- **iprange module**: `github.com/noelruault/nmod/iprange`
- **api module**: `github.com/noelruault/nmod/api`

### Publishing to GitHub

Before you can use versioned tags, you must first push this repository to GitHub:

```bash
# 1. Create a new repository on GitHub: https://github.com/noelruault/nmod
# 2. Add the remote and push
git remote add origin git@github.com:noelruault/nmod.git
git branch -M main
git add .
git commit -m "Initial commit: multi-module repository"
git push -u origin main
```

### Tagging strategy

Each module uses **subdirectory-prefixed tags**:

```bash
# Tag the IP validation library
git tag -a iprange/v1.0.0 -m "iprange: release v1.0.0"
git push origin iprange/v1.0.0

# Tag the API separately
git tag -a api/v1.0.0 -m "api: release v1.0.0"
git push origin api/v1.0.0
```

### Why subdirectory prefixes?

Go's module system requires tags to include the subdirectory path when modules aren't at the repository root. This allows:

- Independent versioning of each module
- Proper resolution by `go get`
- Clear semantic versioning per component

### Consuming tagged versions

**Important**: Tags must be pushed to GitHub before they can be consumed. Create a separate test project:

```bash
# Create a test project outside this repository
cd /tmp
mkdir iprange-test && cd iprange-test
go mod init example.com/test

# Fetch specific versions (only works after pushing tags to GitHub)
go get github.com/noelruault/nmod/iprange@iprange/v1.0.0

# Or use the simpler semantic version (Go strips the prefix)
go get github.com/noelruault/nmod/iprange@v1.0.0
```

Verify available versions:

```bash
# List all versions of iprange module (requires public GitHub repo)
go list -m -versions github.com/noelruault/nmod/iprange

# List all versions of api module
go list -m -versions github.com/noelruault/nmod/api
```

### Local development

For local development without pushing tags, the `api` module already uses a `replace` directive:

```go
replace github.com/noelruault/nmod/iprange => ../iprange
```

This allows the API to use the local iprange code. Run `go mod tidy` in the api directory to sync dependencies.

## Run the API

```bash
make run
```

If you leave `PORT` unset, the server defaults to 9000.

## Sample requests

With the server running:

- Valid IP example: `curl -i "http://localhost:9000/validate?ip=192.168.0.11"`
- Invalid IP example: `curl -i "http://localhost:9000/validate?ip=192.168.0.256"`
- Bad input example: `curl -i "http://localhost:9000/validate"`

All responses are JSON with the shape:

```json
{ "ip": "<value you sent>", "valid": <true|false> }
```
