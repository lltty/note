const wx = require('./wx')
const qq = require('./qq')
const kafka = require('./kafka')

class MysqlModel {

    constructor() {
        this.watchs = []
    }

    register(watch) {
        this.watchs.push(watch)
    }

    update() {
        console.log('我做更新操作')
        if (this.watchs.length > 0) {
            this.watchs.forEach(watch => {
                watch.send('数据库数据更新')
            })
        }
    }
}

const mysql = new MysqlModel()
mysql.register(wx)
mysql.register(qq)
mysql.register(kafka)

mysql.update()