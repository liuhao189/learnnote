# Vue SSR

## 什么是服务器渲染

Vuejs是构建客户端应用程序的框架，默认情况下，可以在浏览器中输出Vue组件，进行生产DOM和操作DOM。然而，也可以将同一个组件渲染为服务器端的HTML字符串，将它们直接发送到浏览器。最后将这些静态标记激活为客户端上完全可交互的应用程序。

服务器渲染Vue.js应用程序也可以被认为是同构或通用的，因为应用程序的大部分代码都可以在服务器和客户端上运行。

## 为什么使用服务器渲染

与传统SPA相比，服务器渲染SSR的优势主要在于：

1、更好的SEO，由于搜索爬虫抓取工具可以直接查看完全渲染的页面。异步获取内容，可能需要服务器端渲染解决此问题。

2、更快的内容到达时间，也别是对于缓慢的网络情况或运行缓慢的设备。无需等待所有的JS都完成下载并之执行，你的用户将会更快速地看到完整渲染的页面。

权衡之处：

1、开发条件所限，浏览器特定的代码，只能在某些生命周期钩子函数中使用，一些外部扩展库可能需要特殊处理，才能在服务器渲染应用程序中运行。

2、设计构建设置和部署的更多要求，与可以部署在任何静态文件服务器上的完全静态单页应用程序不同，服务器渲染应用程序，需要处于Nodejs Server运行环境。

3、更多的服务器负载，在Nodejs中渲染完整的应用程序，比仅仅提供静态文件的server更占用CPU资源，预料在高流量下使用，请准备相应的服务器负载，并明智地采用缓存策略。

是否真的需要SSR，取决于内容到达时间对应用程序的重要程度。

## 服务器端渲染 VS 预渲染

如果你调研服务器渲染只是用来改善少数营销页面的SEO，那么你可能需要预渲染。使用预渲染，在构建时简单地生成针对特定路由的静态HTML文件。特点是设置预渲染更简单，并可以将你的前端作为一个完全静态的站点。

可以使用prerender-spa-plugin轻松地添加预渲染，已经被Vue应用程序广泛测试。

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
    res.status(500).end('Internal Server Error!');
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

将在服务器上预取数据prefetch data，开始渲染时，我们的应用程序就已经解析完成其状态。将数据进行响应式的过程在服务器上是多余的，默认情况下禁用。

## 组件生命周期钩子函数

没有动态更新，所有的生命周期钩子函数中，只有beforeCreate和create会在服务器端渲染过程中被调用。其它生命周期钩子函数中的代码beforeMount，mounted只会在客户端执行。

避免beforeCreate和create生命周期时产生全局副作用的代码。使用setInterval。副作用代码移动到beforeMount或mounted生命周期中。服务器端和客户端执行代码的差异性。

## 访问特定平台的API

通用代码不可接受特定平台的API，直接使用window，document，仅浏览器可用的全局变量，在Nodejs执行时会抛出错误。

对于共享于服务器和客户端，但用于不同平台的API的任务，建议将平台特定实现包含在通用API中，或者为你执行此操作的类库。axios是一个HTTP库，可以向服务器和客户端都暴露相同的API。

对于纯浏览器可用的API，通常方式是，在纯客户端的生命周期钩子函数中惰性访问。

## 自定义指令

大多数自定义指令直接操作DOM，因此会在服务器渲染过程中导致错误。有两者方法可以解决这个问题：推荐使用组件作为抽象机制，并运行在虚拟DOM层；如果你有一个自定义指令，但是不是很容易替换为组件，可以在创建服务器renderer时，使用directives选项所提供的服务器版本。

# 源码结构

## 避免状态单例

当编写纯客户端代码时，习惯在每次在新的上下文中对代码进行取值。Nodejs是一个长期运行的进程，如果创建一个单例对象，它将在每个传入的请求之间共享。

我们不应该直接创建一个应用程序实例，而是应该暴露一个可以重复执行的工厂函数。同样的规则也适用于router，store和event bus实例，不应该直接从模块导出并将其导入到应用程序中，而是需要在createApp中创建一个新的实例，并从根Vue实例注入。 

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

