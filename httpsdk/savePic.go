package httpsdk

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var (
	picSavePath = "/tmp"
)

type PicType int

const (
	PicTypeNone                  PicType = iota
	PicTypeTrigger                       //主动抬杆
	PicTypeRecognizeFile                 //识别大图
	PicTypeRecognizeFragmentFile         //识别小图
	PicTypeMax
)

func SetPicSavePath(dirPath string) {
	picSavePath = strings.TrimRight(dirPath, "/")
}

func GetPicSavePath() string {
	return picSavePath
}

func saveCatpurePicBase64Content(picType PicType, picBase64Content string) {
	if len(picBase64Content) <= 0 {
		return
	}
	t := time.Now().Unix()
	md5Byte := md5.Sum([]byte(picBase64Content))
	md5Str := base64.StdEncoding.EncodeToString(md5Byte[:])
	fName := fmt.Sprintf("%v-%v.png", md5Str, t)
	y, m, d := time.Now().Date()
	fPath := fmt.Sprintf("%v/%v/%v-%v-%v/%v", picSavePath, picType, y, m, d, fName)
	dir := path.Dir(fPath)
	if err := os.MkdirAll(dir, os.ModePerm); nil != err {
		log.Printf("save pic failed when create dir %v\n", err)
		return
	}
	f, err := os.Create(fPath)
	if nil != err {
		log.Printf("save pic failed when create file %v\n", err)
		return
	}
	defer f.Close()

	buf, err := base64.StdEncoding.DecodeString(picBase64Content)
	if nil != err {
		log.Printf("save pic failed when decode file content %v\n", err)
		return
	}
	if l, err := f.Write(buf); nil != err {
		log.Printf("save pic failed when write file content %v\n", err)
		return
	} else {
		log.Printf("save pic to %v (len:%v) \n", fName, l)
	}
}
