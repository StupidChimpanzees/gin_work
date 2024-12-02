package response

type Response struct {
	Code    int         `json:"code" yaml:"code" xml:"code" bson:"code"`
	Data    interface{} `json:"data" yaml:"data" xml:"data" bson:"data"`
	Message string      `json:"message" yaml:"message" xml:"message" bson:"message"`
}

func (*Response) Success(args ...interface{}) *Response {
	response := &Response{
		Code:    200,
		Data:    nil,
		Message: "Success",
	}
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

func (*Response) Fail(args ...interface{}) *Response {
	response := &Response{
		Code:    500,
		Data:    nil,
		Message: "Fail",
	}
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
