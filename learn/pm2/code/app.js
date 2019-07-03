const Koa = require("koa")
const fs = require("fs");
const app = new Koa();
const fileName = new Date().getTime() + '.txt'
fs.writeFileSync(fileName, "I am log file!");

app.use(async ctx => {
    ctx.response.headers["Content-Type"] = "text/html";
    ctx.body = "Hello PM2!" + `From ${process.pid}` + ` Argv:${process.argv.join(',')}` + `Env:${JSON.stringify(process.env)}`;
    ctx.body += '<h1>I am html!</h1>'
})

process.on("SIGINT", () => {
    console.log('SIGINT')
    fs.unlinkSync(fileName);
    process.exit(1);
});

app.listen(3000)