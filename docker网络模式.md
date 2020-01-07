Nat: 网络地址转化，使用iptables的nat docker内网通过docker0进行通讯
Host: 和宿主机共享网络,没有经过iptables转发，性能好，但是网络环境没有和宿主机隔离，流量统计不好区分，端口管理也不容易。
Other Container: 容器间网络通讯频繁的情况
None: 不创建网络，自行配置