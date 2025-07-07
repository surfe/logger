package key

const (
	Email            = "email"
	CompanyKey       = "company_key"
	Latency          = "latency"
	ProcessingTime   = "processing_time"
	Method           = "method"
	URI              = "uri"
	Path             = "path"
	Route            = "route"
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

	DataDogTraceID     = "trace_id"
	DataDogSpanID      = "span_id"
	DataDogVersion     = "version"
	DataDogEnvironment = "env"
	DatadogService     = "service"

	HeaderCorrelationID = "correlation-id"
	HeaderEmail         = "email"
	HeaderCompanyKey    = "company-key"
)

type ctxKey int

const (
	CtxEmail ctxKey = iota
	CtxCompany
	CtxCorrelationID
	CtxTool
	CtxProductFeature
	CtxAPIVersion
	CtxJobDetails
	CtxService
)
