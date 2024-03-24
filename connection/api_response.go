package connection

type ApiResponse interface {
	Retrieve() interface{}
}
