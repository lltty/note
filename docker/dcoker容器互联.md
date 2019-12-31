新建一个docker网络

    -d：参数指定 Docker 网络类型，有 bridge、overlay。
    docker network create -d bridge test-net
    
查看所有的网络

    docker network ls
    
运行两个容器并连接到新的网络

    docker run -itd --name mysql_network --network test-net msyql /bin/bash
    docker run -itd --name redis_network --network test-net redis /bin/bash
    
    
进入两个容器测试下容器是否可互联

    docker exec -it mysql_network /bin/bash
    docker exec -it redis_network /bin/bash
    
    如果没有ping命令(以ubuntuwe为例)：
        apt-get update
        apt install iputils-ping
        
    在两个容器里分别ping下对方:
    ping redis_network
    ping mysql_network