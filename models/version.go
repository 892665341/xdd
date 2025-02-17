package models

import (
	"errors"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
)

var version = "2021081905"
var AppName = "xdd"
var pname = regexp.MustCompile(`/([^/\s]+)`).FindStringSubmatch(os.Args[0])[1]

func initVersion() {
	if Config.Version != "" {
		version = Config.Version
	}
	logs.Info("检查更新" + version)
	value, err := httplib.Get(GhProxy + "https://raw.githubusercontent.com/cdle/xdd/main/models/version.go").String()
	if err != nil {
		logs.Info("更新版本的失败")
	} else {
		// name := AppName + "_" + runtime.GOOS + "_" + runtime.GOARCH
		if match := regexp.MustCompile(`var version = "(\d{10})"`).FindStringSubmatch(value); len(match) != 0 {
			if match[1] > version {
				(&JdCookie{}).Push("版本过低，自动更新")
				err := Update()
				if err != nil {
					logs.Warn("更新失败,", err)
					return
				}
				// rsp, err := httplib.Get(GhProxy + "https://github.com/cdle/jd_study/releases/download/main/" + name).Response()
				// if err != nil {
				// 	logs.Warn("无法下载更新")
				// 	return
				// }
				// filename := ExecPath + "/." + pname
				// f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
				// if err != nil {
				// 	logs.Warn("无法创建更新临时文件：%v"+filename, err)
				// 	return
				// }
				// _, err = io.Copy(f, rsp.Body)
				// f.Close()
				// if err != nil {
				// 	logs.Warn("下载更新失败")
				// 	return
				// }
				// if err := os.Rename(filename, ExecPath+"/"+pname); err != nil {
				// 	logs.Warn("移动临时更新文件失败")
				// }
				Daemon()
			}
		}
	}
}

func Update(msgs ...interface{}) error {
	sendMessagee("小滴滴开始拉取代码", msgs...)
	return 1
}
