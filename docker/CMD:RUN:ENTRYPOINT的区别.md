1  CMD 和 RUN 的区别

两者都是用于执行命令，区别在于执行命令的时机不同，RUN命令适用于在 docker build构建镜像是执行的命令，而CMD命令在docker run 执行docker镜像构建容器时适用，可以动态的覆盖CMD执行的命令

2 CMD 和 ENTRYPOINT 的区别

CMD 命令是用于默认执行的，如果写了多条CMD命令，则只会执行最后一条，如果后续存在ENTRYPINT 命令，则CMD命令或被充当参数或被覆盖，而且Dockerfile中的CMD命令最终可以被在执行docker run 命令是添加的参数所覆盖。而ENTRYPOINT命令则是一定会被执行的，一般用于执行脚本。

2.1 CMD 的几种写法：

* CMD ["executable","param1","param2"]（exec形式，这是首选形式）
* CMD ["param1","param2"]（作为ENTRYPOINT的默认参数）
* CMD command param1 param2（外壳形式）

shell 写法
FORM centos
CMD echo 'hello'

exec 写法
FROM centos
CMD ["echo", "hello"]

CMD只能有一条指令，如果列出多条，只有最后一条会被执行，注意如果是exec的写法，参数必须要使用双引号包裹，因为命令最终会被转换为json序列。

2.2 在shell写法环境下
在shell写法中，如果存在ENTRYPOINT命令，则不管是子Docker中存在CMD命令也好，还是在docker run 执行的后面添加添加的命令也好，都不会被执行。如果不存在ENTRYPOINT命令，则可以被docker run 后面设置的命令覆盖，实现动态操作

2.3 在exec 写法环境下
在exec写法中，如果存在ENTRYPOINT命令，则在Dockerfile中如果存在CMD命令或者是在docker run 执行的后面添加的命令，会被当做是ENTRYPOINT命令的参数来使用


2.4 给ENTRYPOINT传参
docker run xxx --entrypoint xxx




































