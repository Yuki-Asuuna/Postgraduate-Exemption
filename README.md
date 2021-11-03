# Postgraduate-Exemption项目（研究生推免系统）

服务器ip&port：101.34.159.82:8000

系统信息：Linux VM-16-11-centos 3.10.0-1160.11.1.el7.x86_64 #1 SMP Fri Dec 18 16:34:56 UTC 2020 x86_64 x86_64 x86_64 GNU/Linux

## 访问示例

没有申请域名，目前只能通过ip+端口号访问

Request：

```
http://101.34.159.82:8000/ping
```

Response：

```json
{"message":"pong"}
```





## Redis
由于安全问题，Redis不对公网开放

服务器ip：101.34.159.82（与服务器ip相同）

Port：6379

用户名：nil

密码：123456



## Mysql

服务器ip：101.34.159.82（与服务器ip相同）

Port：3306

用户名：root

密码：123456