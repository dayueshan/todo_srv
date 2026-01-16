# todo项目，基础版

自述: 
- 基于goframe
- 原计划是srv、mysql均单独服务，放到Harbor，再部署

## 系统说明:
- /path_to/todo-server/docker/mysql/init.sql表结构文件
  - user表，承接第三方用户的信息
  - todo表，业务数据表
  - docker初始化数据，插入一条管理员
  - 备注:
    - 用户表的数据量，增长情况
    - 用户表的扩展表，以及上一层的宽表
    - 用户数据和权限单元、以及数据权限数的结合部分
    - todo业务表是自己分，还是分布式，聚合的问题
    - 统计，clickhouse
    - 等等
- 模拟第三方oauth，手动访问接口获取token
- restful业务接口带上token校验
- 后期接入通用第三方功能、分布式数据库、部署自动化

## 程序说明:
- 仓库地址，https://github.com/dayueshan/todo_srv
- 依赖：
  - docker
  - 镜像仓库
  - golang
  - mysql
- 运行方式：
  - 1）启动mysql
    - cd /path_to/todo-server/docker
    - docker compose up -d
    - docker ps
    - docker exec -it gf-mysql mysql -uroot -p123456
  - 2）启动srv
    - cd /path_to/todo-server/
    - go build -o todo-srv ./main.go
    - ./todo-srv
    - 镜像上传到habor，打tag
  - 3）测试接口，restful
    - 检查不带token的情况
      - curl -H "Content-Type: application/json" -X GET "http://localhost:8000/user/1"
    - 用户登录，拿token；此接口唯一不需要校验的
      - curl -H "Content-Type: application/json" -X POST -d '{"provider":"google", "provider_user_id":"1002jsldfj78ksahe21faslj9sfsjakj229dfjkj"}' "http://localhost:8000/user/login"
    - 用户信息
      - curl -H "Content-Type: application/json" -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3Njg1NzA3ODYsIm5iZiI6MTc2ODU2NzE4NiwiaWF0IjoxNzY4NTY3MTg2fQ.fm9O6j-UcyDQkgjVKLByfH3LdggIl-IXeUQn4uvbXAc" -X GET "http://localhost:8000/user/1"
    - 新增todo记录
      - curl -H "Content-Type: application/json" -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3Njg1NzA3ODYsIm5iZiI6MTc2ODU2NzE4NiwiaWF0IjoxNzY4NTY3MTg2fQ.fm9O6j-UcyDQkgjVKLByfH3LdggIl-IXeUQn4uvbXAc" -X POST -d '{"user_id":1, "content":"ttttttttttdddddddddddddddddddd....todotodotodo"}' "http://localhost:8000/to_do"
    - 获取todo列表
      - curl -H "Content-Type: application/json" -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3Njg1NzA3ODYsIm5iZiI6MTc2ODU2NzE4NiwiaWF0IjoxNzY4NTY3MTg2fQ.fm9O6j-UcyDQkgjVKLByfH3LdggIl-IXeUQn4uvbXAc" -X GET "http://localhost:8000/to_do?user_id=1"
    - 修改todo记录的状态
      - curl -H "Content-Type: application/json" -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3Njg1NzA3ODYsIm5iZiI6MTc2ODU2NzE4NiwiaWF0IjoxNzY4NTY3MTg2fQ.fm9O6j-UcyDQkgjVKLByfH3LdggIl-IXeUQn4uvbXAc" -X PUT -d '{"status":1}' "http://localhost:8000/to_do/1"
    - 删除某条记录
      - curl -H "Content-Type: application/json" -H "Content-Type: application/json" -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJleHAiOjE3Njg1NzA3ODYsIm5iZiI6MTc2ODU2NzE4NiwiaWF0IjoxNzY4NTY3MTg2fQ.fm9O6j-UcyDQkgjVKLByfH3LdggIl-IXeUQn4uvbXAc" -X DELETE "http://localhost:8000/to_do/1"
    - 更多操作见swagger接口文档
    - 单元测试，behat等
      - 后期
  - 4）接口文档
    - 完整内容，http://127.0.0.1:8000/swagger/
    - 或导出的文件，api.json

## 脚本:
- 后期

## 代码目录:
- api 接口定义
- docker
- hack 便捷脚手架
- internal 实际业务代码
  - controller
  - dao
  - service
  - entity
  - ...
- manifest
- output
  - 运行截图，其它
