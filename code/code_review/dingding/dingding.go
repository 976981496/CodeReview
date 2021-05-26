package dingding

import (
	"code_review/common"
	"fmt"
	"strings"

	"github.com/royeo/dingrobot"
)

func SendDingMsg(git_info common.GitInfo, review_phone []string, hook string) {
	dst := ""
	for _, val := range review_phone {
		dst += "@"
		dst += val
	}
	urlstr := ""
	brancharry := strings.Split(git_info.Ref, "/")

	for _, com := range git_info.Commits {
		urlstr += "["
		strid := com.ID[0:8]
		urlstr += strid
		urlstr += "]("
		urlstr += com.URL
		urlstr += ") "
		urlstr += com.Message
		urlstr += "\n\n >"
	}

	webhook := hook

	robot := dingrobot.NewRobot(webhook)
	title := "codereview"
	text := fmt.Sprintf("%s    %s  to branch %s  at  repository   %s  \n >  %s\n\n >\n   %s    请走查代码!", git_info.UserName, git_info.ObjectKind, brancharry[2], git_info.Project.Name, urlstr, dst)
	atMobiles := review_phone
	fmt.Println("atMobiles", atMobiles)

	isAtAll := false

	err3 := robot.SendMarkdown(title, text, atMobiles, isAtAll)
	if err3 != nil {
		fmt.Println("Error:", err3)
	}
}
