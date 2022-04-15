package autotask

func Init() {
	go registerTaskToSyncPicToOss()
}
