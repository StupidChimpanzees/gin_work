package response

type response struct {
	Code    int         `json:"code" yaml:"code" xml:"code" bson:"code"`
	Data    interface{} `json:"data" yaml:"data" xml:"data" bson:"data"`
	Message string      `json:"message" yaml:"message" xml:"message" bson:"message"`
}

func newResponse(code int, message string, data interface{}) *response {
	return &response{
		Code:    code,
		Data:    data,
		Message: message,
	}
}

func Success(args ...interface{}) (int, *response) {
	response := newResponse(200, "Success", nil)
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
	return 200, response
}

func Fail(args ...interface{}) (int, *response) {
	response := newResponse(500, "Fail", nil)
	if args != nil {
		if len(args) > 2 {
			response.Code = args[0].(int)
			response.Message = args[1].(string)
			response.Data = args[2]
		} else if len(args) > 1 {
			response.Code = args[0].(int)
			response.Message = args[1].(string)
		} else {
			response.Code = args[0].(int)
		}
	}
	return response.Code, response
}
