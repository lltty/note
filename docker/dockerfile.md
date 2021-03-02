构建docker守护程序不是CLI运行的。生成过程要做的第一件事就是将整个上线文(递归)发送到守护进程。

Docker守护程序伊Dockerfile一对一的方式运营指令，如有必要将每条指令的结果提交到新镜像，然后最终输出新镜像的ID.
每条指定都是独立运行的，并且会导致新镜像的创建。因此`RUN cd /tmp` 对下一条指令不会有任何影响

Docker将尽可能的重用中间镜像(缓存)，以加速Docker build的过程。如果使用的缓存，会看到控制台有"using cache" 信息。

# directive
解析器指令是可选的，并且会影响Dockerfile处理后续行为的方式。解析器指令不会再构建中添加层，也不会显示为构建步骤，解析器指令以特殊形式的注释来书写。 `# directive=value`,单个指令只能使用一次。所有解析器指令都必须位于最顶端，docker处理完注释，空行和解析器指令之后，就不会再寻找解析器指令。而是将格式化的解析器指令的任何内容都视为注释

支持的解析器指令有两种：

* syntax #指定构建当前dockerfile的Dockerfile构建起的位置。
* escape #定义转义操作符，主要对windows的文件路径特别有效

#ARG
    ARG VERSION=latest
arg可以在from之前，也可以在from之后申明，如果实在from之前声明的，那要在后面的构建过程使用，就需要使用一个没有值的指令:

    `
    ARG docker_php_version=1.0.0
    FROM lihongweimac/php72:${docker_php_version}
    ARG docker_php_version

    RUN echo $docker_php_version
    `
如果要传递多个参数：
    docker -f xxx -t xxx --build-arg arg_name=xxx --build-arg arg_name_1=xxx .

# FROM

FROM [--platform=<platform>] <image>[:<tag>] [AS <name>]
    OR
FROM [--platform=<platform>] <image>[@<digest>] [AS <name>]

digest是镜像的sha256值，可以作为镜像的区分,所以和tag是一样的作用

docker images  --digests

AS name 是给当前构建步骤取名，便于后面从这个镜像拷贝文件，啥的，比如

    `
        FROM xxxxx AS n1
        ...

        COPY --from=n1 /build/server/xxx /
    `

# ENV 
定义Dockerfile在构建过程中的环境变量，环境变量额外支持一些bash的操作符

    ENV NAME=${name:-toby} #如果name变量设置，则NAME值为$name,否则“toby”
    ENV NAME=${name:+toby} #如果name变量设置，则NAME值为“toby”,否则为空。

# RUN
在构建镜像的过程中执行命令，可以使用shell和数组两种方式
# CMD
在创建容器的时候执行的命令，可以使用shell或者数组的方式
# ENTRYPOINT
容器创建以后执行的命令，如果ENTRYPOINT声明了，则shell方式声明的cmd会被忽略，数组方式声明的cmd会被当成是ENTRYPOINT的参数，但是这个可以被客户端传过来的参数覆盖.

这里重点说明下，CMD,ENTRYPOINT的覆盖情况
 * 如果是数组方式声明
    `
    CMD ["--host=0.0.0.0"]
    ENTRYPOINT ["php", "artisan", "serve"]
    `
被覆盖的情况：
docker run -d --name test_01 --entrypoint \-\-host=127.0.0.1 laravel:01  #--host=127.0.0.1
docker run -d --name fuck_01 laravel:01 ccc                              #php artisan serve ccc
docker run -d --name fuck_01 laravel:01                                  #php artisan serve --host=0.0.0.0

shell 相对于 数组方式 有一个优势，就是，如果容器创建的时候，想执行多个命令,前面的不是常驻进程，后面的时候常驻进程，数组方式是不支持的

  `
```  ENTRYPOINT ["echo", 1, ">>/tmp/config", "&&", "php", "artisan", "serve"] #是不行的，容器会自动退出
     ENTRYPOINT echo 1  >> /tmp/config && php artisan serve #是可以的
  `

查看容器启动的详细信息：
docker ps -a --no-trunc
查看镜像的详细信息:
docker image inspect --format='' myimage

# ONBUILD
在父镜像声明，子镜像执行的命令，换句话说就是子镜像继承父镜像，但是子镜像执行以后会清除，也就是说孙子不会继承爷爷
    `
    ONBUILD RUN echo  "儿子,我是你爸爸呀"
    `

# STOPSIGNAL
定义容器的退出信号

# HEALTHCHECK
容器的健康检查
* HEALTHCHECK [OPTIONS] command
* HEALTHCHECK NONE

OPTIONS：
* --interval=DURATION (default: 30s)
* --timeout=DURATION (default: 30s)
* --start-period=DURATION (default: 0s)
* --retries=N (default: 3)

# SHELL
覆盖默认shell,在windows下需要用power shell时候比较有用

    `
    HEALTHCHECK --interval=5m --timeout=3s \
      CMD curl -f http://localhost/ || exit 1
    `



























