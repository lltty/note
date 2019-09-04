const Base = require('./base')

class WeChat extends Base {
    constructor() {
        super()
    }

    send(msg) {
        console.log(`发送消息[微信消息]:${msg}`)
    }
}

module.exports = new WeChat()