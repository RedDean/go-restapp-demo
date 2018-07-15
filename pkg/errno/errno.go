package errno

import "fmt"

// 错误码
type Errno struct {
	Code    int
	Message string
}

func (errno Errno) Error() string {
	return errno.Message
}

// 具体要自定义的错误
type Err struct {
	Code    int
	Message string
	Err     error
}

func New(errno *Errno, err error) *Err {
	return &Err{errno.Code, errno.Message, err}
}

func (err *Err) Add(message string) error {
	//err.Message = fmt.Sprintf("%s %s", err.Message, message)
	err.Message += " " + message
	return err.Err
}

func (err *Err) Addf(format string, args ...interface{}) error {
	//return err.Message = fmt.Sprintf("%s %s", err.Message, fmt.Sprintf(format, args...))

	err.Message += " " + fmt.Sprintf(format, args...)
	return err.Err
}

func (err *Err) Error() string {
	return fmt.Sprintf("Err - code: %d, message: %s, error: %s", err.Code, err.Message, err.Err)
}

func IsErrUserNotFound(err error) bool {
	code, _ := DecodeErr(err)
	return code == ErrUserNotFound.Code
}

//解析错误
func DecodeErr(err error) (int, string) {
	if err == nil {
		return OK.Code, OK.Message
	}
	/* go语言里的error是个interface
	   这也是为什么不论是Err还是Errno都要有Error()方法
	   通过断言，判断传进来的error类型
	*/
	switch typed := err.(type) {
	case *Err:
		return typed.Code, typed.Message
	case *Errno:
		return typed.Code, typed.Message
	default:
	}

	return InternalServerError.Code, err.Error()
}
