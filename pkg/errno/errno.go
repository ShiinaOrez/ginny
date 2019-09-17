package errno

import "fmt"

type Error struct {
    ErrorCode  string  `json:"error_code"`
    Message    string  `json:"message"`
}

func (err *Error) Error() string {
    return fmt.Sprintf("Error(%s): %s.", err.ErrorCode, err.Message)
}
