package wapper

type ResponseError struct {
	Code       int
	Message    string
	StatusCode int
	ERR        error
}

func (r *ResponseError) Error() string {
	if r.ERR != nil {
		return r.ERR.Error()
	}
	return ""
}
