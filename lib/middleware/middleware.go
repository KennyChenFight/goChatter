package middleware

import (
	"github.com/KennyChenFight/goChatter/lib/constant"
	"github.com/gin-gonic/gin"
	"net/http"
	"xorm.io/xorm"
)

var (
	db *xorm.Engine
)

func Init(database *xorm.Engine) {
	db = database
}

func Wrap() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := db.NewSession()
		err0 := session.Begin()

		if err0 != nil {
			SendResponse(c, http.StatusInternalServerError, map[string]string{"error": err0.Error()})
			return
		}
		defer session.Close()

		c.Set(constant.DbSession, session)
		c.Set(constant.StatusCode, nil)
		c.Set(constant.Error, nil)
		c.Set(constant.Output, nil)
		c.Next()

		session = c.MustGet(constant.DbSession).(*xorm.Session)
		statusCode := c.GetInt(constant.StatusCode)
		err1 := c.MustGet(constant.Error)
		output := c.MustGet(constant.Output)

		if err1 == nil {
			if err := session.Commit(); err != nil {
				session.Rollback()
				SendResponse(c, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			} else {
				SendResponse(c, statusCode, output)
			}
		} else {
			session.Rollback()
			SendResponse(c, statusCode, map[string]string{"error": err1.(error).Error()})
		}
	}
}

func SendResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}
