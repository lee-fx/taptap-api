<p align="center">
  <h2>
	TAPTAP-API
  </h2>
</p>


## 简介

 - [taptap-api](https://github.com/lee-fx/taptap-api) 是一个后端接口解决方案，它基于 [golang](https://github.com/golang/go) 实现。
 - 主要功能是为uniapp的前端，vue的后台管理系统（前端），实现全部后端接口。



## 框架目录（以admin为例）

```
- config         // 配置文件目录
  - config       
- controllers    // 接口逻辑
  - users        // 用户相关接口
  - ...
- dbops          // 数据库相关
  - conn_mysql   // mysql conn
  - conn_redis   // reids conn
  - user         // 用户 db 操作
- defs           // 结构体
  - apidef       // 接口原型
  - errs         // 错误原型与错误自定义
- utils          // mine 封装的工具
  - auth         
  - comm_fun    
  - response     
  - token         
- main           // 入口
- router         // 路由
- README.md      // 说明文档
```

## 开发

```bash
# 克隆项目
git clone https://github.com/lee-fx/taptap-api.git
```


Copyright (c) 2021-present Lee-Fx