<?php
$input = [
    4,
    'A',
    'B',
    'C',
    'D',
    'A-B 1:1',
    'A-C 2:2',
    'A-D 1:0',
    'B-C 0:3',
    'B-D 0:3',
    'C-D 0:3',
];


class Team {
    public $name;
    public $grade; #积分
    public $realWin; #净进球
    public $in; #进球数

    public function __construct($name, $grade = 0, $realWin = 0, $in = 0) {
        $this->name = $name;
        $this->grade = $grade;
        $this->realWin = $realWin;
        $this->in = $in;
    }
}

class ArrayUtil {

    public static function sortAryByKey($team1, $team2) {

        if ($team1['realWin'] == $team2['realWin']) {
            return $team1['in'] < $team1['in'];
        }

        return $team1['realWin'] < $team2['realWin'];

    }
}

function output($input) {
    $n = $input[0];
    $index = 1;
    $footballTeam = [];
    for ($i = 1; $i <= $n; $i++) {
        $name = $input[$i];
        $footballTeam[$input[$i]] = [
            'name' => $name,
            'grade' => 0,
            'realWin' => 0,
            'in' => 0
        ];
    }
    $index += $n;

    for ($i = $index; $i < count($input); $i++) {

        $params = explode(" ", $input[$i]);
        $teams = explode("-", $params[0]);
        $score = explode(":", $params[1]);

        $team1 = $teams[0];
        $in1 = $score[0];

        $team2 = $teams[1];
        $in2 = $score[1];

        //双方进球数增加
        $footballTeam[$team1]['in'] += $in1;
        $footballTeam[$team2]['in'] += $in2;

        //核算积分
        $diff = $in1 - $in2;
        if ($diff == 0) {
            $footballTeam[$team1]['grade'] += 1;
            $footballTeam[$team2]['grade'] += 1;
        } else if ($diff > 0) {
            $footballTeam[$team1]['grade'] += 3;
            $footballTeam[$team1]['realWin'] += $diff;
        } else {
            $footballTeam[$team2]['grade'] += 3;
            $footballTeam[$team2]['realWin'] += -$diff;
        }
    }

    //按照既定规则对球队进行排序
    $footballTeam = array_values($footballTeam);
    usort($footballTeam, 'ArrayUtil::sortAryByKey');

    for ($i = 0; $i < $n/2; $i++) {
        var_dump($footballTeam[$i]);
    }

}