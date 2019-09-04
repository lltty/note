/*
 * 继承该类需要实现send方法
 */

class SendBase {
    constructor() {
        if (!this.hasImplement(['send'])) {
            throw new Error('没有继承send方法')
        }
    }

    hasImplement(func) {

        if (func && typeof func == 'object' && func.length > 0) {

            func.forEach(fun => {
                if (!(typeof this[fun] == 'function')) {
                    return false
                }
            })
            return true
        }
        return false
    }
}

module.exports = SendBase