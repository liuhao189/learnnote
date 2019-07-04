const Koa = require("koa")
const fs = require("fs");
const app = new Koa();
const fileName = new Date().getTime() + '.txt'
fs.writeFileSync(fileName, "I am log file!");

app.use(async ctx => {
    ctx.response.headers["Content-Type"] = "text/html";
    ctx.body = "Hello PM2!" + `From ${process.pid}`;
    ctx.body += '<h1>I am html!</h1>'
})

process.on("SIGINT", () => {
    console.log('SIGINT')
    fs.unlinkSync(fileName);
    process.exit(0);
});

setTimeout(() => {
    process.send('ready');
}, 10000);

app.listen(3000)