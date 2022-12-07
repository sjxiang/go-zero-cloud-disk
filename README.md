
## CloudDisk

```

用户模块
    密码登录
    刷新 Authorization
    邮箱注册
    用户详情


存储池模块
    中心存储池资源管理
        文件上传
            文件秒传
            文件分片上传

    个人存储池资源管理
        文件存储关联
        文件列表
        文件名称修改
        文件夹创建
        文件删除
        文件移动


文件共享模块
    创建分享记录
    获取资源详情
    资源保存

```

```text
集成 go-zero

安装 goctl

goctl api new core  // 单体服务

go run core.go -f etc/core-api.yaml

 curl -i -X GET http://localhost:8888/from/you
```


https://go-zero.dev/cn/docs/quick-start/monolithic-service

https://gorm.io/zh_CN/docs/query.html

https://console.cloud.tencent.com/cos ui

https://cloud.tencent.com/document/product/436/65650 文档

https://console.cloud.tencent.com/cam/capi  权限私钥
