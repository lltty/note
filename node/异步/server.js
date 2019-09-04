const http = require('http')
const fs = require('fs')

function sleep(sec) {
    const end = Date.now() + sec * 1000
    while (Date.now() < end) {
    }
    return
}

const server = http.createServer((req, res) => {
    const start = Date.now()
    sleep(3)
    const end = Date.now()
    fs.writeFileSync('./res.log', `${start} -> ${end}\n`, {flag: 'a+'})
    res.end(`${start} -> ${end}`)
})

//ab测试:ab -c 2 -n 5 http://localhost:8081/
server.listen(8083)