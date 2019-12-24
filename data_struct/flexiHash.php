<?php
/**
 * Created by PhpStorm.
 * User: user
 * Date: 2018/9/19
 * Time: 下午4:23
 */

class CommonHash{

    protected $serverList;
    protected $isSorted = FALSE;


    public function __construct()
    {
    }

    public static function hash($key) {
        if (is_array($key)) {
            $key = json_encode($key);
        }
        $md5 = substr(md5($key), 0, 8);
        $seed = 31;

        $hash = 0;
        for ($i = 0; $i < 8; $i += 2) {
            $hash = $hash * $seed + ord($md5{$i});
        }

        return $hash & 0x07FFFFFFF;
    }

    public function checkServerList() {
        echo "<pre>";
        var_dump($this->serverList);
    }

    public function addServers($servers) {

        if (!is_array($servers)) {
            $servers = [$servers];
        }

        $res = FALSE;
        foreach($servers as $server) {
            $res = $this->addServer($server);
        }

        return $res;
    }

    public function removeServer($server) {
        $hash = $this->hash($server);
        if (isset($this->serverList[$hash])) {
            unset($this->serverList[$hash]);
        }
        $this->isSorted = FALSE;
        return TRUE;
    }

    public function getServerList() {
        return $this->serverList;
    }


    public function addServer($server) {
        $hash = $this->hash($server);
        if (!isset($this->serverList[$hash])) {
            $this->serverList[$hash] = $server;
        }
        $this->isSorted = FALSE;
        return TRUE;
    }

    public function findServer($key) {
        $hash = $this->hash($key);
        if (!$this->isSorted) {
            krsort($this->serverList, SORT_NUMERIC);
            $this->isSorted = TRUE;
        }

        foreach($this->serverList as $pos => $server) {
            if ($hash >= $pos) {
                return $server;
            }
        }

        return $this->serverList[array_keys($this->serverList)[count($this->serverList) - 1]];

    }

    public function test($key) {
        $hash = $this->hash($key);
        echo $hash,"\n";
    }
}

$h = new CommonHash();
$h->addServers(array(
    [
        'host' => '192.168.1.111',
        'port' => '3306'
    ],
    [
        'host' => '192.168.1.112',
        'port' => '3307'
    ],
    [
        'host' => '192.168.1.198',
        'port' => '3308'
    ],
    [
        'host' => '192.168.1.1',
        'port' => '3309'
    ],
    [
        'host' => '192.168.1.10',
        'port' => '3310'
    ],
    [
        'host' => '192.168.1.100',
        'port' => '3311'
    ],
    [
        'host' => '192.168.1.235',
        'port' => '3312'
    ],[
        'host' => '192.168.1.98',
        'port' => '3313'
    ],
));

$h->checkServerList();
exit;
echo "<pre>";
var_dump($h->findServer('hz_name'));
var_dump($h->findServer('hz_age'));
var_dump($h->findServer('hz_addr'));
var_dump($h->findServer('a'));
var_dump($h->findServer('A'));

