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
    echartInstanbce.resize调整尺寸。

theme，应用的主体，可以是一个主题的配置对象，也可以是使用已经通过echarts.registerTheme注册的主体名称。

opts，附加参数：

devicePixelRatio，设备像素比，默认取浏览器的window.devicePixelRatio。

renderer，渲染器，支持canvas或者svg。

width，height显式指定宽高。为空或auto，自动去dom的宽高。

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

接触图标实例的联动，只需要移除单个实例，通过将该图标实例group设为空。

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

## echarts.graphic




