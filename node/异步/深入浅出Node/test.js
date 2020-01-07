var fs = require('fs')

const f1 = fs.readFile('./t/m1.md', function(err, file) {
    if (err != undefined) {
        console.log('报错了', err)
        return
    }
    console.log("读取文件完成", file.toString())
})

console.log("发起读取文件")