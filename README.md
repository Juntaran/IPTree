# IPTree

使用二叉树存储路由规则，可以自由增加、删除、查找IP是否存在于路由


## 现有状态:

`Binary_IP_Tree`:  

前缀匹配路由，端口号配置为 int  
可以修改为 端口号 []int，存储一组端口号，这样形成一个 路由+端口 结构  
两个该结构一个定为 src ，一个定为 dst ，组合为四元组  

`Radix_IP_Tree`:  

前缀匹配，golang 实现 ngx_radix_tree  
value 值也就是 端口号为 Uint32Slice 类型 (uint32[])  


## Todo:

| ID  | Todo  | Done  |
|---|---|---|
| 1  | 增加端口号范围  | √  |
| 2  | 五元组匹配  |   |
| 3  | 黑白名单  | √  |
| 4  | Golang重写  | √  |
| 5  | Radix Tree版本  | √  | 

___

Reference:
* [ip转发二叉树查找方法](http://www.cnblogs.com/letusrock/p/4321983.html)
* [go-iptree](https://github.com/zmap/go-iptree)
* [nginx_core](https://trac.nginx.org/nginx/browser/nginx/src/core/)
* [nradix](https://github.com/asergeyev/nradix)