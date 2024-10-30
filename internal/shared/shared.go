package shared

type IOR struct {
	Host     string
	Port     int
	Id       int
	TypeName string
}

type Invocation struct {
	Ior     IOR
	Request Request
}

type Termination struct {
	Rep Reply
}

type Request struct {
	Operation string
	Params    []interface{}
}

type Reply struct {
	Result []interface{}
}

var MaxConnectionAttempts = 30
var LocalHost = "localhost"
var DefaultPort = 1999
