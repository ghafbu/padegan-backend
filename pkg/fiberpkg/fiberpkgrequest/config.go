package fiberpkgrequest

const (
	// and the content-type is `multipart/form-data` or `application/x-www-form-urlencoded`.
	MultipartForm = "multipartForm"
	// Form indicates that the data to validate is in the request body,
	Form = "form"
	// JSON indicates that the data to validate is in the request body,
	// and the content-type is `application/json`.
	JSON = "json"
	// XML indicates that the data to validate is in the request body,
	// and the content-type is `application/xml` or `text/xml`.
	XML = "xml"
	// Query indicates that the data to validate is in the query string.
	Query = "query"
	// Params indicates that the data to validate is in the route parameters.
	Params = "params"
	// ExtractIDFromToken indicates that the data to validate is in the token
	ExtractIDFromToken = "extractIDFromToken"
)
