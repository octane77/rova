package dtos

type CallServiceApiDto struct {
	Route  string
	Token  string
	Body   interface{}
	Method string
	Path   string // this is the route plus any other query parameter
}
