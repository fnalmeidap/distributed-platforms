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

func NewIOR(h string, p int) IOR {
	return IOR{Host: h, Port: p, Id: 0, TypeName: ""}
}

const MaxConnectionAttempts int = 30
const LocalHost string = "localhost"
const DefaultPort int = 1999
const DefaultLeasingTimeSeconds float64 = 10
const NamingPort int = 1313
const CalculadoraPort int = 1314
