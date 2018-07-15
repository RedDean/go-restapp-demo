package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"go-restapp-demo/pkg/errno"
	. "go-restapp-demo/handler"
)

// 注册
func Create(c *gin.Context) {

	var user CreateRequest
	var err error

	if err := c.Bind(&user); err != nil {
		SendRespnose(c, errno.ErrBind, nil)
		return
	}

	admin2 := c.Param("username")
	log.Infof("URL username: %s", admin2)

	desc := c.Query("desc")
	log.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	log.Infof("Header Content-Type: %s", contentType)

	log.Debugf("username is: [%s], password is [%s]", user.Username, user.Password)
	if user.Username == "" {
		err = errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")).Add("This is add message.")
		SendRespnose(c, err, nil)
	}

	if errno.IsErrUserNotFound(err) {
		log.Debug("err type is ErrUserNotFound")
	}

	if user.Password == "" {
		err = fmt.Errorf("password is empty")
		SendRespnose(c, err, nil)
	}

	resp := CreateResponse{
		user.Username,
	}
	SendRespnose(c, nil, resp)
}
