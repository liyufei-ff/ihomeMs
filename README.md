# ihomeMs
项目介绍：基于微服务架构的综合性租房平台可以作为房东发布房源也可以作为客户寻找房源租房。项目使用go-micro微服务框架有房源地区、获取图片验证码、房屋管理、注册登录、用户、订单六个微服务。
技术要点：
使用go-micro和gin框架构建项目，go mod管理三方包依赖
使用consul作为服务注册中心，grpc进行远程过程调用，protobuf作为grpc间数据交换格式
使用gorm操作mysql数据库，redigo操作redis数据库
使用fastDFS加nginx进行图片的存储管理和通过网络访问
