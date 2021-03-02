<?php
/**
 * Created by PhpStorm.
 * User: user
 * Date: 2019/3/9
 * Time: 11:42 AM
 */


class FlexiHash
{

    private $serverList = [];
    private $isSort = FALSE;

    public function __construct()
    {

    }

    public function hash($key)
    {

        if (is_array($key)) {
            $key = json_encode($key);
        }
        $key = md5($key);
        $hash = 0;
        $seed = 33;
        for ($i = 0; $i < 32; $i++) {
            $hash = ($hash * $seed + ord($key{$i})) & 0x7FFFFFFF;
        }

        return $hash;
    }

    public function addServer($server)
    {

        $hash = $this->hash($server);
        if (!isset($this->serverList[$hash])) {
            $this->serverList[$hash] = $server;
        }

        $this->isSort = FALSE;
        return;
    }

    public function addServers($servers)
    {

        if (!is_array($servers)) {
            $servers = [$servers];
        }

        foreach ($servers as $server) {
            $this->addServer($server);
        }

        return TRUE;

    }

    public function sortServer()
    {
        $this->serverList = krsort($this->serverList);
        $this->isSort = TRUE;
    }

    public function getServer($key)
    {


        if (!$this->isSort) {
            $this->sortServer();
        }


        $hash = $this->hash($key);

        foreach ($this->serverList as $serverHash => $server) {
            if ($hash > $serverHash) {
                return $server;
            }
        }

        return $this->serverList[count($this->serverList) - 1];

    }

    public function removeServer($server)
    {
        $hash = $this->hash($server);
        if (!isset($this->serverList[$hash])) {
            unset($this->serverList[$hash]);
        }

        return TRUE;
    }

}

//无极限分类
$data = [
    [
        "name" => "toby1",
        "age" => "123",
        "id" => 1,
    ],
    [
        "name" => "toby2",
        "age" => "123",
        "id" => 2,
    ],
    [
        "name" => "toby3",
        "age" => "123",
        "id" => 3,
    ],
    [
        "name" => "toby1_children1",
        "age" => "123",
        "id" => 4,
        "pid" => 1
    ],
    [
        "name" => "toby1_children2",
        "age" => "123",
        "id" => 5,
        "pid" => 1
    ],
    [
        "name" => "toby1_children2_children1",
        "age" => "123",
        "id" => 6,
        "pid" => 5
    ],
    [
        "name" => "toby2_children1",
        "age" => "123",
        "id" => 7,
        "pid" => 2
    ],
    [
        "name" => "toby2_children2",
        "age" => "123",
        "id" => 8,
        "pid" => 2
    ],
];

//无极限分类实现函数
function infiniteList($data, $pk = 'id', $pid = 'pid', $children = 'children')
{

    if (!is_array($data)) {
        return $data;
    }

    $ref = [];
    foreach ($data as $key => $d) {
        $ref[$d[$pk]] = &$data[$key];
    }

    $root = 0;
    $list = [];
    foreach ($data as $key => $d) {
        $parentId = $d[$pid] ?? 0;
        if ($parentId == $root) {
            $list[] = &$data[$key];
        } else {
            if (isset($ref[$parentId])) {
                $parent = &$ref[$parentId];
                $parent[$children][] = &$data[$key];
            }
        }
    }

    return $list;
}

//找出数组里组合等于某个数字的元素
function findTargetCombine($source, $target)
{
    $source = array_unique($source);
    $source = array_flip($source);
    $eles = [];
    foreach ($source as $key => $val) {
        if (isset($source[$target - $key])) {
            $eles[] = [$key, $target - $key];
        }
    }

    return $eles;
}

//计算场次总分
function sumRoundScore($input)
{

    $index = 0;
    $filterScores = [];
    foreach ($input as $key => $score) {
        $score = strtoupper($score);
        switch ($score) {
            case "C":
                $index--;
                unset($filterScores[$index]);
                break;
            case "D":
                $filterScores[$index] = $filterScores[$index - 1] * 2;
                $index++;
                break;
            case "+":
                $filterScores[$index] = $filterScores[$index - 1] + $filterScores[$index - 2];
                $index++;
                break;
            default:
                $filterScores[$index] = $input[$key];
                $index++;
        }
    }

    return $filterScores;
}


function _sumRoundScore()
{
    $input = [5, 2, "C", "D", "+"];
    //c 上一场分数是无效的
    //d 这场分数是有效的上一场分数的2倍
    //+ 这场分数是有效的上两场分数的总和
    var_dump(sumRoundScore($input));
}

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

function wordsToMap ($words) {

    $words = str_replace('，', ',', $words);
    $words = preg_replace('/^\s*,\s*|\s*,\s*$/', '', $words);
    $words = preg_split('/\s*,\s*/', $words);

    $tree = [];
    foreach($words as $word) {

        $word = preg_split('//u', $word);
        array_shift($word);
        array_pop($word);

        $node = [];
        $len = count($word);
        for($i = $len -1 ; $i >= 0; $i--) {
            if ($i == $len - 1) {
                $node = [
                    $word[$i] => 1
                ];
            } else {
                $node[$word[$i]] = $node;
                unset($node[$word[$i+1]]);
            }
        }

        $tree[$word[0]] = $node[$word[0]];
    }

    return $tree;

}

