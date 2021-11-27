package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/gomarkdown/markdown"
	"github.com/hunterhug/go_image"
)

var logger = logs.GetBeeLogger()

func TimestampToDate(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func GetCurrentUnix() int64 {
	fmt.Println(time.Now().Unix())
	return time.Now().Unix()
}

func GetCurrentUnixNano() int64 {
	return time.Now().UnixNano()
}

func GetCurrentDate() string {
	template := "2006-01-02 15:04:05"
	return time.Now().Format(template)
}

func GetCurrentDay() string {
	template := "20060102"
	return time.Now().Format(template)
}

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return string(hex.EncodeToString(m.Sum(nil)))
}

func VerifyEmail(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func GenerateOrderId() string {
	template := "200601021504"
	return time.Now().Format(template) + GetRandomNum()
}

func SendMsg(str string) {
	ioutil.WriteFile("test_send.txt", []byte(str), 06666)
}

func ResizeImage(filename string) {
	extName := path.Ext(filename)
	resizeImageSize, _ := web.AppConfig.String("resizeImageSize")
	resizeImage := strings.Split(resizeImageSize, ",")

	for i := 0; i < len(resizeImage); i++ {
		w := resizeImage[i]
		width, _ := strconv.Atoi(w)
		savepath := filename + "_" + w + "x" + w + extName
		err := go_image.ThumbnailF2F(filename, savepath, width, width)
		if err != nil {
			logger.Error("error", err)
		}
	}
}

func FormatImage(picName string) string {
	ossStatus, err := web.AppConfig.Bool("ossStatus")
	if err != nil {
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName
	}

	if ossStatus {
		ossDomain, err := web.AppConfig.String("ossDomain")
		if err != nil {
			return ossDomain + "/" + picName
		} else {
			return "error"
		}
	} else {
		flag := strings.Contains(picName, "/static")
		if flag {
			return picName
		}
		return "/" + picName
	}
}

func FormatAttribute(str string) string {
	md := []byte(str)
	htmlByte := markdown.ToHTML(md, nil, nil)
	return string(htmlByte)
}

func Mul(price float64, num int) float64 {
	return price * float64(num)
}

func GetRandomNum() string {
	var str string
	for i := 0; i < 4; i++ {
		current := rand.Intn(10)
		str += strconv.Itoa(current)
	}
	return str
}
