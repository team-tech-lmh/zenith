package httpsdk

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/team-tech-lmh/zenith/utils"
)

var (
	openbarrireMap = sync.Map{}
)

func openBarrierAt(ipAddr string) {
	utils.DefaultSwitchLogger.Printf("wait open barrier on %v", ipAddr)
	ch := make(chan int, 1024)
	v, has := openbarrireMap.Load(ipAddr)
	if has {
		ch = v.(chan int)
	}
	ch <- 1
	openbarrireMap.Store(ipAddr, ch)
}

func shouldOpenBarrierAt(ipAddr string) bool {
	v, has := openbarrireMap.Load(ipAddr)
	if !has {
		return false
	}
	select {
	case <-v.(chan int):
		return true
	default:
		return false
	}
}

func handleHeartBeat(ctx *gin.Context) {
	baseBeforeHandle(ctx)
	defer baseDeferHandle(ctx)

	if ip, has := ctx.RemoteIP(); has && shouldOpenBarrierAt(ip.String()) {
		utils.DefaultSwitchLogger.Printf("open barrier at %v\n", ip.String())
		ctx.JSON(http.StatusOK, openResult)
	}
}
