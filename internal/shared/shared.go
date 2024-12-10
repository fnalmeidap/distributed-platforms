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

var MaxConnectionAttempts int = 30
var LocalHost string = "localhost"
var DefaultPort int = 1999
var DefaultLeasingTimeSeconds float64 = 20
