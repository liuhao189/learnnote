# 进入/离开&列表过渡

Vue在插入、更新或移除DOM时，提供多种不同方式的应用过渡效果。

1、在CSS过渡和动画中自动应用class。

2、可以配合使用第三方CSS动画类，如Animate.css

3、在过渡钩子函数中使用JS直接操作DOM

4、可以配合使用第三方JS动画库，如Velocity.js

## 单元素/组件的过渡

Vue提供了transition的封装组件，可以给任何元素和组件添加进入/离开过渡。

1、条件渲染，条件展示，动态组件，组件根节点。

当插入或删除包含在Transition组件中的元素时，Vue将会做以下处理：

1、自动嗅探目标元素是否应用了CSS过渡或动画，如果是，在恰当的时机添加/删除CSS类名。

2、如果过度组件提供了JS钩子函数，这些钩子函数将在恰当的时机被调用。

3、如果没有找到JS钩子并且也没有检测到CSS过度/动画，DOM操作在下一帧立即执行。

# 过渡的类名

在进入/离开的过渡中，会有6个class切换。

1、v-enter，定义进入过渡的开始状态，元素被插入之前生效，元素被插入之后的下一帧移除。

2、v-enter-active，定义进入过渡时的状态，在整个进入过渡的阶段中应用，元素被插入之前生效，过渡/动画完成之后移除，可以定义进入过渡的过程时间，延迟和曲线函数。

3、v-enter-to，定义进入过渡的结束状态，元素被插入之后下一帧生效，过渡/动画完整之后移除。

4、v-leave，离开过渡的开始状态，离开过渡被触发时立刻生效，下一帧被移除。

5、v-leave-active，定义离开过渡生效时的状态，整个离开过渡的阶段中应用，过渡完成之后立刻移除。

6、v-leave-to，离开过渡的结束状态，离开过渡下一帧生效，过渡/动画完成之后移除。

    enter,leave两个大状态。CSS动画主要依赖CSS数字属性变化，进入时添加类，马上添加另一个类&&过渡属性类，过渡完成后删除。

## CSS动画

CSS动画用法同CSS过渡，区别在于动画中v-enetr类名在节点插入DOM后不会立即删除，而是在animationend事件触发时删除。

## 自定义过渡的类名

可以通过以下特性来自定义过渡类名

enter-class，enter-active-class，enter-to-class，leave-class，leave-active-class，leave-to-class。优先级高于普通的类名，对于Vue的过渡系统和其他第三方CSS动画库，如Animate.css结合使用十分有用。

## 同时使用过渡和动画

Vue需要知道过渡的完成，必须设置相应的事件监听器。可以是tranitionend或animationend，这取决于给元素应用的CSS属性。如果要给一个元素同时设置两种过渡动效，必须使用type特性并设置animation或transition来明确什么你需要Vue监听的类型。

## 显性地过渡时间

默认情况下，Vue会等待其在过渡效果的根元素的第一个transitionend或animationend事件。我们可以拥有一个精心编排的过渡效果，嵌套的内部元素相比于过渡效果的根元素有延迟或更长的过渡效果。

可以使用duration属性定制一个显性地过渡持续时间。

也可以定制进入和移出的持续时间。

## JS钩子

可以在属性中声明JS钩子

当只用JS过渡的时候，在enter和leave中必须使用done进行回调，否则，它们将被同步调用，过渡会立即完成。

推荐对于仅使用JS过渡的元素添加:css="false"，vue会跳过CSS的检测，避免过渡过程中的CSS的影响。

## 初始渲染的过渡

可以通过appear特性设置节点在初始渲染的过渡。也可以自定义CSS类名。或自定义JS钩子。

# 多个元素的过渡

对于原生标签可以使用v-if/v-else。最常见的多标签过渡是一个列表和描述这个列表为空消息的元素。

    当有相同标签名的元素切换时，需要通过key特性设置唯一的来标记以让vue区分它们，否则Vue为了效率只会替换相同标签内部的内容。

