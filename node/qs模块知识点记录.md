## 该库主要是用于字符串的解析和字符串化库,增加了一些安全性

### 1 基本用法
    const qs = require('qs'')
    qs.parse('a=b') // {a:'b'}
    qs.parse('a=b&c=d') // {a:'b', c: 'd'}
    qs.parse('foo[a]=b') // {foo: {a: b}}
    qs.parse('foo[a]=b&foo[c]=d') // {foo: {a: 'b', c: 'd'}}
    qs.parse('a%5Bb%5D=c') // 'a[b]=c' {a:{b:c}}
    qs.parse('foo[bar][baz]=1') // {foo: {bar: {baz: 1}}}
    qs.parse('a[]=1&a[]=2') // {a: [1, 2]}
    qs.parse('a[]=1&a[]=2') // {a: [1, 2]}
    qs.parse('a[10]=1&a[9]=2') // {a: [10: 1, 9: 2]}
    qs.parse('a[][b]=c') // {a:[{b:c}]}
    
    
    
### 2 参数说明
    
    qs.parse('xxxx', {
        plainObjects: true,
        allowPrototypes: true,
        depth: 5,
        parameterLimit: 1000,
        ignoreQueryPrefix: true,
        delimiter: ';'
        allowDots: true, //弃用点表示法
        arrayLimit: 20 , //允许解析数组的最大索引
        parseArrays: false, //禁用数组解析
        comma: true,
        indices: false,
        arrayFormat: indices
    })
   
   
2.1 plainObjects 参数是为了控制有关键字时候，避免覆盖原型属性,以对象解析：
    
    qs.parse('obj[hasOwnProperty]=toby', {plainObjects: true})
    
2.2 allowPrototypes和"1"类似

2.3 qs最多只能解析5个子项, depth可以控制解析的深度

    const s = 'a[b][c][d][e][f][g][h][i]=j'
    { 
        a ：{ 
            b ：{ 
                c ：{ 
                    d ：{ 
                        e ：{ 
                            f ：{ 
                                ' [g] [h] [i] '：' j ' 
                            }
                        }
                    }
                }
            }
        }
    } ;
    
2.4 为了避免滥用参数，允许最多解析1000个参数，parameterLimit可以修改该默认

2.5 ignoreQueryPrefix允许解析的时候忽略前缀

    qs.parse('?a=b&c=d')
    
2.6 delimiter允许传递可选的分隔符

    qs.parse('a=b;c=d', {delimiter: ';'})  
当然分隔符也支持正则表达式

    qs.parse('a=b;c=d,e=f', {delimiter: /[;,]/})  
    
    
    
### 3 解析数组的特殊说明
qs解析数组的时候，会将稀疏数组压缩为仅保留其顺序的现有值

    qs.parse('a[14]=1&a[10]=2&a[17]=3');
    //{a:['2','1','3']}
qs会将数组中的索引指定为最大为20,超过该值转换为索引为键的对象

    qs.parse('a[14]=1&a[10]=2&a[21]=3');
    //{ a: { '10': '2', '14': '1', '21': '3' } }
    
    
有些人用逗号来连接数组:
    
    qs.parse('a=b,c', { comma: true })
    // {a: ['b', 'c']}
    
当数组被字符串化时，默认情况下被赋予显示索引
    
    qs.stringify({ a: ['b', 'c', 'd'] })
    //a[0]=b&a[1]=c&a[2]=d
   
可以设置indices来覆盖该设置 

    qs.stringify({ a: ['b', 'c', 'd'] }, { indices: false });
    // 'a=b&a=c&a=d'
    
也可以设置arrayFormat选项指定输出数组的格式

    qs.stringify({ a: ['b', 'c'] }, { arrayFormat: 'indices' })
    // 'a[0]=b&a[1]=c'
    qs.stringify({ a: ['b', 'c'] }, { arrayFormat: 'brackets' })
    // 'a[]=b&a[]=c'
    qs.stringify({ a: ['b', 'c'] }, { arrayFormat: 'repeat' })
    // 'a=b&a=c'
    qs.stringify({ a: ['b', 'c'] }, { arrayFormat: 'comma' })
    // 'a=b,c'
    
### 4 对象字符串化的说明

对象进行字符串化时，默认情况下使用括号表示法：
    
    qs.stringify({a: {b: {c: 'd', e: 'f'}}})
    // a[b][c]=d&a[b][e]=f
    
可以通过allowDots设置为使用点表示法：

    qs.stringify({ a: { b: { c: 'd', e: 'f' } } }, { allowDots: true });
    // 'a.b.c=d&a.b.e=f'
    
可以通过filter参数传递一个回调函数来控制哪些键被过滤

    function filterFunc(prefix, value) {
        if (prefix == 'b') {
            // Return an `undefined` value to omit a property.
            return;
        }
        if (prefix == 'e[f]') {
            return value.getTime();
        }
        if (prefix == 'e[g][0]') {
            return value * 2;
        }
        return value;
    }
    qs.stringify({ a: 'b', c: 'd', e: { f: new Date(123), g: [2] } }, { filter: filterFunc });
    // 'a=b&c=d&e[f]=123&e[g][0]=4'
    qs.stringify({ a: 'b', c: 'd', e: 'f' }, { filter: ['a', 'e'] });
    // 'a=b&e=f'
    qs.stringify({ a: ['b', 'c', 'd'], e: 'f' }, { filter: ['a', 0, 2] });
    // 'a[0]=b&a[2]=d'
        
    
    
    
    
    
    
    
    
    
    
    

