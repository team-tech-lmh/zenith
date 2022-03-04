package httpsdk

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func openBarrierAt(ipAddr string) {
}

func handleHeartBeat(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)

	// 如果有需要开门
	select {
	case <-openCh:
		ctx.JSON(http.StatusOK, openResult)
	default:
		return
	}
}
