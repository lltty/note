<?php

function bubbleSort($arr)
{

    $len = count($arr);
    for ($i = 1; $i < $len; $i++) {
        for ($k = 0; $k < $len - $i; $k++) {
            if ($arr[$k] > $arr[$k + 1]) {
                $tmp = $arr[$k + 1];
                $arr[$k + 1] = $arr[$k];
                $arr[$k] = $tmp;
            }
        }
    }
    return $arr;
}

function selectSort($arr)
{
    for ($i = 0, $len = count($arr); $i < $len - 1; $i++) {
        $p = $i;
        for ($j = $i + 1; $j < $len; $j++) {
            if ($arr[$p] > $arr[$j]) {
                $p = $j;
            }
        }
        if ($p != $i) {
            $tmp = $arr[$p];
            $arr[$p] = $arr[$i];
            $arr[$i] = $tmp;
        }
    }
    return $arr;
}

function insertSort($arr)
{
    for ($i = 1, $len = count($arr); $i < $len; $i++) {
        $tmp = $arr[$i];
        for ($j = $i - 1; $j >= 0; $j--) {
            if ($tmp < $arr[$j]) {
                $arr[$j + 1] = $arr[$j];
                $arr[$j] = $tmp;
            } else {
                break;
            }
        }
    }
    return $arr;
}

function quickSort($arr)
{
    $length = count($arr);
    if ($length <= 1) {
        return $arr;
    }
    $middle = $arr[0];
    $left_array = array();
    $right_array = array();
    for ($i = 1; $i < $length; $i++) {
        if ($middle > $arr[$i]) {
            $left_array[] = $arr[$i];
        } else {
            $right_array[] = $arr[$i];
        }
    }
    $left_array = quick_sort($left_array);
    $right_array = quick_sort($right_array);
    return array_merge($left_array, array($middle), $right_array);
}