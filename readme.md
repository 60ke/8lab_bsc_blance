此项目用于重启bsc docker,需赋予操作docker的权限

项目支持通过当前目录下的config.toml来配置(如不含配置文件则会使用默认配置)

### 相关配置参数
```
ding-prefix : 钉钉通知接口关键字
ding-url : 钉钉机器人接口
ding-token : 钉钉通知接口token
port : 服务端口默认为8546
log-name : 日志路径,默认为bsc_balance.log
log-level : 日志等级int类型默认为info; debug || info || warn
token : 用于客户端token校验"`
docker-name : 待重启的docker名字,默认为trust_zkbsc
```


### 接口说明

```
curl --location --request POST '127.0.0.1:8546/restart_bsc' \
--header 'Content-Type: application/json' \
--data-raw '{"token":"3D3781351A3EE9E4"}'
```

出于安全考虑不支持docker名称参数,后续如有需要可做相应调整

### 其它说明
由于geth pid为1,在docker中无法通过kill -9来杀死,故将此程序放入docker外部运行