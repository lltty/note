<?php
/**
 * 将一个给定的整数上的数字进行反转
 */

/*
 * 整数转成字符串反转
 */
function intReverseStr(int $n) {
    $n = $n . '';
    $reverse = '';
    for ($i = strlen($n) - 1; $i >= 0; $i--) {
        $reverse .= $n{$i};
    }
    echo $reverse;
}

/*
 * 使用数学方法反转
 */
function intReverseMath(int $n) {
    $ans = 0;
    while ($n != 0) {
        $pop = $n % 10;
        if ($ans > PHP_INT_MAX / 10 || ($ans == PHP_INT_MAX / 10 && pop > 7)) {
            echo $ans;
            return 0;
        }
        if ($ans < PHP_INT_MIN / 10 || ($ans == PHP_INT_MIN / 10 && pop < -8)) {
            echo 2;
            return 0;
        }
        $ans = $ans * 10 + $pop;
        $n /= 10;
    }
    var_dump($ans);
}

intReverseMath(PHP_INT_MAX - 10);