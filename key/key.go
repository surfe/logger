package key

const (
	Email            = "email"
	CompanyKey       = "company_key"
	Latency          = "latency"
	Method           = "method"
	URI              = "uri"
	Status           = "status"
	UserAgent        = "user_agent"
	APIVersion       = "api_version"
	ExtensionVersion = "extension_version"
	Payload          = "payload"
	User             = "user"
	CorrelationID    = "correlation_id"
	Tool             = "tool"
)

type ctxKey int

const (
	CtxEmail ctxKey = iota
	CtxCompany
	CtxCorrelationID
	CtxTool
)
