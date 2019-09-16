# GSAP

GSAP(GreenSock Animation Platform)可以给JS可控制的元素变化(CSS属性，SVG，React，Canvas，Generic objects)添加动画。

GASP还解决了浏览器的兼容性问题，同时性能比较高。被很多大中型网站使用。

## 属性控制者

Animation的本质是不断改变某个属性值。GASP用一个初始值，一个结束值和一些配置即可实现属性值的变化。

开发者需要决定改变什么属性，改变速度，延迟等属性。

## DOM，SVG，Canvas and beyond

GASP依赖插件来支持不同的属性变换。

1、CSS，2D&&3D transforms，colors，width，opacity，border-radius，margin和其它CSS属性。

2、SVG Attribute,viewBox, width, height,fill,stroke, cx,r,opacity.

3、Any numberic value. 

## GSPA 包含什么

GSPA包含一系列脚本动画工具。

1、TweenLite, 轻量版本，包含改变属性功能的核心引擎。

2、TweenMax，包含常用功能。包含TweenLite，TimelineLite，TimeLineMax，CSSPlugin，AttrPlugin，RoundPropsPlugin，BezierPlugin和EasePack。

3、TimelineLite&&TimelineMax，组动画管理。

4、其它插件。

## Loading GSAP
 
CDN && NPM && Download Zip && Github

## Tweening Basics

我们使用TweenMax来完成以下示例。

### TweenMax.to

为了创建一个动画，TweenMax.to需要3个东西。

target：需要动画的元素，可以是一个JS对象，对象数组，选择器。

duration：单位秒。

vars：你希望动画to的属性值，或一些回调函数。

    TweenMax变换不止作用于DOM，还可以作用于对象的属性。

## Plugins

插件是为了提供额外的能力。这使得核心代码小巧且高效，同时允许功能扩展。

TweenMax包括CSSPlugin，AttrPlugin，BezierPlugin。

### CSSPlugin

CSSPlugin首先会识别出目标对象是DOM元素。然后把属性转换到内联样式。

CSSPlugin提供的功能：

1、浏览器的兼容性问题。

2、提高性能，自动加层，缓存，阻止布局抖动。

3、使2D和3D不同属性的变换独立。

4、读取初始值。

5、复合值的动画。

6、处理颜色插值。

### 2D && 3D 变换

GSAP可以支持所有transform属性。我们强烈建议您使用GASP简称写法。因为它们更加快速和精确。

## TweenMax.from

从end-value to current-value的动画。

## TweenMax.formTo

formTo方法允许你定义开始和结束的值。

## 特别属性

onComplete。特别属性不同于普通的动画属性值。特别属性用于定义回调，延迟，动画缓动函数等等。

1、onComplete，动画完成时的回调。

2、onUpdate，动画更新渲染时的回调。

3、ease，缓动函数

## Easing

如果你想要成为动画大师，需要开发一个easing函数。easing函数决定了一个从点A到点B的运动。

如果想自定义Ease，GSAP提供一个可视化工具。

## Callbacks

回调事件在特定的动画事件发生时会调用。

1、onComplete

2、onStart

3、onUpdate

4、onRepeat

5、onReverserComplete

### 回调函数参数和作用域

任何回调都可以传递任意数量参数，必须以数组形式传递给lifeCycleParams属性。

默认情况下，回调函数的this是tween实例，你可以定义为其它对象。lifeCycleScope属性。

### 控制动画

为了控制动画，你需要一个动画实例。to，from，fromTo方法都返回一个实例。

## 时间轴序列动画

可以通过TimelineLite或TimnelineMax方便地创建序列动画。

一个事件轴是一个动画的容器。可以调用timeline的play，reverse，pause等方法。timeline甚至可以嵌套timeline。

什么场景下使用timeline：

1、控制一组动画

2、按时间顺序发生的动画，timeline容易书写和维护

3、模块化动画代码

to，from，formTo是添加动画到timeline的快速方法。

默认情况下，动画会顺序插入到timeline中。

### 链式调用











