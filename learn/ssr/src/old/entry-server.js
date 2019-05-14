import { createApp } from './app';

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