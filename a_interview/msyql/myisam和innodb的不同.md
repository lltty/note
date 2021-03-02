1  MyiSAM默认对字符串使用压缩索引，这会使得select的性能慢很多

2  mysql的存储引擎在工作时需要在引擎之间和服务器之间通过行缓存格式进行数据拷贝，然后再服务器层将缓冲内容解码成各个列。从行缓冲中将编码过的列转换成数据结构的代价是非常高的。而Myisam的定长行结构实际上与服务器层的行机构正好匹配，所以不需要转换。然而Myisam的变长行结构和innodb的行结构总是需要转换的。转化的代码依赖于行数。

3  myisam表，索引，数据是分开存储的，所以修改表可以通过替换.frm文件的方式

4  myisam使用前缀压缩技术使得索引更小，但inodb则是按照原数据格式进行存储。mysiam索引通过数据的物理位置引用被索引的行，而innodb则根据主键引用被索引的行。

5 数据分布不同

  innodb的数据结构和数据是存储的一起的，myisam是分开存储的，是因为innodb使用的聚簇索引。
  myisam和innodb索引都是使用的b-tree结构，但是后者是聚簇索引。myisam的主键是和其他索引一样的唯一非空索引而已。

6 Myisam在内存中只缓存索引，数据则是依赖于操作系统来缓存。

7 Innodb的二级索引的叶子节点保存的是主键值，所以通常的查询会需要依赖聚簇索引进行二次查询。

对于一些无法直接使用覆盖索引的查询，我们可以优先使用覆盖索引找出要的值，再根据覆盖索引去关联要查找的数据。

CREATE TABLE `t1` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(10) DEFAULT NULL,
  `age` int(11) DEFAULT '0',
  `nick` varchar(10) NOT NULL DEFAULT '_',
  `addr` varchar(10) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `name` (`name`,`age`,`nick`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

# 没法使用覆盖索引的情况：
SELECT * FROM t1 WHERE age = 10 and name = 'hehe';

#使用覆盖索引找到要的列再去关联数据(延迟关联)
SELECT * FROM T1 JOIN (SELECT id FROM t1 WHERE age = 10 and name = 'hehe') tmp ON t1.id = tmp.id

8 Innodb有redo-log,而Myisam没有，只能依靠服务器层面的bin-log.

9 Myisam殷勤不支持行锁。就意味着并发控制只能使用表锁，对于这种引擎的表，同一张表上任何时刻只能有一个更新在执行。Innodb支持行锁，这也是Innodb取代Myisam的重要原因之一。

10 Myisam将表的总数存在磁盘上，因此执行效率很高，而innodb执行count(*)需要把数据从引擎里面读出来，累计计算。

11 Innodb的事务支持，并发能力，数据安全方面是优于Myisqm的。

12 Myisam 引擎的自增值保存在文件中，Innodb引擎的自增值保存在内存中，Mysql8以后才有了自增值持久化功能。Mysql重启会修改AUTO_INCREMENT值。Mysql8以后自增值放在redo-log中，重启时候依靠redo_log恢复重启之前的值。

13 Myisam分区表通用分区策略，Innodb是本地分区策略，Myisam如果表分区过多，插入数据会报打开文件个数上线，Myisam分区在5.7.19以后开始禁用。