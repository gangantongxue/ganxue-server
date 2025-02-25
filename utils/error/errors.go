package error

import (
	"encoding/json"
	"log"
	"strconv"
)

type Error struct {
	Code    int    `json:"code" bson:"code"`
	Err     string `json:"error" bson:"error"`
	Message string `json:"message" bson:"message"`
}

// New 新建错误
func New(code int, err error, message string) *Error {
	if err == nil {
		return nil
	}
	return &Error{
		Code:    code,
		Err:     err.Error(),
		Message: message,
	}
}

func (e *Error) Is(code int) bool {
	return e.Code == code
}

func (e *Error) As(code int) {
	e.Code = code
}

func (e *Error) Print() {
	log.Println("code", e.Code, "error", e.Err, "message", e.Message)
}

func (e *Error) ToJson() (string, error) {
	data, err := json.Marshal(e)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (e *Error) FromJson(data string) error {
	err := json.Unmarshal([]byte(data), e)
	if err != nil {
		return err
	}
	return nil
}

func (e *Error) ToString() string {
	var str string
	str = "Error: code " + strconv.Itoa(e.Code) + " error " + e.Err + " message " + e.Message
	return str
}
