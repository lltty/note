1 lua数据转化成redis有一条额外规则,true/false没有对应的redis类型

``
eval "return true" 0 #(integer) 1
``

``
eval "return fasle" 0 #(nil)
``
2 lua中整数和浮点数是没有区别的,小数部分是会被忽略的,所以如果需要返回浮点数，请返回字符串

3 lua脚本是原子性的，所以可以用lua脚本封装一些原子操作

4 可以使用script load 加载一个脚本返回sha值,然后可使用类似的函数调用

``
eval "return KEYS[1] + KEYS[2]" 2 10 20 #30
``

使用函数的方式：
``
script load "return KEYS[1]+KEYS[2]" #7b23d2a5829679ac50baf7c8e105904a3e9e69bb
evalsha 7b23d2a5829679ac50baf7c8e105904a3e9e69bb 10 20 #30
``
``
script flush #删除所有的lua脚本
script exists #判断脚本是否存在
``

应该在避免在evalsha内部再使用evalsha，因为需要在代码的顶端判断使用的脚本是否存在,不存在的话,需要script load下，才可以使用。

    