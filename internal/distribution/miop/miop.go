package miop

import (
	"distributed-platforms/internal/shared"
)

type Packet struct {
	Hdr Header
	Bd  Body
}

type Header struct {
	Magic       string // remove
	Version     string // remove
	ByteOrder   bool
	MessageType int
	Size        int
}

type Body struct {
	ReqHeader RequestHeader
	ReqBody   RequestBody
	RepHeader ReplyHeader
	RepBody   ReplyBody
}

type RequestHeader struct {
	Context          string
	RequestId        int
	ResponseExpected bool
	ObjectKey        int
	Operation        string
}

type RequestBody struct {
	Body []interface{}
}

type ReplyHeader struct {
	Context   string
	RequestId int
	Status    int
}

type ReplyBody struct {
	OperationResult []interface{}
}

func CreateRequestMIOP(op string, p []interface{}) Packet {
	r := Packet{
		Hdr: Header{},
		Bd: Body{
			ReqHeader: RequestHeader{
				Operation: op,
			},
			ReqBody: RequestBody{
				Body: p,
			},
		},
	}

	return r
}

func CreateReplyMIOP(params []interface{}) Packet {
	r := Packet{
		Hdr: Header{},
		Bd: Body{
			RepBody: ReplyBody{
				OperationResult: params,
			},
		},
	}

	return r
}

func ExtractRequest(m Packet) shared.Request {
	i := shared.Request{}

	i.Operation = m.Bd.ReqHeader.Operation
	i.Params = m.Bd.ReqBody.Body

	return i
}

func ExtractReply(m Packet) shared.Reply {
	var r shared.Reply

	r.Result = m.Bd.RepBody.OperationResult

	return r
}
