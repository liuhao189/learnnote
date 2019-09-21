# 前言

2019年，主流PC&&移动操作系统均支持暗黑模式(dark mode)。用户在系统中设置了暗黑模式后，为了统一、良好的使用体验，app和网页应该支持暗黑模式。

# 网页暗黑模式

默认情况下，浏览器中的网页不会有任何变化。但是，这会在暗黑模式下留下大片高亮度的内容，一些App中内嵌webview会对网页进行简单处理。eg：Mac的mail客户端会将背景色设为暗色，文字设为亮色。

并不是所有的网页内容简单的(可以被很好处理的)，所以浏览器不会自动改变网页的展示效果。

## Safari声明支持简单内容暗黑模式

```css
:root{
 color-schemes: light dark;
}
```
声明支持暗黑模式后，浏览器会自动改变背景颜色，文本颜色，标准控件的颜色以match暗黑模式。

除了声明color-schemes变量外，还可以通过meta标签声明支持暗黑模式。

```html
<meta name="color-schemes" value="light dark">
<meta name="color-schemes" value="light only">
```

  color-scheme属性从supported-color-schemes更名而来。

#  复杂内容的暗黑模式

对于复杂的网页，你需要使用prefers-color-scheme的媒体查询来变更颜色，图片，布局等样式信息。

比较好的方式支持暗黑模式，是使用CSS变量。

```css
:root{
--bg-color: white;
--font-color: black;
}

@media (prefers-color-scheme: dark){
  :root {
    --bg-color:black;
    --font-color: white;
    }
}
```
## 图片和暗黑模式

可以在picture元素里应用media query选择合适的图片来展示。

```html
    <picture>
        <source srcset="./imgs/black.jpg"
            media="(prefers-color-scheme:dark)">
        <img src="./imgs//white.jpg">
    </picture>
```

## 目前支持的网站

[Twitter](https://mobile.twitter.com/)
[CSS DB](https://cssdb.org/)

# 兼容性问题

IOS13，win10 && mac 10.14.4+的 chrome浏览器均支持。

[https://www.caniuse.com/#search=color-scheme](https://www.caniuse.com/#search=color-scheme)

