package shared

type Request []any
type Response struct {
	StatusCode int
	Result     any
}

type Controller interface {
	Handle(Request) Response
}
