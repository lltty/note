<?php
/**
 * Created by PhpStorm.
 * User: user
 * Date: 2019/9/3
 * Time: 12:16 PM
 */

//创建一个worker

$worker = new GearmanWorker();

//添加一个job服务

$worker->addServer('127.0.0.1', 4730);

//从redis批量获取数据转换成map
function batchGetDataFromRedis($keys) {

}

//从mysql批量获取数据转换成map
function batchGetDataFromMysql($ids) {

}

//将计算结果重新写入redis
function saveDataToRedis($data) {

}

//注册一个回调函数，用于业务处理

$worker->addFunction('sum', function ($job) {

    //workload()获取客户端发送来的序列化数据

    $data = unserialize($job->workload());

    $max = $data[0];

    $redisMap = batchGetDataFromRedis($max);
    $mysqlMap = batchGetDataFromMysql($max);

    $sum = 0;
    for($i = 1; $i <= $max ; $i++) {
        $sum += $i * $i * $i;
    }

    saveDataToRedis($sum);

    echo "计算结果是{$sum}\n";
    return $sum;

});


//死循环

while (true) {

    //等待job提交的任务

    $ret = $worker->work();

    if ($worker->returnCode() != GEARMAN_SUCCESS) {

        break;

    }

}