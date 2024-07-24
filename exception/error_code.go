package exception

type ErrorCode int

const (
	NotFound ErrorCode = iota
	InternalServer
)