function findFilterWordFromContent($words, $contents, $findMany = FALSE) {
    $filterWordMap = wordsToMap($words);
    if (is_string($contents)) {
        $contents = [$contents];
    }

    $filter = [];
    $contentLen = count($contents);

    for ($i = 0; $i < $contentLen; $i++) {
        $t = $contents[$i];
        $contentWords = preg_split('//u', $contents[$i]);
        array_shift($contentWords);
        array_pop($contentWords);
        $wordLen = count($contentWords);
        $filterWordMapTemp = $filterWordMap;

        for($j = 0; $j < $wordLen; $j++) {

            if (!isset($filterWordMapTemp[$contentWords[$j]])) {
                continue;
            }

            $start = $j;
            $node = $filterWordMapTemp[$contentWords[$j]];
            for ($x = $j + 1; $x < $wordLen; $x++) {

                if (!isset($node[$contentWords[$x]])) {
                    break;
                }

                if ($node[$contentWords[$x]] == 1) {

                    $filter[] = [
                        "文章_{$i}" => [
                            'start' => $start,
                            'end' => $x,
                            'word' => mb_substr($t, $start, $x - $start + 1)
                        ]
                    ];

                    unset($filterWordMapTemp[$contentWords[$x]]);
                    if (!$findMany) {
                        break 2;
                    }
                }

                $node = $node[$contentWords[$x]];

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
    $filterWords = '中华人名共和国,我贼,fuck,cao尼玛,法轮功,习近平';
    $res = findFilterWordFromContent($filterWords, $article, TRUE);
    var_dump($res);
}


//实现PHP的hashTable
class Bucket {

    public $key;
    public $val;
    public $nextNode;
    public $prevNode;

    public function __construct($key, $val, $nextNode)
    {
        $this->key = $key;
        $this->val = $val;
        $this->nextNode = $nextNode;
        if ($nextNode != NULL) {
            $nextNode->prevNode = $this->nextNode;
        }
    }
}


class MYHash {


    private $size;
    private $buckets;

    public function __construct ($size = 20) {
        $this->size = $size;
        $this->buckets = new SplFixedArray($size);
    }

    private function hash($str) {
        $len = strlen($str);
        $hash = 0;
        for($i = 0; $i < $len; $i++) {
            $hash += ord($str{$i});
        }

        return $hash % $this->size;
    }

    public function insert($key, $val) {
        $hash = $this->hash($key);
        if (isset($this->buckets[$hash])) {
            $this->buckets[$hash] = new Bucket($key, $val, $this->buckets[$hash]);
        } else {
            $this->buckets[$hash] = new Bucket($key, $val, NULL);
        }
    }

    /*
     * 这里其实是有问题的，因为insert的时候没有判断值是否存在，所以导致相同的键值可能回被存储为多个数据
     * 所以这里才会有getAll,理论上hashTable是不允许存储重复键值的
     */
    public function get($key, $getAll = FALSE) {
        $hash = $this->hash($key);
        $res = [];
        if (isset($this->buckets[$hash])) {
            $current = $this->buckets[$hash];
            while ($current) {
                if ($current->key == $key) {
                    if ($getAll) {
                        $res[] = $current->val;
                    } else {
                        return $current->val;
                    }
                }
                $current = $current->nextNode;
            }
        }

        if ($getAll) return $res;
        return NULL;
    }

}

function testHashTable() {
    $h = new MYHash();
    $h->insert("toby", "101");
    $h->insert("toby", 102);
    var_dump($h->get("toby", TRUE));
}

testHashTable();


class CircleQueue {

    public $size;
    public $data = [];
    public $head = -1;
    public $tail = -1;

    public function __construct($size) {
        $this->size = $size;
    }

    public function enQueue($val) {
        if ($this->isFull()) return FALSE;
        if ($this->isEmpty()) $this->head = 0;
        $this->tail = ($this->tail + 1) % $this->size;
        $this->data[$this->tail] = $val;
    }

    public function deQueue() {

        if ($this->isEmpty()) return FALSE;
        if ($this->head == $this->tail) {
            $this->head = -1;
            $this->tail = -1;
            return TRUE;
        }

        $this->head = ($this->head + 1 ) % $this->size;
    }

    public function isFull() {
        return ($this->tail + 1 ) % $this->size == $this->head;
    }

    public function isEmpty() {
        return $this->head == -1;
    }

    public function front() {
        $t = $this->data[$this->head];
        unset($this->data[$this->head]);
        return $t;
    }

    public function rear() {
        return $this->data[$this->tail];
    }

}

function testCircleQueue() {
    $circleQueue = new CircleQueue(5);
    $circleQueue->enQueue('a');
    $circleQueue->enQueue('b');
    $circleQueue->enQueue('c');
    $circleQueue->enQueue('d');
    $circleQueue->enQueue('e');

    //队列，取数据，先进先出
    while(!$circleQueue->isEmpty()) {
        echo $circleQueue->front(),"\n";
        $circleQueue->deQueue();
    }

    $circleQueue->enQueue('a1');
    $circleQueue->enQueue('b1');
    $circleQueue->enQueue('c1');
    $circleQueue->enQueue('d1');
    $circleQueue->enQueue('e1');

    //队列，取数据，先进先出
    while(!$circleQueue->isEmpty()) {
        echo $circleQueue->front(),"\n";
        $circleQueue->deQueue();
    }
}





















