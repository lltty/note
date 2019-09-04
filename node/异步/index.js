async function asyncFunc() {
    console.log('NO.2')
    await new Promise((res, rej) => {
        console.log('No.3')

        //这是一个异步方法，会放到next tick执行
        setTimeout(() => {
            console.log('No.6')
            res()
        }, 0)

        console.log('No.4')
    })
    console.log('No.7')
    return 'No.8'
}

console.log('No.1')
let asyncPromise = asyncFunc()
console.log('No.5')
asyncPromise.then(value => {
    console.log(value)
})