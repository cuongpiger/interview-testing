package server

type (
	ResponseBuilder interface {
		SetCode(string) ResponseBuilder
		SetMessage(string) ResponseBuilder
		SetData(interface{}) ResponseBuilder
		GetResponse() Response
	}

	Response struct {
		Code    string      `json:"code"`           // Response code of the request
		Message string      `json:"message"`        // Response message of the request
		Data    interface{} `json:"data,omitempty"` // Response data
	}
)

func NewResponse() ResponseBuilder {
	return new(Response)
}

func (s *Response) SetCode(code string) ResponseBuilder {
	s.Code = code
	return s
}

func (s *Response) SetMessage(message string) ResponseBuilder {
	s.Message = message
	return s
}

func (s *Response) SetData(data interface{}) ResponseBuilder {
	s.Data = &data
	return s
}

func (s *Response) GetResponse() Response {
	return *s
}
