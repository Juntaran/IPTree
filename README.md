# IPTree

使用二叉树存储路由规则，可以自由增加、删除、查找IP是否存在于路由


## 现有状态:

前缀匹配路由，端口号配置为 int  
可以修改为 端口号 []int，存储一组端口号，这样形成一个 路由+端口 结构  
两个该结构一个定为 src ，一个定为 dst ，组合为四元组  

## Todo:

 1. 增加端口号范围  
 2. 增加功能使之成为防火墙  
 2.1 五元组匹配  
 2.2 黑白名单  
 3. Golang重写  
 4. 修改出一个Radix Tree版本，删除操作可以采用惰性删除  

___

Reference:
* [ip转发二叉树查找方法](http://www.cnblogs.com/letusrock/p/4321983.html)
* [go-iptree](https://github.com/zmap/go-iptree)