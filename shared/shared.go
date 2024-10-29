package shared

type IOR struct {
	Host     string
	Port     int
	Id       int
	TypeName string
}

type Request struct {
	Operation string
	Params    []interface{}
}

type Reply struct {
	Result []interface{}
}
