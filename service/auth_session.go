// 鉴权
package service

import (
	"Postgraduate-Exemption/constant"
	"Postgraduate-Exemption/utils/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		store := sessions.GetSessionClient()
		session, err := store.Get(c.Request, "dotcomUser")
		if err != nil {
			c.Abort()
			logrus.Errorf(constant.Service+"AuthMiddleWare Store Get Session Failed, err= %v", err)
			c.JSON(http.StatusUnauthorized, GenResponseWithUnauthorized())
			return
		}
		if session.IsNew {
			c.Abort()
			logrus.Errorf(constant.Service+"AuthMiddleWare Store Session Is New, err= %v", err)
			c.JSON(http.StatusUnauthorized, GenResponseWithUnauthorized())
			return
		}
		if isauth, ok := session.Values["authenticated"].(bool); !ok || !isauth {
			c.JSON(http.StatusUnauthorized, GenResponseWithUnauthorized())
			logrus.Infof(constant.Service + "AuthMiddleWare Store Values Is False")
			c.Abort()
			return
		}
		c.Next()
	}
}
