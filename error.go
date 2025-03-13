package gokit

import "fmt"

type ErrCommon struct {
	Msg    *string
	Code   *string
	Op     *string
	Params *interface{}
}

func (e *ErrCommon) Error() string {
	return fmt.Sprintf("Error: Op: %v, Code: %v, Message: %v", e.Op, e.Code, e.Msg)
}

type ErrDbOperationFail struct {
	ErrCommon
}
