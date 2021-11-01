package sessions

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/boj/redistore.v1"
)

var client *redistore.RediStore

func SessionInit() error {
	var err error
	// TODO: Add secret key
	client, err = redistore.NewRediStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret key"))
	if err != nil {
		return err
	}
	return nil
}

func GetSessionClient() *redistore.RediStore {
	return client
}

func GetUserNameBySession(c *gin.Context) string {
	session, _ := client.Get(c.Request, "dotcomUser")
	ret, ok := session.Values["username"]
	if !ok {
		return ""
	}
	return ret.(string)
}
