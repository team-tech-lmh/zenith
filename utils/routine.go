package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/ztrue/tracerr"
)

// ExecWithRecovery :
func ExecWithRecovery(op func()) {
	if nil == op {
		return
	}
	defer func() {
		if r := recover(); nil != r {
			str := fmt.Sprintf("routine panicd [%v] :\n", time.Now())
			frames := tracerr.StackTrace(tracerr.New(""))
			for _, f := range frames {
				str += fmt.Sprintf("\t %v\n", f.String())
			}
			log.Println(str)
		}
	}()
	op()
}