尽管Nodejs最新版本能够完全支持ES2015特性，还是需要转译客户端代码以适应老版浏览器。还是需要转译客户端代码以适应老版浏览器。

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
app.js是我们应用程序的通用entry，在纯客户端应用程序中，将在此文件中创建根Vue实例，并直接挂在到DOM。但是，对于服务器渲染SSR，责任转移到纯客户端entry文件，app.js简单地使用export导出一个createApp函数。

entry-client.js只需创建应用程序，并且将其挂在到DOM中。

entry-server.js每次渲染中重复调用此函数。除了创建和返回应用程序实例之外，不会做太多事情。稍后将在此执行服务器路由匹配和数据预取逻辑。

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

# 数据预取和状态

## 数据预取存储容器Data Store

在服务器端渲染期间，我们本质上是在渲染我们应用程序的快照，所以如果应用程序依赖于一些异步数据，那么在开始渲染过程之前，需要先预取和解析好这些数据。

另一个需要关注的问题是在客户端，在挂载到客户端应用程序之前，需要获取到与服务器端应用程序完全相同的数据。否则-客户端应用程序会因为使用与服务器端应用程序不同的状态，然后导致混合失败。

为了解决这个问题，获取的数据需要位于视图组件之外，放置在专门的数据预取存储容器或状态容器中。服务器端，可以在渲染之前预取数据，并将数据填充到store中。html中序列化和内置预置状态，挂载到客户端应用程序之前，可以直接从store获取到内联预置状态。

为此，将使用官方状态管理库Vuex。

```js
import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

import { fetchItem } from './api';

export function createStore() {
    return new Vuex.Store({
        state: {
            items: {}
        },
        actions: {
            fetchItem({ commit }, id) {
                return fetchItem(id).then(item => {
                    commit('setItem', { id, item })
                })
            }
        },
        mutations: {
            setItem(state, { id, item }) {
                Vue.set(state.items, id, item);
            }
        }
    });
}
```

## 带有逻辑配置的组件

需要通过访问路由，来决定获取哪部分数据-这也决定了哪些组件需要渲染。事实上，给定路由所需的数据，也是在该路由上渲染组件时所需的数据。

将在路由组件上暴露出一个自定义静态函数asyncData，此函数会在组件实例化之前调用，所以无法访问this。

```html
<template>
  <div>{{item.title}}</div>
</template>
<script>
export default {
  asyncData({ route, store }) {
    return store.dispatch("fetchItem", route.params.id);
  },
  computed: {
    item() {
      return this.$store.state.items[this.$route.params.id];
    }
  }
};
</script>
```

## 服务器端预取数据

在entry-server.js中，可以通过路由获得与router.getMatchedComponnets相匹配的组件，如果组件暴露出asyncData，我们就调用这个方法。然后我们需要解析完成的状态，附加到渲染上下文中。

```js
export default context => {
    return new Promise((resolve, reject) => {
        const { app, router, store } = createApp();
        router.push(context.url);
        router.onReady(() => {
            const matchComponents = router.getMatchedComponents();
            if (!matchComponents) {
                return reject({ code: 404 });
            }
            Promise.all(matchComponents.map((Component) => {
                if (Component.asyncData) {
                    return Component.asyncData({
                        store, route: router.currentRoute
                    })
                }
            })).then(() => {
                context.state = store.state;
                // 在所有预取钩子preFetch hook resolve后
                // 我们的store现在已经填充入渲染应用程序所需的状态
                // 当我们把状态附加到上下文，并且template选项用于renderer时，
                // 状态将自动序列化为window.__INITIAL_STATE__，并注入HTML
                resolve(app);
            }).catch(reject);
        }, reject);
    });
}
```
当使用template时，context.state将作为window.__INITIAL_STATE__状态，自动嵌入到最终的HTML中。客户端，在挂载到应用程序之前，store就应该获取到状态。

## 客户端数据预取

在路由导航之前解析数据，应用程序会等待视图所需数据全部解析之后，再传入数据并处理当前视图。建议提供一个数据加载指示器。

