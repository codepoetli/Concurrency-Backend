# Concurrency-Backend
面向高并发评论业务场景的后端

```shell
cd cmd
go run main.go
```

# 关于三层架构
三层架构的核心设计理念是“关注点分离”（Separation of Concerns, SoC），旨在将不同的功能区分开来，以降低系统各部分之间的依赖性  

- 控制层（Controller）：控制层位于架构的最前端，直接面对用户或外部请求，提供对外部的接口。它负责接收用户的输入，并将请求转发给相应的服务层处理，最后返回处理结果给用户。控制层的主要任务是请求处理和数据转发，它不包含业务逻辑，确保了用户接口的轻量化
  - 不在Controller里暴露Service的业务逻辑，而是直接转发Service的业务处理结果 
  - 接受请求，返回响应对象或错误信息
- 业务逻辑层（Service）：服务层是三层架构中的中心，承担着处理应用程序核心业务逻辑的任务。这一层解释用户的请求，执行必要的业务计算，调用数据访问层进行数据持久化操作，并返回执行结果。将业务逻辑封装在Service层中有助于保持业务处理的一致性和复用性
  - 实现并封装复杂的业务流程
  - 注意事务控制：处理涉及多个数据库操作的业务时，确保这些操作要么全部成功，要么全部失败（即事务的原子性）
  - 可能还负责与外部服务的集成，例如调用其他微服务、消息队列、第三方API等
- 数据访问层（DAO）：数据访问层是与数据库或其他持久化存储方式直接交互的层次。DAO层的职责是执行具体的数据库操作，如增删查改（CRUD），并返回操作结果。通过抽象化数据访问，DAO层使得业务逻辑层与数据存储细节解耦，提高了系统的适应性和稳定性
  - 管理与数据库的连接以及连接池的使用（通常由ORM工具或框架自动管理）；实现CRUD
  - 数据转换：将数据库中的数据转换为应用程序中的对象（DTO，Data Transfer Object），以及将对象转换为适合存储在数据库中的格式
  - 使Service不必关心SQL语句具体的数据库实现

---

## todo
加入Redis  
扩展为微服务架构  
加入Kafka  
（Zookeeper维护Kafka集群中各组件的状态信息）  