package shared

type IOR struct {
	Host      string
	Port      int
	Id        int
	TypeName  string
	LeaseName string
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

const ServerPort int = 1999
const ClientServerPort int = 1998
const NamingServicePort int = 1313
const CalculatorPort int = 1314
const DefaultLeasingTimeSeconds float64 = 10
