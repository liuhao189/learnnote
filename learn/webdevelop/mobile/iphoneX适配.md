# 前言

iPhoneX取消了物理按键，改成底部小黑条，这一改动导致网页出现了比较尴尬的屏幕适配问题。对于网页来说，顶部刘海部分的适配问题浏览器已经做了处理，开发者只需要关注底部于小黑条的适配问题即可。

# 新概念

## 安全区域

安全区域指的是一个可视窗口范围，处于安全区域的部分不受圆角，齐刘海，小黑条的影响。

做好适配，必须保证页面可视，可操作区域是在安全区域内。

## viewport-fit

IOS11新增特性，为了适配iPhoneX对现有的viewport meta标签的一个扩展，用于设置网页在可视窗口的布局方式。

contain，可视窗口完全包含网页内容；cover，网页内容完全覆盖可视窗口；auto，默认值，跟containe表现一致。

  需要适配iPhoneX必须设置viewport-fit=cover，这是适配的关键步骤。

## env()和constant()

IOS11新增特性，Webkit的一个CSS函数，用于设定安全区域与边界的距离。有四个预定义的变量。

safe-area-inset-left，safe-area-inset-right,safe-area-inset-top,safe-area-inset-bottom。

需要关注的safe-area-inset-bottom这个变量，因为它对应的就是小黑条的高度。横竖屏时值不一样。

  当viewport-fit=contain时env是不起作用，必须要配合viewport-fit=cover使用。不支持的浏览器，将忽略env。

env和constant作用相同，但是，IOS11.2之后constant是不能使用的，需要做向后兼容。

# 如何适配

1、新增viewport-fit属性，使得页面内容完全覆盖整个窗口。

2、页面主体内容限定在安全区域内。

## fixed元素的适配

1、fixed完全吸底元素，bottom：0。

padding-bottom，calc覆盖原来高度，

2、fixed非完全吸底元素bottom!=0，返回顶部，侧边广告等

需要对应向上调整即可，margin，calc覆盖原来bottom值。

## @supports来隔离兼容样式

@support来隔离兼容样式。