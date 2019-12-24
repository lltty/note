<?php

/*
 *  这是一个hash
 *  每个hash的元素又是一个链表
 */

class HashNode{

    public $key;
    public $val;
    public $nextNode;

    public function __construct($key, $val, $nextNode)
    {
        $this->key = $key;
        $this->val = $val;
        $this->nextNode = $nextNode;
    }

}

class Hash{

    private $size = 5;
    private $buckets;


    public function __construct()
    {
        $this->buckets = new SplFixedArray($this->size);
    }

    public function hashStr($str) {
        $len = strlen($str);
        $hash = '';
        for($i = 0; $i < $len; $i++) {
            $hash += ord($str{$i});
        }

        return $hash % $this->size;
    }

    public function insert($key, $val)
    {
        $hash = $this->hashStr($key);

        if (isset($this->buckets[$hash])) {
            $this->buckets[$hash] = new HashNode($key, $val, $this->buckets[$hash]);
        } else {
            $this->buckets[$hash] = new HashNode($key, $val, NULL);
        }
    }

    public function get($key) {
        $hash = $this->hashStr($key);
        $current = $this->buckets[$hash];

        while($current) {
            if ($current->key == $key) {
                return $current->val;
            }
            $current = $current->nextNode;
        }

        return NULL;
    }

}

$h = new Hash();
$h->insert('hz_name', 'h_lihongwei');
$h->insert('media_name', 'm_lihongwei');
$h->insert('adx_name', 'adx_lihongwei');

echo $h->get('hz_name');