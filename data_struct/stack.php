<?php

/*
 *  这是一个循环堆栈
 *  保证先进先出，且循环保存
 */
class CircleStack{
	public $head = -1;
	public $tail = -1;
	public $data = [];
	public $size = 0;

	public function __construct($size) {

		$this->size = $size;
		$this->head = $size;
		$this->tail = $size;
	}

	public function enStack($val) {
		if ($this->isFull()) return FALSE;

        if ($this->isEmpty()) $this->tail = $this->size - 1;

        if ($this->head > 0) {
        	$this->head -= 1;
        } else {
        	$this->head += $this->size - 1;
        }
        $this->data[$this->head] = $val;

        return true;
	}

	public function deStack() {
		if ($this->isEmpty()) return FALSE;

		if ($this->head == $this->tail) {
			$this->head = $this->size;
			$this->tail = $this->size;
			return true;
		}

		$this->head = $this->head + 1;

	}

	public function isEmpty() {
		return $this->head == $this->size;
	}

	public function isFull() {
		return $this->head == 0;
	}

	public function front() {
		return $this->data[$this->head];
	}

	public function rear() {
		return $this->data[$this->tail];
	}


}

$circleStack = new CircleStack(5);
$circleStack->enStack('a');
$circleStack->enStack('b');
$circleStack->enStack('c');
$circleStack->enStack('d');
$circleStack->enStack('e');

//栈，取数据，先进后出
while(!$circleStack->isEmpty()) {
	echo $circleStack->front(),"\n";
	$circleStack->deStack();
}



