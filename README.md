# MiraiMirai
基于mirai的论坛更新提醒QQ机器人
## 环境配置
- 安装miraiOK一键启动器:
  
  https://github.com/LXY1226/MiraiOK
  
- 配置mirai-api-http: 
  
  https://github.com/project-mirai/mirai-api-http

- 下载gomirai库
  
  执行命令```go get github.com/Logiase/gomirai```

## 代码思路
- 使用gomirai库初始化更新提醒机器人

- 根据论坛的RSS功能，爬取xml信息

- 创建一个goroutine每隔1min检查论坛有无更新

- 使用gomirai封装好的函数实现QQ机器人消息的收发

## 使用
- 运行miraiOK，登录机器人账号
- 按需求更改mirai文件中的机器人信息
- 运行

