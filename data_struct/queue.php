<?php

/*
 *  这是一个循环队列
 *  保证先进先出，且循环保存
 */
class CircleQueue
{
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

        return true;
    }

    public function deQueue() {
        if ($this->isEmpty()) return FALSE;

        if ($this->head == $this->tail) {
            $this->head = -1;
            $this->tail = -1;
            return true;
        }

        $this->head = ($this->head + 1) % $this->size;
        return true;

    }

    public function isEmpty() {

        return $this->head == -1;
    }

    public function isFull() {

        $head = $this->head;
        $tail = $this->tail;

        return ($tail + 1) % $this->size == $head;
    }

    public function front() {
        return $this->data[$this->head];
    }

    public function rear() {
        return $this->data[$this->tail];
    }

}

$circleQueue = new CircleQueue(3);
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







