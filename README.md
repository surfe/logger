# Logger

Logging library wrapper with `Echo`.

## Install

```
go get -u github.com/surfe/logger
```

## Usage

### Initialization

Initiate a Zap logger;
```go
zapLogger, err := zap.Init()
if err != nil {
	log.Panic(err)
}
defer zapLogger.Sync()
```

Use the logger;
```go
l := logger.Use(zapLogger)
```

### Logging

Error with a message and extra fields;
```go
import "github.com/surfe/logger"

...

fields := []any{l.UserKey, "abc@xyz.com"}
logger.Log(ctx).Err(err).With(fields...).Error("Add Contact (SF)")
```

### Echo Middleware

```go
e.Use(l.EchoMiddleware())
```

## Development

You can use `go work` to develop this module:

```bash
go work init .
go work use ../logger
```
