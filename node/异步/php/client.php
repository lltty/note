<?php
/**
 * Created by PhpStorm.
 * User: user
 * Date: 2019/9/3
 * Time: 12:11 PM
 */

//创建一个客户端

$client = new GearmanClient();

//添加一个job服务

$client->addServer('127.0.0.1', 4730);

//doNormal是同步的，等待worker处理完成返回结果

$client->doNormal('sum', serialize(array(100)));