```js
router.onReady(() => {
    // 添加路由钩子函数，用于处理asyncData
    // 在初始路由resolve后执行，不会二次预取已有的数据
    // 使用router.beforeResolve，以便确保所有异步组件都resolve
    router.beforeResolve((to, from, next) => {
        const matched = router.getMatchedComponents(to);
        const preMatched = router.getMatchedComponents(from);
        let diffed = false;
        const activated = matched.filter((c, i) => {
            return diffed || (diffed = (preMatched[i] !== c))
        });
        if (!activated.length) {
            return next();
        }
        Promise.all(activated.map(c => {
            if (c.asyncData) {
                return c.asyncData({ store, route: to })
            }
        })).then(() => {
            next();
        }).catch(next)
        app.$mount('#app');
    });
});
```
## 匹配要渲染的视图后，再获取数据

此策略将客户端数据预取逻辑，放在视图组件的beforeMount函数中。当路由导航被触发时，可以立即切换视图，因此应用程序具有更快的响应速度。对于使用此策略的每个视图组件，都需要具有条件加载状态。

```js
Vue.mixin({
    beforeMount() {
        const { asyncData } = this.$options;
        if (asyncData) {
            this.dataPromise = asyncData({
                store: this.$store,
                route: this.$route
            })
        }
    },
    beforeRouteUpdate(to, from, next) {
        const { asyncData } = this.$options;
        if (asyncData) {
            asyncData({ store: this.$store, route: to }).then(next).catch(next)
        } else {
            next();
        }
    }
})
```

## Store代码拆分

大型应用程序中，Vuex store可能会分为多个模块。可以将这些模块代码，分割到相应的路由组件chunk中。
模块是路由组件的依赖，所以它被webpack移动到路由组件的异步chunk中。

```js
export default {
    namespaced: true,
    state: () => {
        return { count: 0 }
    },
    actions: {
        inc: ({ commit }) => commit('inc')
    },
    mutations: {
        inc: state => state.count++
    }
}
```

```html
<template>
  <div>{{fooCount}}</div>
</template>
<script>
import fooStoreModule from "../store.foo";
export default {
  asyncData({ route, store }) {
    store.registerModule("foo".fooStoreModule);
    return store.dispatch("foo/inc");
  },
  destroyed() {
    this.$store.unregisterModule("foo");
  },
  computed: {
    fooCount() {
      return this.$store.state.foo.count;
    }
  }
};
</script>
```
## 客户端激活

所谓客户端激活，指的是Vue在浏览器端接管由服务器端发送的静态HTML，使其变为由Vue管理的动态DOM的过程。

无需将服务器渲染好的HTML丢弃再重复创建所有的DOM元素，相反，我们需要激活这些静态HTML，然后使他们成为动态的，能够响应后续的数据变化。

服务器渲染的输出结果，在应用程序的根元素上添加了一个特殊的属性data-server-rendered。这个属性让客户端Vue知道这部分HTML是由Vue在服务器端渲染的，应该以激活模式进进行挂载。

```html
<div id="app" data-server-rendered="true"></div>
```

可以向$mount函数的hydrating参数为主传入true，来强制开启激活模式。

在开发模式下，Vue将推断客户端生成的虚拟DOM树，是否与服务器渲染的DOM结构匹配，无法匹配，将推出混合模式，丢弃现有DOM并从头开始渲染，在生产模式下，此检测会被跳过，以避免性能损耗。

```js
app.$mount('#app',true)
```
## 一些需要注意的坑

使用SSR+客户端混合时，需要了解的一件事是，浏览器可能会更改一些特殊的HTML结构。所以导致无法匹配。

# Bundle Renderer指引

## 使用基本SSR的问题

假设打包的服务器端代码，将由服务器通过require直接使用，这是里所应当的，然而在每次编辑过应用程序源代码之后，都必须停止并重启服务。开发过程中影响开发效率。

## 传入BundleRenderer

vue-server-render提供一个creteBundleRenderer的API，用于处理此问题，通过使用webpack的自定义插件，server bundle将生成可传递到bundle renderer的特殊JSON文件。

bundle renderer提供一下优点：内置的source map支持；在开发环境甚至部署过程中热重载。关键CSS注入，自动内联在渲染过程中用到的组件所需的CSS。使用clientManifest进行资源注入，自动推断出最佳的预加载preload和预取prefetch指令，以及初始渲染所需的代码分割chunk。

# 构建配置

## 服务器配置

是用于生成传递给createBUndleRenderer的server bundle。它应该是这样的。


