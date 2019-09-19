# 定义
以下摘抄自MDN定义：

    CSS变量是由CSS作者定义的实体，其中包含要在整个文档中重复使用的特定值，使用自定义属性来设置变量名，并使用特定的var()来访问。

解读：

1、声明变量时前面要加两根连词线--，更加容易区分，同时避免其它CSS属性冲突。

2、使用var方法来访问即可。

3、既然是变量，除了声明和使用外，重点在于作用域。

# 基本用法
## 1、局部变量

```css
element {
   --main-bg-color: brown;
}

element {
  background-color: var(--main-bg-color)
}
```
## 2、全局变量

```css
:root {
   --global-color: #666;
  --pane-padding: 5px 42px;
}

.demo {
  color: var(--global-color);
}
```

# 解决的问题

##  可维护性方面

1. 一个网站一般都有特定的UI设计方案，这意味着一些块颜色，空白，字体属性，阴影等会重复出现在CSS文件中。不使用CSS变量，修改时繁琐且易出错。

2. 变量名称包含了语义化的信息，CSS文件变得更加易读和容易理解。

## CSS变量 VS 预处理器

1. 预处理器需要添加额外的构建步骤，增加了系统的复杂性。编译后输出的CSS没有体现出变量的语义化信息。

# CSS变量作用域

CSS引入了一种层级变量的概念，自定义变量和CSS属性一样，作用在当前层级，若没有定义，则从其父元素继承其值。

同一个CSS变量，可以在多个选择器内声明，读取的时候，优先级最高的声明生效。

```css
:root{
 --color: blue;
}
.title {
 --color: green;
}

#info{
--color: #333;
}
```


# 变量值的类型

## 字符串

字符串可与其它字符串拼接。

```css
:root {
 --hi-str:  'Hello';
 --hi-css-var: var(--hi-str)'  CSS var';
}
```

## 数值

不能与数值单位连用，必须使用calc函数。

```css
:root{
  --normal-space: 20;
}

.card{
padding:  var(--normal-space)px;
margin: calc(var(--normal-space) * 1px );
}
```

# 其它知识点

1、var函数第二个参数，表示变量的默认值，如果该变量不存在，就会使用这个默认值。

# 兼容性

2019年9月19日 查询[CSS Variables ](https://www.caniuse.com/#search=css%20Variables )，92.85%可用。基本上除了IE和一些简化版手机浏览器之外均支持。