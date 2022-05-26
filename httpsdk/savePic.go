package httpsdk

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/team-tech-lmh/zenith/utils"
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

func saveCatpurePicBase64Content(picType PicType, picBase64Content string) (string, error) {
	if len(picBase64Content) <= 0 {
		return "", nil
	}
	t := time.Now().Unix()
	md5Byte := md5.Sum([]byte(picBase64Content))
	md5Str := base64.StdEncoding.EncodeToString(md5Byte[:])
	fName := fmt.Sprintf("%v-%v.png", md5Str, t)
	y, m, d := time.Now().Date()
	fPath := fmt.Sprintf("%v/%v/%v-%v-%v/%v", picSavePath, picType, y, m, d, fName)
	dir := path.Dir(fPath)
	if err := os.MkdirAll(dir, os.ModePerm); nil != err {
		utils.DefaultSwitchLogger.Printf("save pic failed when create dir %v\n", err)
		return "", err
	}
	f, err := os.Create(fPath)
	if nil != err {
		utils.DefaultSwitchLogger.Printf("save pic failed when create file %v\n", err)
		return "", err
	}
	defer f.Close()

	buf, err := base64.StdEncoding.DecodeString(picBase64Content)
	if nil != err {
		utils.DefaultSwitchLogger.Printf("save pic failed when decode file content %v\n", err)
		return "", err
	}
	if l, err := f.Write(buf); nil != err {
		utils.DefaultSwitchLogger.Printf("save pic failed when write file content %v\n", err)
		return "", err
	} else {
		utils.DefaultSwitchLogger.Printf("zenith sdk save pic to %v (len:%v) \n", fName, l)
	}
	return fName, nil
}
