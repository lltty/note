<?php

/*
 * 给定一个数组，和目标值，在数组中找出和为目标值的那两个整数
 * @params $ary 用来查找的数组
 * @params $target 查找的目标
 * @params $repeat 找出的值是否重复, 比如1,2和2,1
 */

function filterTargetCombine(array $ary, int $target, bool $repeat = true) {
    //转成map结构
    $map = [];
    for ($i = 0, $j = count($ary); $i < $j; $i++) {
        $map[$ary[$i]] = $i;
    }
    $filters = [];
    foreach($map as $key => $val) {
        $diff = $target - $key;
        if (isset($map[$diff])) {
            if ($repeat) {
                $filters[] = [$val, $key, $map[$diff], $diff];
            } else {
                $index = noOrderIndex($key, $diff, $target);
                if (!isset($filters[$index])) {
                    $filters[$index] = [$val, $key, $map[$diff], $diff];
                }
            }
        }
    }
    echo "<pre>";
    var_dump($filters);
}

function noOrderIndex(int $n1, int $n2, int $target = 0) {
    return $n1 * $n2 + $target;
}

$ary = [1, 20, 16, 20, 9, -5, -10, 21, -11];
$target = 10;
filterTargetCombine($ary, $target, false);