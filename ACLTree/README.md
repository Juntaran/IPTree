# ACL Tree

ACL Tree 是五元组匹配的基本结构，底层为 radix tree  
能够存储、查询相应的节点  
底层 radix tree 参考 nginx 源码实现  
当前版本并未考虑 ipv6 ，只支持 ipv4  

通过 ACL Tree 可以查询相应的 IP 的端口黑白名单  
也可以轻易地增加/删除对应的 cidr 规则  
其中可以直接增加一个端口段，便于配置  

而五元组匹配包含  

> 源 IP  
> 源端口  
> 目的 IP  
> 目的端口
> 协议号  

刨除协议号，其余四元组可以很容易地由两个 ACL Tree 表示    
一棵为 src Tree  
一棵为 dst Tree  

协议号待加入，会设置 protocol:port_list 的结构  
