package defs

type Err struct {
	Error string `json:"name"`
	ErrorCode string `json:"error_code"`

}

type ErrorResponse struct {
	HttpSC int
	Error Err
}

var (
	ErrRequestBodyParseFailed = ErrorResponse{HttpSC:400, Error:Err {Error:"Request body parse failed", ErrorCode:"001" }}
	ErrorNotAuthUser = ErrorResponse{HttpSC:401, Error:Err{Error:"User authentication failed", ErrorCode:"002"}}
)
