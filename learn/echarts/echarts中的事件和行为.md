# ECharts中的事件和行为

ECharts的图表中用户的操作将会触发相应的事件，开发者可以监听这些事件，然后通过回调函数做相应的处理。

Echarts中绑定事件通过on方法，事件名称更加简单了。事件名称对应DOM事件名称，均为小写字符串。

Echarts中事件分为两种类型，一种是用户鼠标操作点击，还有一种是用户在使用可以交互的组件后触发的行为事件。

## 鼠标事件的处理

ECharts支持常规的鼠标事件类型，包括click，dbclick，mousedown，mousemove，mouseup，mouseout，globalout，contentmenu事件。

所有的鼠标事件包含参数params，这是一个包含点击图形的数据信息的对象。

chart.on(eventName,query,handler)，query为string或object，如果为string表示组件类型。格式可以是mainType或者mainType.subType。

如果query为object，可以包含以下一个或多个属性，每个属性都是可选的。Index，Name，Id，dataIndex,name,dataType,element。

可以在回调中获取这个对象中的数据名，系列名称后进行任何操作。





