import { createApp } from './app';

const { app, router, store } = createApp();

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