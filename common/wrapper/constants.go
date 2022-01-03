package wrapper

type TypeError uint16

const (
	ErrDatabase    TypeError = iota + 10 // 10
	ErrFromUseCase                       // 11
	ErrValidation                        // 12
	ErrUnknown                           // 13
)
