const Base = require('./base')

class Kafka extends Base {
    constructor() {
        super()
    }

    send(msg) {
        console.log(`发送消息[KAFKA]:${msg}`)
    }
}

module.exports = new Kafka()