<?php

$file1 = 'a.txt';           //随便拷一个图片作为测试用
$file2 = 'b';         //生成的二进制流保存在这个文件里
$file3 = 'c.txt';           //由二进制流还原成的文件

$size = filesize($file1);

echo '文件大小为：' . $size, "\n";
echo "\n转化为二进制 ...\n";

$content = file_get_contents($file1);
$content = bstr2bin($content);


$fp = fopen($file2, 'w');
fwrite($fp, $content);
fclose($fp);

$size2 = filesize($file2);

echo '转化成二进制后文件大小为：' . $size2. "\n";

$content = bin2bstr($content);

$fp = fopen($file3, 'w');
fwrite($fp, $content);
fclose($fp);


function bin2bstr($input)
// Convert a binary expression (e.g., "100111") into a binary-string
{
    echo "解码~~~~~~~~~~~~~~~~~\n";
    if (!is_string($input)) return null; // Sanity check
    // Pack into a string
    $input = str_split($input, 4);
    $str = '';
    foreach ($input as $v) {
        $str .= base_convert($v, 2, 16);
    }

    echo "待解码:",$str,"\n";

    $str = pack('H*', $str);

    return $str;
}

function bstr2bin($input)
{
    if (!is_string($input)) return null; // Sanity check

    echo "输入：",$input,"\n";
    // Unpack as a hexadecimal string
    $value = unpack('H*', $input);
    echo "输出：",
    var_dump($value);

    // Output binary representation
    $value = str_split($value[1], 1);
    $bin = '';
    foreach ($value as $v) {
        $b = str_pad(base_convert($v, 16, 2), 4, '0', STR_PAD_LEFT);

        $bin .= $b;
    }

    return $bin;
}
