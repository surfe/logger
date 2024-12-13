package key

const (
	Email            = "email"
	CompanyKey       = "company_key"
	Latency          = "latency"
	Method           = "method"
	URI              = "uri"
	Path             = "path"
	External         = "external"
	Status           = "status"
	UserAgent        = "user_agent"
	APIVersion       = "api_version"
	ExtensionVersion = "extension_version"
	Payload          = "payload"
	User             = "user"
	CorrelationID    = "correlation_id"
	Tool             = "tool"
	ProductFeature   = "product_feature"
	JobDetails       = "job_details"
)

type ctxKey int

const (
	CtxEmail ctxKey = iota
	CtxCompany
	CtxCorrelationID
	CtxTool
	CtxProductFeature
	CtxJobDetails
)
