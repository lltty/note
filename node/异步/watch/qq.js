const Base = require('./base')

class QQ extends Base {
    constructor() {
        super()
    }

    send(msg) {
        console.log(`发送消息[QQ]:${msg}`)
    }
}

module.exports = new QQ()