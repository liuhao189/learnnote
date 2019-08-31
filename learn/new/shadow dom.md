# Shadow DOM

Shadow DOM解决了构建网络应用的脆弱性问题。脆弱性是由HTML、CSS和JS的全局性引起的。

Shadow DOM在网络平台中引入作用域样式。

Shadow DOM是四大网络组件标准之一。HTML模板，Shadow DOM、自定义元素以及HTML导入。

## 特点

隔离DOM，组件的DOM是独立的。

作用域CSS：shadow DOM内部定义的CSS在其作用域内。样式规则不会泄露，页面样式也不会渗入。

组合：为组件设计一个声明性，基于标记的API。

简化CSS：作用域DOM意味着您可以使用简单的CSS选择器，更同样的id/类名称，而无需担心命名冲突。

效率：将应用看成是多个DOM块，而是一个大的页面。

## 什么是shadow DOM

Shadow DOM与普通DOM相同，但有两点区别。1、创建/使用的方式；2、也页面其它部分有关的行为方式。

作用域子树称为影子树，被附着的元素称为影子宿主。在影子中添加的任何项都将称为宿主元素的本地项，包括style。

## 创建shadow DOM

 element.attachShadow()。

 有些元素无法托管影子树，其原因如下：
 
 1、浏览器已未该元素托管其自身的内部shadow DOM<textarea>,input。

 2、让元素托管shadow DOM毫无意义img

 ## 为自定义元素创建shadow DOM

 创建自定义元素时，Shadow DOM尤其有用。使用shadow DOM来分隔元素的HTML，CSS和JS，从而生成一个网络组件。

 ## 组合和Slot

 组合是shadow DOM最难理解的功能之一，但可以说是最重要的功能。

 slot元素：Shadow DOM使用slot元素将不同的DOM树组合在一起。Slot是组件内部的占位符，用户可以使用自己的标记来填充。

 通过定义一个或多个slot，您可将外部标记引入到组件的shadow DOM中进行渲染。

 如果slot引入了元素，这些元素可跨越shadow DOM的边界。这些元素称为分布式节点。

 组件可在其shadow DOM中定义零个或多个slot，Slot可以为空，或者提供回退内容。

 您还可以创建已命名slot，已命令slot是shadow DOM中用户通过名称引用的特定槽。

 