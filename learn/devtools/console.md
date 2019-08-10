# console

## 消息堆叠

如果一条消息连续重复，控制台将堆叠消息并在左侧外边距显示一个数字。

倾向于为每一个日志使用一个独特的行条目，请在DevTools设置中启用Show timestamps。

## 清除历史记录

控制台中输入clear()，JS代码内调用console.clear()。

## 保留console记录

设置preserve log即可，消息将一直显示，直至清除控制台或关闭标签。

## console.table

console.table(arrs)，可打印数组，对象。

## console.group

console.group可对消息分组，使用console.group(name)开始，console.groupEnd(name)结束。

```js
let labelName='testGroup'
console.group(labelName);
console.log(1);
console.log(2)
console.log(3)
console.groupEnd(labelName)
```

## 自定义样式消息

使用console.log('%c msg',styles)即可。

```js
{
    let styles=`padding:5px;color:red;font-size:3em;`
    console.log("%c grow up",styles);
}
```

## 支持正则过滤消息 && 支持消息来源过滤

正则直接/^[AH]/即可，消息来源过滤在左侧展示侧边栏，messages下拉里选择即可。

## Console和其它Tab共存

ESC即可。

## REPL

REPL是READ，Evaluate，Print和Loop的简称。

## $&&debug&&keys

$受到了JQuery的启发，可以快捷选择元素。

debug(function)可以在函数的第一行设置断点。

## Eager Evaluation

输入JS表达式时是否展示预览信息。

## console.assert

console.assert(expression,object)。

```js
console.assert(x<y,{x,y,reason})
```

## console.count([label])&&console.countReset

相同label的计数器。

countReset清空计数器。

## console.debug && console.dir

dir展示对象的详细JSON属性。

debug和console.log(object,[,object,...])一样。

dirxml以XML格式展示元素的后代节点。

## groupCollapsed(label)

收缩group的信息们。

## console.time([label]) && console.timeEnd

以ms形式记录时间差。

## console.trace

打印当前的调用堆栈信息。

## $$

$$同document.querySelectorAll类似。

$$(selectorStr,startNode)。

## getEventListeners

返回注册的所有事件监听器。

# monitor

监控某个函数，函数被执行时打印出函数调用参数等信息。

# monitorEvent(object,[,events])

当该对象的某个事件触发时，打印相关的信息。

# profile && profileEnd

profile开始一个JS的CPU分析会话。profileEnd完成并显示分析结果。

## 实时表达式

如果你需要不断查看某个表达式的值，你可以创建实时表达式。

