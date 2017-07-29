# IPTree

使用二叉树存储路由规则，可以自由增加、删除、查找IP是否存在于路由

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