# Logger

Logging library for Leadjet backend services.

## Install

```
go get -u github.com/Leadjet/logger
```

## Initiate

Initiate a Zap logger;
```
err := zap.RegisterLog()
if err != nil {
    log.Panic(err)
}
```

### Echo Middleware

```
e.Use(zap.EchoMiddleware(logger.Log))
```

### Usage

Error with a message and extra fields;
```
fields := []interface{}{l.UserKey, x.UserWithCompany}
l.Log.Errorw("Add Contact (SF)", err, fields...)
```

Just an error;
```
l.Log.Error("Add Contact (SF)", err)
```

## Development

Add above `replace` directive to `go.mod` file pointing to the Logger project location.

```
replace (
	github.com/Leadjet/logger v1.0.0 => ../logger
)
```
