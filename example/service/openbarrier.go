package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"zenith/httpsdk"

	"github.com/gin-gonic/gin"
)

func openBarrier(ctx *gin.Context) {
	buf, err := ioutil.ReadAll(ctx.Request.Body)
	if nil != err {
		ctx.JSON(http.StatusOK, err)
		return
	}
	var po httpsdk.OpenBarrierPO
	err = json.Unmarshal(buf, &po)
	if nil != err {
		ctx.JSON(http.StatusOK, err)
		return
	}
	httpsdk.OpenBarrier(po)
	ctx.JSON(http.StatusOK, "")
}
