<?php
/**
 * Created by PhpStorm.
 * User: user
 * Date: 2019/3/19
 * Time: 11:10 AM
 */
function findKing($m , $n) {
    $monkey =   range(1, $m);
    $index = 0;
    while(count($monkey) > 1) {
        $index++;
        $head = array_shift($monkey);
        if ($index % $n != 0) {
            array_push($monkey, $head);
        }
    }
    return $monkey[0];
}

function xipai($num) {
    $tmp = range(0, $num);
    $card = [];
    for($i = 0; $i < $num; $i++) {
        $rand = rand(0, $num - ($i + 1));
        $card = $tmp[$rand];
        array_shift($tmp);
    }
}

//echo findKing(40, 6);