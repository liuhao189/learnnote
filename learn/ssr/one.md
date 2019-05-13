# Vue SSR

## 安装

```bash
npm install vue vue-server-renderer --save
```
推荐使用Nodejs版本6+；vue-server—renderer和vue必须匹配版本；vue-server-render依赖一些Node.js原生模块，只能在Nodejs中使用。

## 与服务器集成

```js
const Vue = require('vue');
const server = require('express')();
const render = require('vue-server-renderer').createRenderer();

server.get('*', (req, res) => {
    const app = new Vue({
        data: {
            url: req.url
        },
        template: `<div>Url is {{url}}</div>`
    });

    render.renderToString(app, (err, html) => {
        if (err) {
            res.status(500).end('Internal Server Error');
            return;
        }
        res.end(`<!DOCTYPE html>
                 <html lang="en">
                    <head><title>Url</title></head>
                    <body>${html}</body>
                 </html>`);
    });
});

server.listen(8080);
```

## 使用一个页面模板

当你在渲染Vue应用程序时，renderer只从应用程序生成HTML标记。可以直接在创建renderer时提供一个页面模板。
<!--vue-ssr-outlet-->注释指定应用程序HTML标记注入的地方。

```html
<body>
    <!--vue-ssr-outlet-->
</body>
```
```js
const render = require('vue-server-renderer').createRenderer({
    template: require('fs').readFileSync('./template.html', 'utf-8')
})
//2.5.0+中，没有传入回调函数，则会返回Promise
render.renderToString(app).then(html => {
    res.end(html);
}).catch(err => {
    res.status(500).end('Internal Server Error!s');
});
```

## 插值模板

模板还支持简单插值。

```html
<title>{{title}}</title>
<!--三花括号进行html不转义插值-->
{{{meta}}}
```
可以通过传入一个渲染上下文对象，作为renderToString函数的第二个参数，来提供插值数据。
也可以与Vue应用程序实例共享context对象，允许模板插值中的组件动态地注册数据。
```js
render.renderToString(app, context)
```
# 编写通用代码

运行在服务器和客户端的代码，由于用例和平台API的差异，运行在不同环境中时，代码将不会完全相同。

## 服务器上的数据响应

在纯客户端应用程序中，每个用户会在他们各自的浏览器中使用新的应用程序实例。对于服务器端渲染，我们也希望如此，每个请求都是全新的，独立的应用程序实例，以便不会有交叉请求造成的状态污染。

将在服务器上预取数据，开始渲染时，我们的应用程序就已经解析完成其状态。将数据进行响应式的过程在服务器上是多余的，默认情况下禁用。

## 组件生命周期钩子函数

没有动态更新，所有的生命周期钩子函数中，只有beforeCreate和create会在服务器端渲染过程中被调用。其它生命周期钩子函数中的代码beforeMount，mounted只会在客户端执行。

避免beforeCreate和create生命周期时产生全局副作用的代码。使用setInterval。副作用代码移动到beforeMount或mounted生命周期中。

## 访问特定平台的API

通用代码不可接受特定平台的API，直接使用window，document，仅浏览器可用的全局变量，在Nodejs执行时会抛出错误。

对于共享于服务器和客户端，但用于不同平台的API的任务，建议将平台特定实现包含在通用API中，或者为你执行此操作的类库。axios是一个HTTP库，可以向服务器和客户端都暴露相同的API。

对于纯浏览器可用的API，通常方式是，在纯客户端的生命周期钩子函数中惰性访问。

## 自定义指令

大多数自定义指令直接操作DOM，因此会在服务器渲染过程中导致错误。有两者方法可以解决这个问题：推荐使用组件作为抽象机制，并运行在虚拟DOM层；如果你有一个自定义指令，但是不是很容易替换为组件，可以在创建服务器renderer时，使用directives选项所提供的服务器版本。

# 源码结构

## 避免状态单例

当编写纯客户端代码时，习惯在每次在新的上下文中对代码进行取值。Nodejs是一个长期运行的进程，如果创建一个单例对象，它将在每个传入的请求之间共享。

我们不应该直接创建一个应用程序实例，而是应该暴露一个可以重复执行的工程函数。同样的规则也适用于router，store和event bus实例，不应该直接从模块导出并将其导入到应用程序中，而是需要在createApp中创建一个新的实例，并从根Vue实例注入。 

```js
function cereateApp(context) {
    return new Vue({
        data: {
            url: context.url
        },
        template: `<div>Url is {{url}}</div>`
    });
}
```

## 介绍构建步骤

需要使用Webpack来打包我们的Vue应用程序，事实上，我们可能需要在服务器上使用webpack打包vue应用程序。

通常Vue应用程序是由webpack和vue-loader构建，并且许多webpack特定功能不能直接Node.js中运行。

尽管Nodejs最新版本能够完全支持ES205特性，还是需要转译客户端代码以适应老版浏览器。

基本的步骤是，对于客户端应用程序和服务器应用程序，都要使用webpack打包，服务器需要服务器bundle用于服务器端渲染，而客户端bundle会发送给浏览器，用于混合静态标记。


## 使用webpack的源码结构

正在使用webpack来处理服务器和客户端的应用程序，大部分源码可以使用通用方式编写，可以使用webpack支持的所有功能。

```bash
src
--components
  --Foo.vue
  --Bar.vue
  --Baz.vue
--App.vue
--app.js
--entry-client.js
--entry-server.js 
```
app.js是我们应用程序的通用entry，在纯客户端应用程序中，将在此文件中创建根Vue实例，并直接挂在到DOM。但是，对于服务器渲染SSR，责任转译到纯客户端entry文件，app.js简单地使用export导出一个createApp函数。

entry-client.js只需创建应用程序，并且将其挂在到DOM中。

entry-server.js每次渲染中重复调研此函数。除了创建和返回应用程序实例之外，不会做太多事情。稍后将在此执行服务器路由匹配和数据预趣逻辑。

```js
import Vue from 'vue';
import App from './App.vue';

export function createApp() {
    const app = new Vue({
        render: (h) => h(App)
    });
    return { app };
}
//entry-client.js
import { createApp } from './app';

const { app } = createApp();

app.$mount('#app');
//entry-server.js
import { createApp } from './app';

export default context => {
    const { app } = createApp();
    return app;
}
```

# 路由和代码分割

使用官方提供的vue-router，首先创建一个文件，在其中创建router，需要给每个请求创建一个新的router实例。文件导出一个createRouter函数。

```js
import Vue from "vue";
import Router from "vue-router";

Vue.use(Router);

export function createRouter() {
    return new Router({
        mode: 'history',
        routes: [{
            path: 'one',
            component: import('./components/one.vue')
        }, {
            path: 'two',
            component: import('./components/two.vue')
        }]
    })
}
//
import { createApp } from './app';

export default context => {
    return new Promise((resolve, reject) => {
        const { app, router } = createApp();
        router.push(context.url);
        router.onReady(() => {
            const matchComponents = router.getMatchedComponents();
            if (!matchComponents) {
                return reject({ code: 404 });
            }
            resolve(app);
        }, reject);
    });
}
```
## 代码分割

应用程序的代码分割或惰性加载，有助于减少浏览器在初始渲染中下载的资源体积，可以极大地改善大体积bundle的可交互时间。
在Vue2.5以下的版本中，服务器端渲染时异步组件只能用在路由组件上，2.5+的版本中，得益于核心算法的升级，异步组件现在可以用在应用的任何地方。

