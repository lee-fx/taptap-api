<p align="center">
  <h2>
	TAPTAP-ADMIN
  </h2>
</p>


## 简介

 - [taptap-admin](https://github.com/lee-fx/taptap-admin) 是一个后端接口解决方案，它基于 [golang](https://github.com/golang/go) 实现。
 - 主要功能是为uniapp的前端，vue的后台管理系统，实现全部接口。



## 框架目录（以api为例）

```
- controllers  // 接口逻辑
  - games        // 游戏相关接口
  - home         // 主页相关接口
  - users        // 用户相关接口
- dbops       // 数据库相关
  - api_***      // 控制器逻辑对应的db 
  - conn         // 数据库连接
  - internal     // session db 相关
- defs       // 静态资源目录
  - apidef       // 接口原型
  - image        // 错误处理格式
- session    // session
- utils      // 工具
  - response     // 响应同意处理
  - uuid         // 唯一id
- README.md  // 说明
```

## 开发

```bash
# 克隆项目
git clone https://github.com/lee-fx/taptap-admin.git
```


Copyright (c) 2021-present Lee-Fx