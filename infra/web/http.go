package web

type Headers struct {
	ContentType string
}

type HttpResponse struct {
	StatusCode int
	Body       []byte
	Headers    Headers
}

type HttpRequest struct {
	Body []byte
}
