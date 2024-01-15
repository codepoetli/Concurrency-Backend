## internal

- controller 层 ：负责与用户（前端）交互
- service 层 ： 处理业务逻辑
- dao 层 ： 与数据库或其他数据存储系统进行交互，获取和存储数据； 基于Gorm框架
  - database.go 连接数据库
  - 所有写操作都通过一个事务来完成
- model : 定义对象数据模型（数据库实体），对应数据库里的表； 基于Gorm框架
- oss : 对象存储相关逻辑