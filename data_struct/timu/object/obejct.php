<?php
/**
 * Created by PhpStorm.
 * User: user
 * Date: 2019/3/10
 * Time: 7:49 PM
 */

$input = [
    ['01', '204521'],
    ['23', '204523'],
    ['22', '204526'],
    ['01', '204528'],
    ['22', '204527']
];


function getConflictObject($input) {
    $len = count($input);
    $objectMap = [];
    for($i= 0; $i < $len; $i++) {

        $day = $input[$i][0];
        $hour = $input[$i][1];


    }
}

class MyDB extends SQLite3{

}