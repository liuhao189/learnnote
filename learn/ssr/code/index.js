const Vue = require('vue');
const server = require('express')();
const render = require('vue-server-renderer').createRenderer({
    template: require('fs').readFileSync('./template.html', 'utf-8')
});

function cereateApp(context) {
    return new Vue({
        data: {
            url: context.url
        },
        template: `<div>Url is {{url}}</div>`
    });
}

server.get('*', (req, res) => {
    const app = cereateApp({ url: req.url });
    let context = {
        title: 'SSR Demo',
        meta: '<link rel="shortcut icon" href="https://img-oss.yunshanmeicai.com/cfile/25b31924f41fcbb14a4d2fc0ee9dc605.ico">'
    };
    render.renderToString(app, context).then(html => {
        res.end(html);
    }).catch(err => {
        res.status(500).end('Internal Server Error!s');
    });
});

server.listen(8080);