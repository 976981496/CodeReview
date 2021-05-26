    export GO111MODULE=on
    export GOPROXY=https://goproxy.io
    #生成go.mod,在当前目录下初始化 go.mod(就是会新建一个 go.mod 文件)
    go mod init  code_review
    #整理依赖关系，会添加丢失的 module，删除不需要的 module
    go mod tidy
    #如果有版本不一致，编辑go.mod, replace
    #例如本例, 在go.mod中增加replace github.com/Sirupsen/logrus v1.4.2 => github.com/sirupsen/logrus v1.0.6
    #运行本文件时如果有库需要replace终端会提示，然后按提示修改就行，替换的版本可以v0.0.0或和前面的版本一致，这个替换的名字就是
    #GOPATH下的pkg/mod找到对应的包里面的go.mod的定义的模块名字

    #以上会生成go.mod文件的及内容