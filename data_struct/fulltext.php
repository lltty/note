<?php
/**
 * Created by PhpStorm.
 * User: user
 * Date: 2019/3/9
 * Time: 3:42 PM
 */

/*
 * 查找字符串是否存在(全文查找)
 * @param $key 要查找的字符串
 * @param $content 目标文章
 * @index 最终找到的位置
 * @repeat boolean 是否重复查找所有位置
 * @return boolean
 */

function findKey($key, $content, &$index) {
    $len = mb_strlen($content);
    $keyLen = mb_strlen($key);
    $j = 0;
    $index = [];
    $findOIndex = 0;

    for ($i = 0; $i < $len; $i++) {
        $t = mb_substr($content, $i, 1);
        $findOIndex++;

        for (; $j < $keyLen; $j++) {
            $findOIndex++;
            $k = mb_substr($key, $j, 1);
            if ($k != $t) {
                $j = 0;
            } else {
                //如果当前查找的是关键字的最后一个字符，则已经找到该字符串
                if ($j == $keyLen - 1) {
                    $index[] = $i - $keyLen + 1;
                    break;
                }
                $j++;
            }
            continue 2;
        }
    }

    echo  "要查找的关键字长度：", $keyLen,"\n";
    echo  "要查找的文章长度：", $len,"\n";
    echo  "总共查找次数：", $findOIndex,"\n";
    echo  "查找到的所有结果：", "\n";
    var_dump($index);


    if (!empty($index)) {
        return TRUE;
    } else {
        return FALSE;
    }

}

function testFindKey()
{
    $key = '李红卫';
    $content = '这篇李文章是红卫我写的，爸爸吧，吧啦啦啦啦，李红卫,i am 红色的， toby,toby是谁呀，他怎么呢，卫红是个傻逼吗，哦，对了，李红卫,wo cuo le ,试试看，搜索时间, 李红卫';
    $index = [];
    findKey($key, $content, $index);
}

/*
 * 生成关键字索引map
 */

function filterWordsMap($filterWords) {
    $filterWords = str_replace('，', ',', $filterWords);
    $filterWords = preg_replace('/^\s*,\s*|\s*,\s*$/', '', $filterWords);
    $words = preg_split('/\s*,\s*/', $filterWords);
    $wordNum = count($words);

    $tree = [];
    for($i = 0; $i < $wordNum; $i++) {
        //模式修正符 u 表示一个符合unicode编码规则的串，比如utf-8编码的串
        //在u修饰符下，一个汉字被当做一个字符被处理。\w由原来的[_0-9A-Za-z]扩展到汉字
        $word = preg_split('//u', $words[$i]);
        array_shift($word);
        array_pop($word);

        $len = count($word);
        $node = [];
        for($j = $len - 1; $j >= 0; $j--) {
            if ($j == $len - 1) {
                $node = [
                    $word[$j] => 1
                ];
            } else {
                $node[$word[$j]] = $node;
                unset($node[$word[$j + 1]]);
            }
        }
        $tree[$word[0]] = $node[$word[0]];
    }

    return $tree;
}

/*
 * 改进版的查找字符串是否存在(全文查找)
 * $filterMany 是否对所有关键字都查找
 */
function findKeyPerformance($contents, $filterMany = FALSE) {

    $filterWords = '中华人名共和国,我贼,fuck,cao尼玛,法轮功,习近平';
    $filterWordMapCommon = filterWordsMap($filterWords);

    //var_dump($filterWordMapCommon);

    $filter = [];

    for($i = 0; $i< count($contents); $i++) {
        $filterWordMap = $filterWordMapCommon;
        $t = $temp = $contents[$i];
        $t = preg_split('//u', $t);
        array_shift($t);
        array_pop($t);
        $tlen = count($t);

        for ($j = 0; $j < $tlen; $j++) {

            if (!isset($filterWordMap[$t[$j]])) {
                continue;
            }

            $node = $filterWordMap[$t[$j]];
            $start = $j;
            for($x = $j + 1; $x < $tlen; $x++) {
                if (!isset($node[$t[$x]])) {
                    break;
                }

                if ($node[$t[$x]] == 1) {
                    $filter['文章_' . $i][] = [
                        'start' => $start,
                        'end' => $x,
                        'word' => mb_substr($temp, $start, $x - $start + 1)
                    ];

                    //避免同一个关键字被再次查找
                    unset($filterWordMap[$t[$start]]);

                    if (!$filterMany) {
                        break 2;
                    }
                }

                $node = $node[$t[$x]];
            }
        }

    }

    return $filter;

}

function testFindKeyPerformance()
{
    $article = [
        '我日习近平,haoeya,啥玩意啊，习大大，fuck ,就习近平啊，，好吧，wocao尼玛',
        '我我都算得上是的所得税，安静安静计算机是多少，中华人名共和国,我屁呀，fuck,我就不相信，我最爱法轮功',
        'baojia，安静安静计算机是多少，中华人名共和国,我屁呀，fuck,我就不相信，我最爱法轮功',
        'babab，安静安静计算机是多少，cao尼玛，fuck,我就不相信，我最爱法轮功',
        '啦啦，cao尼玛，我的法轮功啊,cao尼玛，fuck,我就不相信，我最爱法轮功',
        '舅舅，我贼，中华人名共和国,我屁呀，fuck,我就不相信，我最爱法轮功',
    ];
    $res = findKeyPerformance($article);
    var_dump($res);
}

testFindKeyPerformance();