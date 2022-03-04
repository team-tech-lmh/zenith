package utils

import "sync"

var (
	gLocalMessageQueue = &LocalMessageQueue{
		ch: make(chan int, 1),
	}
)

type LocalMessageQueueHandler func(interface{})
type LocalMessageQueue struct {
	QueueList sync.Map
	ch        chan int
}

func MessageSub(subkey string, f LocalMessageQueueHandler) {
	gLocalMessageQueue.ch <- 1
	defer func() {
		<-gLocalMessageQueue.ch
	}()

	list := []LocalMessageQueueHandler{}
	v, has := gLocalMessageQueue.QueueList.Load(subkey)
	if has {
		list = v.([]LocalMessageQueueHandler)
	}
	list = append(list, f)
	gLocalMessageQueue.QueueList.Store(subkey, list)
}

func MessagePub(subkey string, msg interface{}) {
	list := []LocalMessageQueueHandler{}
	v, has := gLocalMessageQueue.QueueList.Load(subkey)
	if has {
		list = v.([]LocalMessageQueueHandler)
	}
	for _, f := range list {
		go ExecWithRecovery(func() {
			f(msg)
		})
	}
}
