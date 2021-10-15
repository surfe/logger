# Logger

Logging library for Leadjet backend services.

## Install

```
go get -u github.com/Leadjet/logger
```

## Development

Add above `replace` directive to `go.mod` file pointing to the Logger project location.

```
replace (
	github.com/Leadjet/logger v1.0.0 => ../logger
)
```
