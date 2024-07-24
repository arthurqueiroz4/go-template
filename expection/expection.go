package expection

type ErrorBase struct {
	Body    any
	Message string
	Status  int
}

func (e *ErrorBase) Error() string {
	return e.Message
}

type (
	ErrBadRequest interface {
		Error() string
	}
	ErrNotFound interface {
		Error() string
	}
)

func NewErrBadRequest(body any, message string) ErrBadRequest {
	return &ErrorBase{Body: body, Message: message, Status: 400}
}

func NewErrNotFound(body any, message string) ErrNotFound {
	return &ErrorBase{Body: body, Message: message, Status: 404}
}

type NewErrInternalServer struct {
	ErrorBase
}
