<?php

$a = [
    "我的" => 30,
    "你的啊啊啊" => 10
];

asort($a);
echo "<pre>";
var_dump($a);

exit;

function getTarget($num, $target) {
	
	$temp = array_flip($num);

	$len = count($num);
	$output = [];
	for($i = 0; $i < $len; $i++) {
		if(isset($temp[$target - $num[$i]])) {
			$output[] = [
				$i, $temp[$target - $num[$i]],
				$num[$i], $target - $num[$i]
			];
		}
	}

	return $output;
}


function testGetTarget() {
	$num = [1, 10, 13, 9, 2, 8, 7];
	var_dump(getTarget($num, 9));
}

function sumRoundScore($input) {

	$len = count($input);
	$index = 0;
	$t = [];
	for ($i = 0; $i< $len; $i++) {
		$s = strtoupper($input[$i]);
		switch($s) {
			case 'C':
				$index--;
				unset($t[$index]);
				break;
			case 'D':
				$t[$index] = $t[$index - 1] * 2;
				$index++;
				break;
			case '+':
				$t[$index] = $t[$index - 1] + $t[$index - 2];
				$index++;
				break;
			default:
				$t[$index] = $input[$i];
				$index++;

		}

	}

	return $t;
}


function _sumRoundScore(){
	$input = [5, 2,"C","D","+"];
	//c 上一场分数是无效的
	//d 这场分数是有效的上一场分数的2倍
	//+ 这场分数是有效的上两场分数的总和
	var_dump(sumRoundScore($input));
}


//二分递归查找
function binarySearch($source, $target, $start, $end) {
	
	//二分查找的时候要注意这里必须用floor(向下取整)，如果用ceil(向上取整),会跳过边界值
	$middle = floor(($start + $end) / 2 );

	if ($target > $source[$middle]) {
		$start = $middle + 1;

		if ($start > $end) {
			echo $start, '->', $end, "\n";
			return FALSE;
		}
		return binarySearch($source, $target, $start, $end);
	} else if ($target < $source[$middle]) {
		$end = $middle - 1;
		if ($end < 0) {
			return FALSE;
		}
		return binarySearch($source, $target, $start, $end);

	} else {
		return $middle;
	}
}

//二分循环查找
function binarySearchW($source, $target) {
	sort($source);
	$start = 0;
	$end = count($source);

	while (true) {
		$middle = floor(($start + $end) / 2);
		if ($source[$middle] > $target) {
			$end = $middle - 1;
			if ($end < 0) return FALSE;
		} else if($source[$middle] < $target) {
			$start = $middle + 1;
			if ($start > $end) return FALSE;
		} else {
			return $middle;
		}
	}
}

function testBinarySearch() {
	$source = [9, 10, 31, 2, 4, 19, 3, 14, 20, 17];
	sort($source);
	var_dump(binarySearch($source, 17, 0, count($source)));
	var_dump(binarySearchW($source, 17));
}

/*
 * 无极限分类
 */
function listToTree($list, $pk='id', $pid = 'pid', $child = '_child', $root = 0) {
    // 创建Tree
    $tree = array();
    if(is_array($list)) {
        // 创建基于主键的数组引用
        $refer = array();
        foreach ($list as $key => $data) {
            $refer[$data[$pk]] = &$list[$key];
        }

        foreach ($list as $key => $data) {
            // 判断是否存在parent
            $parentId =  $data[$pid];
            if ($root == $parentId) {
                $tree[] = &$list[$key];
            }else{
                if (isset($refer[$parentId])) {
                    $parent = &$refer[$parentId];
                    $parent[$child][] = &$list[$key];
                }
            }
        }
    }

    return $tree;
}


function testListToTree() {
	$list = [
	    [
	        'id' => 6,
	        'pid' => 2,
	        'name' => 'b2'
	    ],
	    [
	        'id' => 7,
	        'pid' => 6,
	        'name' => 'b21'
	    ],
	    [
	        'id' => 8,
	        'pid' => 6,
	        'name' => 'b22'
	    ],
		[
			'id' => 1,
			'pid' => 0,
			'name' => 'a'
		],
		[
			'id' => 2,
			'pid' => 0,
			'name' => 'b'
		],
		[
			'id' => 3,
			'pid' => 0,
			'name' => 'c'
		],
		[
			'id' => 4,
			'pid' => 0,
			'name' => 'd'
		],
		[
			'id' => 5,
			'pid' => 2,
			'name' => 'b1'
		]
	];

	$tree = list_to_tree($list);
}
