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

type Request struct {
	Operation string
	Params    []interface{}
}

type Reply struct {
	Result []interface{}
}
