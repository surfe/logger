# Logger

Logging library for Leadjet backend services.

## Install

```
go get -u github.com/Leadjet/logger
```

## Usage

Initiate a Zap logger;
```
err := zap.RegisterLog()
if err != nil {
    log.Panic(err)
}
```

Error with a message and extra fields;
```
fields := []interface{}{l.UserKey, x.UserWithCompany}
l.Log.Errorw("Add Contact (SF)", err, fields...)
```

Or, simply add key-value pairs;
```
l.Log.Errorw("Add Contact (SF)", l.CompanyKey, x.CompanyKey, l.EmailKey, x.User.Email)
```

Only add an error (company key won't be sent thus will not be filtered by company!);
```
l.Log.Error("Add Contact (SF)", err)
```

### Echo Middleware

```
e.Use(zap.EchoMiddleware(logger.Log))
```

## Development

Add above `replace` directive to `go.mod` file pointing to the Logger project location.

```
replace (
	github.com/Leadjet/logger v1.0.0 => ../logger
)
```