使用了多个v-if的多个元素的过渡可以重写为绑定了动态属性的单个元素过渡。

```html
    <div id="app">
        <button @click="toggle">Toggle</button>
        <button @click="changeType">ChangeType</button>
        <transition appear
            name="fade">
            <table v-if="items.length>0">
                <tr v-for="item of items">
                    <td>{{item}}</td>
                </tr>
            </table>
            <p v-else>Sorry, no items found.</p>
        </transition>
        <transition appear
            name="fade">
            <span :key="currentStr"
                style="display:inline-block;">{{currentStr}}</span>
        </transition>
    </div>
```

## 过渡模式

transition的默认行为，进入和离开同时发生，一个离开过渡的时候另一个进入过渡。

过渡模式：

1、in-out，新元素先进行过渡，完成之后当前元素过渡离开。

2、out-in, 当前元素先进行过渡，完成之后新元素过渡进入。

# 多个组件的过渡

多个组件的过渡简单很多，不需要使用key特性，只需要使用动态组件即可。

```html
    <transition name="fade"
        appear
        mode="out-in">
        <component :is="view"></component>
    </transition>
```

# 列表过渡

目前为止，关于过渡我们已经讲到：单个节点，同一时间渲染多个节点中的一个。

同时渲染整个列表，使用v-for。这种场景中，使用transition-group组件。

1、不同于transition，它会以一个真实元素呈现，默认为一个span，也可以通过tag特性更换为其它元素。

2、过渡模式不可用，不再相互切换特有的元素。

3、内部元素总是需要提供唯一的key属性值。

4、CSS过渡的类将会应用在内部的元素中，而不是这个组/容器本身。

## 列表的进入/离开过渡

```html
    <transition-group name="list"
        appear
        tag="p">
        <span v-for="item in list"
            :key="item"
            class="list-item">{{item}}</span>
    </transition-group>
```

## 列表的排序过渡

transition-group组件还有一个特殊之处，不仅可以进入和离开动画，还可以改变定位，要使用这个新功能只需了解新增的v-move特性。它会在元素的改变定位的过程中应用。可以通过move-class属性手动设置。

内部的实现，Vue使用了一个叫FLIP简单的动画队列，使用transforms将元素从之前的位置平滑过渡到新的位置。

    使用FLIP过渡的元素不能设置为display:inline，作为替代方案，可以设置为display:inline-block或者放置于flex中。

FLIP动画不仅可以实现单列过渡，多维网格也同样可以过渡。

## 列表的交错过渡

通过data属性与JS通信，就可以实现列表的交错过渡。

```html
    <transition-group tag="ul"
        :css="false"
        appear
        @before-enter="beforeEnter"
        @enter="enter"
        @leave="leave">
        <li v-for="(item,inx) of computedList"
            :key="item.msg"
            :data-index="inx">
            {{item.msg}}
        </li>
    </transition-group>
```

# 可复用的过渡

过渡可以通过Vue的组件系统实现复用，创建一个可复用过渡组件，需要做的就是讲transition或transition-group作为根组件，然后将任何子组件放置在其中就可以了。

函数式组件更适合完成这个任务。

```js
 Vue.component('my-transition', {
            functional: true,
            render: function (createElement, context) {
                let data = {
                    props: {
                        name: 'reuse',
                        mode: 'out-in',
                        appear: true
                    },
                    on: {
                        beforeEnter: function (el, done) {
                            Velocity(el, {
                                opacity: 1,
                                translateX: 0,
                                fontSize: '38px'
                            }, {
                                    complete: () => {
                                        el.style.transition = ""
                                        done();
                                    }
                                })
                        },
                        enter: function (el) {
                            el.style.opacity = 0;
                            el.style.transform = "translateX(100px)"
                            el.style.transition = "all .5s linear";
                        }
                    }
                };

                return createElement('transition', data, context.children);
            }
```