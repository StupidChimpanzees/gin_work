package response

import "net/http"

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
	response := newResponse(http.StatusOK, "Success", nil)
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
	return response.Code, response
}

func Fail(args ...interface{}) (int, *response) {
	response := newResponse(http.StatusInternalServerError, "Fail", nil)
	if args != nil {
		if len(args) > 2 {
			response.Message = args[0].(string)
			response.Code = args[1].(int)
			response.Data = args[2]
		} else if len(args) > 1 {
			response.Message = args[0].(string)
			response.Code = args[1].(int)
		} else {
			response.Code = args[0].(int)
		}
	}
	return response.Code, response
}

func RequestFail(args ...interface{}) (int, *response) {
	response := newResponse(http.StatusBadRequest, "Fail", nil)
	if args != nil {
		if len(args) > 2 {
			response.Message = args[0].(string)
			response.Code = args[1].(int)
			response.Data = args[2]
		} else if len(args) > 1 {
			response.Message = args[0].(string)
			response.Code = args[1].(int)
		} else {
			response.Code = args[0].(int)
		}
	}
	return response.Code, response
}
