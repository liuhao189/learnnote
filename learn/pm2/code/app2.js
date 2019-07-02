const Koa = require("koa")
const app = new Koa();

app.use(async ctx => {
    ctx.body = "Hello, I am app2!"
})

app.listen(3100)