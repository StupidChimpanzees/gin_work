package response

type Response struct {
	Code    int         `json:"code" yaml:"code" xml:"code" bson:"code"`
	Data    interface{} `json:"data" yaml:"data" xml:"data" bson:"data"`
	Message string      `json:"message" yaml:"message" xml:"message" bson:"message"`
}

func (r *Response) Success(args ...interface{}) *Response {
	response := r.response(200, "Success", nil)
	if args != nil {
		if len(args) > 2 {
			response.Data = args[0]
			response.Code = args[1].(int)
			response.Message = args[2].(string)
		} else if len(args) > 1 {
			response.Data = args[0]
			response.Code = args[1].(int)
		} else {
			response.Data = args[0]
		}
	}
	return response
}

func (r *Response) Fail(args ...interface{}) *Response {
	response := r.response(500, "Fail", nil)
	if args != nil {
		if len(args) > 2 {
			response.Message = args[0].(string)
			response.Code = args[1].(int)
			response.Data = args[2]
		} else if len(args) > 1 {
			response.Message = args[0].(string)
			response.Code = args[1].(int)
		} else {
			response.Message = args[0].(string)
		}
	}
	return response
}

func (*Response) response(code int, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Data:    data,
		Message: message,
	}
}
