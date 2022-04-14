package autotask

import (
	"io/fs"
	"path/filepath"
	"strings"
	"time"
	"zenith/zenithsdk/httpsdk"
)

func registerTaskToSyncPicToOss() {
	c := time.NewTicker(time.Minute * 5).C
	for {
		<-c
		rootPath := httpsdk.GetPicSavePath()
		if err := filepath.Walk(rootPath, func(filepath string, info fs.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			go func() {
				if err := uploadFileToOss(filepath, strings.TrimLeft(filepath, rootPath)); nil != err {
					return
				}
				if err := removeFile(filepath); nil != err {
					return
				}
			}()
			return nil
		}); nil != err {
			continue
		}
	}
}

func uploadFileToOss(filePath, objKey string) error {
	return nil
}

func removeFile(filePath string) error {
	return nil
}
