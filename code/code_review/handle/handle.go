package handle

import (
	"code_review/common"
	"encoding/json"
	"fmt"
)

func HandleGitInfo(str string) (common.GitInfo, error) {
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	//str :=`{"object_kind":"push","before":"95790bf891e76fee5e1747ab589903a6a1f80f22","after":"da1560886d4f094c3e6c9ef40349f7d38b5d27d7","ref":"refs/heads/master","checkout_sha":"da1560886d4f094c3e6c9ef40349f7d38b5d27d7","user_id":4,"user_name":"wangda","user_username":"wangda","user_email":"wangda@orbbec.com","project_id":15,"project":{"name":"Diaspora","description":"","web_url":"http://example.com/mike/diaspora","avatar_url":null,"git_ssh_url":"git@example.com:mike/diaspora.git","git_http_url":"http://example.com/mike/diaspora.git","namespace":"Mike","visibility_level":0,"path_with_namespace":"mike/diaspora","default_branch":"master"},"commits":[{"id":"b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327","message":"Update Catalan translation to e38cb41.","timestamp":"2011-12-12T14:27:31+02:00","url":"http://example.com/mike/diaspora/commit/b6568db1bc1dcd7f8b4d5a946b0b91f9dacd7327","author":{"name":"Jordi Mallach","email":"jordi@softcatala.org"},"added":["CHANGELOG"],"modified":["app/controller/application.rb"],"removed":[]},{"id":"da1560886d4f094c3e6c9ef40349f7d38b5d27d7","message":"fixed readme","timestamp":"2012-01-03T23:36:29+02:00","url":"http://example.com/mike/diaspora/commit/da1560886d4f094c3e6c9ef40349f7d38b5d27d7","author":{"name":"GitLab dev user","email":"gitlabdev@dv.com"},"added":["CHANGELOG"],"modified":["app/controller/application.rb"],"removed":[]}],"total_commits_count":4}`
	b := []byte(str)
	var ret common.GitInfo
	//解析git返回的所有信息
	m := common.AllGitInfo{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		//加错误信息处理
		return ret, err
	}
	fmt.Println("m:", m)
	ret.ObjectKind = m.ObjectKind
	ret.Ref = m.Ref
	ret.UserName = m.UserName
	ret.UserEmail = m.UserEmail
	ret.Project.Name = m.Project.Name
	ret.Project.Description = m.Project.Description
	ret.Project.Namespace = m.Project.Namespace
	ret.Project.URL = m.Project.URL

	ret.Commits = make([]common.Commit, 0)

	for _, value := range m.Commits {
		var commit common.Commit
		commit.ID = value.ID
		commit.Message = value.Message
		commit.URL = value.URL

		ret.Commits = append(ret.Commits, commit)
	}
	// ret.Commits. Message = m.Commits. Message
	// ret.Commits. URL = m.Commits. URL

	return ret, nil
}
