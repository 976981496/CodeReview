数据配置文件当前使用yaml文件

整个项目当前主要分为三个模块

模块一：消息处理模块，接受gitlab的消息。
模块二：数据筛选。读取配置文件，查找对应提交人员的review人，@走查代码。
模块三：钉钉发送消息模块。已指定消息格式发送需要的消息



配置文件里面需要效验X-Gitlab-Token，在推送消息的gitlab中配置，不一致讲导致不满足触发钉钉机器人的推送消息的条件
