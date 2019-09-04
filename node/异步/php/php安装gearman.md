因为php7的版本没有支持gearman,所以目前以php5.6作为安装说明,本机开发机是mac,以mac为例

### 1 安装gearman

    brew install gearmand 
    
安装完成的gearman位于：/usr/local/Cellar/gearman
gearmand: gearman守护进程，类似 mysqld(/usr/local/Cellar/gearman/1.1.18_2/sbin)
gearadmin: 用于管理gearman,类似 mysqladmin (/usr/local/Cellar/gearman/1.1.18_2/bin)
gearman: gearman 客户端，类似mysql (/usr/local/Cellar/gearman/1.1.18_2/bin)

### 2 安装php扩展

    phpbrew ext install gearman
    
