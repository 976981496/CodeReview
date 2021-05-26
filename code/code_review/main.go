package main

import (
	"code_review/common"
	"code_review/data_fetch"
	"code_review/dingding"
	"code_review/golog"
	"code_review/handle"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
)

func main() {

	// init config and ...
	err := common.Init()
	if err != nil {
		log.Error(fmt.Sprintf("init config err: %s", err))
		return
	}
	//log
	hook, err := golog.NewVislogHook(common.GlobalConfig.Log.FileName)
	if err != nil {
		log.Error("can't add logfile hook")
		return
	}
	log.AddHook(hook)
	tf := log.TextFormatter{
		DisableColors: true,
	}
	log.SetFormatter(&tf)

	var logLevel log.Level
	switch common.GlobalConfig.Log.LogLevel {
	case "DEBUG":
		logLevel = log.DebugLevel
	case "INFO":
		logLevel = log.InfoLevel
	case "WARN":
		logLevel = log.WarnLevel
	case "ERROR":
		logLevel = log.ErrorLevel
	default:
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
	// 遍历 os.Args 切片，就可以得到所有的命令行输入参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%v\n", i, v)
	}
	log.Info("test")
	// data_fetch.ReadYamlConfFile("wangda@orbbec.com")
	// log.Info("os.Args[1]:",os.Args[1] )
	//str :=`{"object_kind":"push","before":"95790bf891e76fee5e1747ab589903a6a1f80f22","after":"da1560886d4f094c3e6c9ef40349f7d38b5d27d7","ref":"refs/heads/master","checkout_sha":"da1560886d4f094c3e6c9ef40349f7d38b5d27d7","user_id":4,"user_name":"wangda","user_username":"wangda","user_email":"wangda@orbbec.com","project_id":15,"project":{"name":"Diaspora","description":"","web_url":"http://example.com/mike/diaspora","avatar_url":null,"git_ssh_url":"git@example.com:mike/diaspora.git","git_http_url":"http://example.com/mike/diaspora.git","namespace":"Mike","visibility_level":0,"path_with_namespace":"mike/diaspora","default_branch":"master"},"commits":[{"id":"b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327","message":"Update Catalan translation to e38cb41.","timestamp":"2011-12-12T14:27:31+02:00","url":"http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327","author":{"name":"Jordi Mallach","email":"jordi@softcatala.org"},"added":["CHANGELOG"],"modified":["app/controller/application.rb"],"removed":[]},{"id":"da1560886d4f094c3e6c9ef40349f7d38b5d27d7","message":"fixed readme","timestamp":"2012-01-03T23:36:29+02:00","url":"http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7","author":{"name":"GitLab dev user","email":"gitlabdev@dv.com"},"added":["CHANGELOG"],"modified":["app/controller/application.rb"],"removed":[]}],"total_commits_count":4}`
	
	git_info, err := handle.HandleGitInfo(os.Args[1])
	if err != nil {
		log.Error(err)
		return
	}
	//	review, webhook,err1 := data_fetch.ReadYamlConfFile("wangda@orbbec.com")
	review, webhook, err1 := data_fetch.ReadYamlConfFile(git_info.UserEmail)
	// fmt.Printf("args%s\n",git_info.UserEmail)
	if err1 != nil {
		log.Info(err1)
	}
	dingding.SendDingMsg(git_info, review, webhook)

}
