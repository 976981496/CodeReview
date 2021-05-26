package data_fetch

import (
	"code_review/common"
	// "encoding/json"
	// "os"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// func ReadConfFile(email string) ([]common.ReviewItem, string) {
// 	// 打开文件
// 	file, _ := os.Open("./conf/conf.json")
// 	// 关闭文件
// 	defer file.Close()

// 	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
// 	decoder := json.NewDecoder(file)

// 	conf := common.ConfigInfo{}
// 	//Decode从输入流读取下一个json编码值并保存在v指向的值里
// 	err := decoder.Decode(&conf)
// 	// for range 遍历数组\
// 	var submitter []common.ReviewItem
// 	for index, value := range conf.CodeReview {
// 		//与配置文件做对比
// 		if email == value.Submitter {
// 			fmt.Println(index, value)
// 			submitter = value.ReviewList
// 			break
// 		}
// 		fmt.Println(submitter)
// 	}
// 	return submitter, conf.WebHook
// }

func ReadYamlConfFile(email string) ([]string, string, error) {
	// 打开文件
	data, _ := ioutil.ReadFile("./conf/config.yaml")
	fmt.Println(string(data))
	var t common.YamlConfig
	//把yaml形式的字符串解析成struct类型
	err := yaml.Unmarshal(data, &t)

	if err != nil {
		fmt.Println("Umarshal failed yaml:", err)
		//加错误信息处理
		return nil, "", err
	}
	fmt.Println("初始数据", t.WebHook)
	fmt.Println("初始数据", t.CodeReview)
	fmt.Println("email", email)

	var phoneList []string
	for index, value := range t.CodeReview {
		//与配置文件做对比
		if email == value.Submitter {
			fmt.Println(index, value)
			for _, value_item := range value.ViewerList {
				phoneList = append(phoneList, value_item.Phone)
			}
			break
		}
	}
	fmt.Println("phoneList====",phoneList)
	return phoneList, t.WebHook, nil
}
