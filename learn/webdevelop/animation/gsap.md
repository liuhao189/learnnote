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

vars：你希望动画到的属性值，或一些回调函数。

    TweenMax变换不止作用于DOM，还可以作用域对象的属性。

## Plugins

插件是为了提供额外的能力。这使得核心代码小巧且高效，同时允许功能扩展。

TweenMax包括CSSPlugin，AttrPlugin，BezierPlugin。

### CSSPlugin

CSSPlugin首先会识别出目标对象是DOM元素。然后把属性转换到内联样式。

CSSPlugin提供的功能：

1、浏览器的兼容性问题。

2、提高性能，自动加层，缓存组件，组织布局抖动。

3、使2D和3D不同属性的变换独立。

4、读取初始值。

5、复合值的动画。

6、处理颜色插值。

### 2D && 3D 变换

GSAP可以支持所有transform属性。我们强烈建议您使用GASP简称写法。因为它们更加快速和精确。




