package myzap

type Category string

func (c Category) String() string {
	return string(c)
}

type SubCategory string

func (c SubCategory) String() string {
	return string(c)
}

type ExtraKey string

const (
	CategoryGeneral Category = "general"
	IO              Category = "IO"
	Internal        Category = "Internal"
	HTTPServer      Category = "HTTPServer"
	RabbitMQ        Category = "RabbitMQ"
	EventDriven     Category = "Event-driven"
	GRPCServer      Category = "GRPCServer"
	Postgres        Category = "Postgres"
	MongoDB         Category = "MongoDB"
	Redis           Category = "Redis"
	Validation      Category = "Validation"
	RequestResponse Category = "RequestResponse"
	Prometheus      Category = "Prometheus"
	Service         Category = "Service"
	Router          Category = "Router"
	Controller      Category = "Controller"
	Usecase         Category = "Usecase"
	Model           Category = "Model"
	Trace           Category = "Trace"
	Collector       Category = "Collector"
)

const (
	// General
	Startup         SubCategory = "Startup"
	ExternalService SubCategory = "ExternalService"
	InternalService SubCategory = "InternalService"

	// DB
	Migration SubCategory = "Migration"
	Select    SubCategory = "Select"
	Rollback  SubCategory = "Rollback"
	Update    SubCategory = "Update"
	Delete    SubCategory = "Delete"
	Insert    SubCategory = "Insert"

	// Internal
	Api                 SubCategory = "Api"
	HashPassword        SubCategory = "HashPassword"
	DefaultRoleNotFound SubCategory = "DefaultRoleNotFound"

	// Validation
	MobileValidation   SubCategory = "MobileValidation"
	PasswordValidation SubCategory = "PasswordValidation"

	// IO
	RemoveFile SubCategory = "RemoveFile"
)

const (
	AppName       ExtraKey = "AppName"
	Configuration ExtraKey = "config"
	LoggerName    ExtraKey = "Logger"
	ClientIp      ExtraKey = "ClientIp"
	HostIp        ExtraKey = "HostIp"
	Method        ExtraKey = "Method"
	StatusCode    ExtraKey = "StatusCode"
	BodySize      ExtraKey = "BodySize"
	Path          ExtraKey = "Path"
	Latency       ExtraKey = "Latency"
	RequestBody   ExtraKey = "RequestBody"
	ResponseBody  ExtraKey = "ResponseBody"
	ErrorMessage  ExtraKey = "ErrorMessage"
	Status        ExtraKey = "Status"
	Data          ExtraKey = "Data"
)
