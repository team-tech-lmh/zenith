package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func GinJSONParameterFromCtx(ctx *gin.Context, obj interface{}) error {
	buf, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		return err
	}
	return json.Unmarshal(buf, obj)
}
