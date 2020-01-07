因为每次从docker官方拉取镜像，是很慢的，所以我们需要有自己的仓库来管理，docker
官方提供了docker-registry来实现。

启动registry:

    docker run -d -p 5000:5000 --name registry registry:0.9.1
    
修改本地tag名称为仓库地址

push本地image到远程仓库

jenkins 要做的操作：
    
       1 自动检查git仓库的代码是否有更新
       2 用docker build生成docker镜像
       3 docker push 将镜像传递给registry
       4 docker-compose up -d 启动容器

//为了是jenkins具备docker的功能，需要将本地docker挂载到jenkins的目录     
docker run -d -p 8080:8080 --name jenkins -v /usr/local/bin/docker:/usr/bin/docker
-v /var/run/docker.sock:/var/run/docker.sock jenkens:1.609