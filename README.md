# Logger

Logging library wrapper with `Echo` and discard rules support.

## Install

```
go get -u github.com/Leadjet/logger
```

## Usage

Initiate a Zap logger;
```
zapLogger, err := zap.Init()
if err != nil {
	log.Panic(err)
}
defer zapLogger.Sync()
```

Optionally, set DiscardRules;
```
l.DiscardRules = config.Config.Log.DiscardRules
```

Use the logger;
```
l := logger.Use(zapLogger)
```

Error with a message and extra fields;
```
import l "github.com/Leadjet/logger"
...

fields := []interface{}{l.UserKey, x.UserWithCompany}
l.Log().Errorw("Add Contact (SF)", err, fields...)
```

Or, simply add key-value pairs;
```
l.Log().Errorw("Add Contact (SF)", l.CompanyKey, x.CompanyKey, l.EmailKey, x.User.Email)
```

Only add an error (company key won't be sent thus will not be filtered by company!);
```
l.Log().Error("Add Contact (SF)", err)
```

### Echo Middleware

```
e.Use(l.EchoMiddleware())
```

## Development

Add above `replace` directive to `go.mod` file pointing to the Logger project location.

```
replace (
	github.com/Leadjet/logger v1.0.0 => ../logger
)
```
