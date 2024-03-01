# projectzero 
个人用GIN项目脚手架

## 已使用模块

1. [Gin](https://github.com/gin-gonic/gin): 轻量级Web框架
2. [entgo](https://github.com/ent/ent): 数据库ORM组件
3. [Golang-JWT](https://github.com/golang-jwt/jwt): Golang JWT模块
4. [viper](https://github.com/spf13/viper): 使用viper为项目提供配置文件支持
5. [zap](https://github.com/uber-go/zap): uber开源的json日志库


当前已实现接口
1. ```/api/v1/user/register```用户注册
2. ```/api/v1/user/login```用户登录
3. ```/api/v1/user/info```用户信息
4. ```/api/v1/user/update```更新信息

MySQL数据表将会在启动项目时自动创建

## 参考项目
[Singo](https://github.com/gourouting/singo): Gin+Gorm开发Golang API快速开发脚手架

[go-zero](https://github.com/zeromicro/go-zero): 具有CLI工具的云原生GO微服务框架
