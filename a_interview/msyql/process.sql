--  用来批量生成大量数据的存储过程

USE test;

-- 创建测试table
DROP TABLE IF EXISTS users;
CREATE TABLE users (
  id int primary key auto_increment,
  name varchar(10) not null,
  addr varchar(2) not null,
  url varchar (90) not null
) ENGINE = INNODB DEFAULT CHARSET UTF8;

-- 重新设定定界符
delimiter //

-- 定义一个用于生成随机字符串的函数,返回的字符串长度控制在20个字符以内
DROP FUNCTION IF EXISTS randStr  //
CREATE FUNCTION randStr(num INT) RETURNS varchar(20)
BEGIN
    DECLARE randStr varchar(62) DEFAULT 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
    DECLARE returnStr varchar(20) DEFAULT '';
    IF num > 20 THEN
      SET num = 20;
    END IF;
    WHILE num > 0 DO
        SET returnStr = concat(returnStr,substring(randStr , FLOOR(1 + RAND()*62),1));
        SET num = num - 1;
    END WHILE;
    RETURN returnStr;
END;
//


DROP FUNCTION IF EXISTS rand_0_1  //
CREATE FUNCTION rand_0_1() RETURNS tinyint
BEGIN
    DECLARE randStatus tinyint DEFAULT 0;
    SELECT ROUND(RAND() * 10) % 2 INTO randStatus;
    RETURN randStatus;
END;

-- 定义一个随机函数用于随机用户地址
DROP FUNCTION IF EXISTS randAddr  //
CREATE FUNCTION randAddr() RETURNS varchar(2)
BEGIN
    DECLARE bj varchar(2) DEFAULT '北京';
    DECLARE sh varchar(2) DEFAULT '上海';
    DECLARE gz varchar(2) DEFAULT '广州';
    DECLARE sz varchar(2) DEFAULT '深圳';
    DECLARE randNum int DEFAULT 0;
    DECLARE returnStr varchar(20) DEFAULT '';

    SET randNum = FLOOR(RAND() * 4);
    IF randNum = 0 THEN
      SET returnStr = bj;
    ELSEIF randNum = 1 THEN
      SET returnStr = sh;
    ELSEIF randNum = 2 THEN
      SET returnStr = gz;
    ELSE
      SET returnStr = sz;
    END IF;
    RETURN returnStr;
END;
//

-- 定义一个生成随机url的函数
DROP FUNCTION IF EXISTS randUrl //
CREATE FUNCTION randUrl() RETURNS varchar(90)
BEGIN

  DECLARE baidu varchar(100) DEFAULT "http://www.baidu.com";
  DECLARE aiqyi varchar(100) DEFAULT "http://www.aiqiyi.com";
  DECLARE mengtui varchar(100) DEFAULT "http://www.me.com";
  DECLARE duoduo varchar(100) DEFAULT "http://www.pingduoduo.com";
  DECLARE cctv varchar(100) DEFAULT "http://edu.cctx1.com";
  DECLARE mongo varchar(100) DEFAULT "http://edu.mongo.com";
  DECLARE csdn varchar(100) DEFAULT "https://www.csdn.net";

  DECLARE randNum int DEFAULT 0;
  DECLARE returnStr varchar(90) DEFAULT '';
  DECLARE subfix varchar(60) DEFAULT '';

  SET randNum = FLOOR(RAND() * 7);

  SET subfix = CONCAT(randStr(10), "/", randStr(20) , "/" , randStr(10), "/" , randStr(15));

  IF randNum = 0 THEN
    SET returnStr = CONCAT(baidu, "/", subfix);
  ELSEIF randNum = 1 THEN
    SET returnStr = CONCAT(aiqyi, "/", subfix);
  ELSEIF randNum = 2 THEN
    SET returnStr = CONCAT(mengtui, "/", subfix);
  ELSEIF randNum = 3 THEN
    SET returnStr = CONCAT(duoduo, "/", subfix);
  ELSEIF randNum = 4 THEN
    SET returnStr = CONCAT(cctv, "/", subfix);
  ELSEIF randNum = 5 THEN
    SET returnStr = CONCAT(mongo, "/", subfix);
  ELSE
    SET returnStr = CONCAT(csdn, "/", subfix);
  END IF;
  RETURN returnStr;

END;
//

-- 定义存储过程, 输入参数num表示一次插入多少数据,输入参数debug用于调试
DROP PROCEDURE IF EXISTS createMassData  //
CREATE PROCEDURE createMassData(IN num int, IN debug tinyint)
BEGIN
  DECLARE debugSql varchar(100) DEFAULT '';
  DECLARE nameField varchar(20) DEFAULT '';
  DECLARE addrField varchar(20) DEFAULT '';
  DECLARE urlField  varchar(100) DEFAULT '';
  DECLARE insertSqlPrefix varchar (50) DEFAULT 'INSERT INTO test.users(name, addr, url) VALUES ';
  DECLARE insertValue varchar (140) DEFAULT '';
  WHILE num > 0 DO
    SET nameField = randStr(10);
    SET addrField = randAddr();
    SET urlField = randUrl();
    SET insertValue = concat('(', nameField, ',', addrField, ',', urlField, ')');
    IF debug THEN
      SET debugSql = concat(insertSqlPrefix, insertValue);
      SELECT debugSql AS debug;
    END IF;
    INSERT INTO test.users(name, addr, url) VALUES (nameField, addrField, urlField);
    SET num = num - 1;
  END WHILE;
END;
//

/*DROP PROCEDURE IF EXISTS createMassData_t01_bitch  //
CREATE PROCEDURE createMassData_t01_bitch(IN num int, IN debug tinyint)
BEGIN
  DECLARE debugSql varchar(100) DEFAULT '';
  DECLARE nameField varchar(20) DEFAULT '';
  DECLARE statusField varchar(20) DEFAULT '';
  DECLARE insertSqlPrefix varchar (50) DEFAULT 'INSERT INTO test.user(name, status) VALUES ';
  DECLARE insertValue varchar (50) DEFAULT '';
  DECLARE insertValueAll text DEFAULT '';
  WHILE num > 0 DO
    SET nameField = randStr_10();
    SET statusField = rand_0_1();
    SET insertValue = concat('(', nameField, ',', statusField, ')');
    IF debug THEN
      SET debugSql = concat(insertSqlPrefix, insertValue);
      SELECT debugSql AS debug;
    END IF;

    IF num > 1 THEN
      SET insertValueAll = concat(insertValueAll, insertValue, ',');
    ELSE
      SET insertValueAll = concat(insertValueAll, insertValue);
    END IF;

    SET num = num - 1;
  END WHILE;

  INSERT INTO test.user(name, status) VALUES insertValueAll;

END;
//*/

-- 查看所有自己定义的函数：
-- SELECT SPECIFIC_NAME FROM information_schema.Routines WHERE ROUTINE_TYPE = 'FUNCTION' limit 100 \G;
-- 恢复定界符
delimiter ;

