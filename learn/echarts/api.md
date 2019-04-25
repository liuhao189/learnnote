# echarts

全局echarts对象，在script引入echarts.js文件后获得，或者在AMD环境中通过require('echarts')获得。

## init
```typescript
echarts.init(dom:HTMLDivElement|HTMLCanvasElement,theme?:Object|String,opts?:{
    devicePixelRatio?:number,
    renderer?:string,
    width?:number|string,
    height?:number|string
})
```

dom，实例容器，一般具有高宽的div元素。

    如果div是隐藏的，ECharts可能会获取不到div的高宽导致初始化失败，可以明确指定div的width和高，或者div显示后手动调用
    echartInstance.resize调整尺寸。

theme，应用的主体，可以是一个主题的配置对象，也可以是使用已经通过echarts.registerTheme注册的主题名称。

opts，附加参数：

devicePixelRatio，设备像素比，默认取浏览器的window.devicePixelRatio。

renderer，渲染器，支持canvas或者svg。

width，height显式指定宽高。为空或auto，自动取dom的宽高。

## connect

多个图标实现联动。

```js
connect(group:string|Array)
```
group的id，或者图标实例的数组。

## disconnect

```js
disconnect(group:string)
```

解除图表实例的联动，只需要移除单个实例，通过将该图表实例group设为空。

## dispose

销毁实例，实例销毁后无法再被使用。

```js
echarts.dispose(target:ECharts|HTMLDivElement|HTMLCanvasElement)
```

## getInstanceByDom 

获取dom容器上的实例

```js
getInstance(target:HTMLDivElement|HTMLCanvasElement)
```

## registerMap

```js
registerMap(mapName:string,geoJson:Object,spcialAreas?:Object)
```

注册可用的地图，必须在包括geo组件或者map图表类型的时候才能使用。

参数：

mapName，地图名称，在geo组件或者map图表类型中设置的map对应的就是该值。

geoJSON，GeoJson格式的数据。

specialAreas，将地图中的部分区域缩放到合适的位置，可以使得整个地图更加好看。

## getMap

获取已注册的地图。
```js
getMap(mapName:string)
```

## registerTheme

注册主题，用于初始化实例的时候指定。


# echartsInstance

通过init创建的实例。

## group

图表的分组，用于联动。

```js
group:string|number
```

## setOption

设置图表实例的配置项以及数据，万能接口，所有参数和数据的修改都可以通过setOption完成。ECharts会合并新的参数和数据，然后刷新图标。如果开启动画的话，ECharts找到两组数据之间的差异然后通过合适的动画去表现数据的变化。

```js
setOption(option:Object,notMerge:boolean,lazyUpdate?:boolean)
setOption(option:Object,opts?:Object)
```

notMerge可选，是否不跟之前设置的option进行合并，默认为false，即合并。

lazyUpdate可选，在设置完option后是否不利己更新图标，默认为false。

slient，可选，阻止调用setOption时抛出事件，默认为false，即抛出事件。

## getWidth && getHeight

获取ECharts实例容器的宽度和高度

## getDom && getOption

获取echarts实例容器的DOM节点，获取当前实例中维护的option对象，包含用户多次setOption合并得到的配置项和数据，也记录了用户交互的状态，图例的开关，数据区域缩放选择的范围。这份option可以恢复或者得到一个新的一模一样的实例。

## resize

改变图标尺寸，在容器大小发生改变时需要手动调用。

```js
resize(opts:{width:number|string,height:number|string,slient:boolean})
```
## dispacthAction

触发图标行为，例如图例开关legendToggleSelect，数据区域缩放dataZoom，显示提示框showTip，更多见action和events的文档。

```js
dispacthAction(payload:Object)
```

## on

绑定事件处理函数，ECharts中的事件有两种，一种是鼠标事件，一种是调用dispatchAction后触发的事件。每个action都会有对应的事件。

```js
on(eventName:string,handler:Function,context:Object)
on(eventName:string,query:string|Object,handler:Function,context?:Object)
```

eventName：事件名称，全小写。统一使用跟DOM事件一样的全小写字符串作为事件名。

query：可选的过滤条件，能够只在指定的组件或者元素上进行响应。

## off

解绑事件处理函数

```js
off(eventName:string,handler?:Function)
```

eventName：事件名称

handler：可以传入需要解绑的处理函数，不传的话解绑所有类型的事件函数。




