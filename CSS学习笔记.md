css

css全名是cascading style sheets，意为层叠样式表，用于定义网页样式和布局的样式表语言。通过css可以指定页面中各个元素的颜色、字体、大小、间距、边框、背景等样式，从而实现更精确的页面设计。

语法

css由选择器以及一条或多条声明组成，

![picture](https://www.runoob.com/wp-content/uploads/2013/07/632877C9-2462-41D6-BD0E-F7317E4C42AC.jpg)

选择器通常是需要改变样式的 HTML 元素。每条声明由一个属性和一个值组成。

属性是您希望设置的样式属性。每个属性有一个值。属性和值被冒号分开

选择器的声明可以写无数条属性

声明的每一行属性，都要以英文分号结尾

p {    color:red;    text-align:center; }

css导入方式

1. 内联样式
2. 内部样式表
3. 外部样式表

选择器

类选择器用于描述一组元素的样式

类选择器以一个点 **.** 号显示

.center {text-align:center;}



元素选择器

选择并设置所有 <p> 元素的样式：

```css
p
{ 
background-color:yellow;
}
```



id 选择器可以为标有特定 id 的 HTML 元素指定特定的样式，id 选择器以 "#" 来定义。

\#para1 {    text-align:center;    color:red; }



通用选择器

选择所有元素，并设置它们的背景色：

```css
*
{ 
background-color:yellow;
}
```

设置元素

p.ex
{
	height:100px;
	width:100px;
}

盒子模型



![picture](https://www.runoob.com/images/box-model.gif)



- **Margin(外边距)** - 清除边框外的区域，外边距是透明的。

- **Border(边框)** - 围绕在内边距和内容外的边框。

- **Padding(内边距)** - 清除内容周围的区域，内边距是透明的。

- **Content(内容)** - 盒子的内容，显示文本和图像。

   margin

margin 清除周围的（外边框）元素区域。margin 没有背景颜色，是完全透明的。

margin 可以单独改变元素的上，下，左，右边距，也可以一次改变所有的属性。

![picture](https://www.runoob.com/wp-content/uploads/2013/08/VlwVi.png)

可以指定不同的侧面不同的边距

margin-top:100px; margin-bottom:100px; margin-right:50px; margin-left:50px;

 **margin:25px 50px 75px 100px;**

- 上边距为25px
- 右边距为50px
- 下边距为75px
- 左边距为100px

**margin:25px 50px 75px;**

- 上边距为25px
- 左右边距为50px
- 下边距为75px

outline

| 属性                                                         | 说明                           | 值                                                           |
| :----------------------------------------------------------- | :----------------------------- | :----------------------------------------------------------- |
| [outline](https://www.runoob.com/cssref/pr-outline.html)     | 在一个声明中设置所有的轮廓属性 | *outline-color outline-style outline-width *inherit          |
| [outline-color](https://www.runoob.com/cssref/pr-outline-color.html) | 设置轮廓的颜色                 | *color-name hex-number rgb-number *invert inherit            |
| [outline-style](https://www.runoob.com/cssref/pr-outline-style.html) | 设置轮廓的样式                 | none dotted dashed solid double groove ridge inset outset inherit |
| [outline-width](https://www.runoob.com/cssref/pr-outline-width.html) | 设置轮廓的宽度                 | thin medium thick *length *inherit                           |

border-style值

none: 默认无边框

dotted: 定义一个点线边框

dashed: 定义一个虚线边框

solid: 定义实线边框

double: 定义两个边框。 两个边框的宽度和 border-width 的值相同

groove: 定义3D沟槽边框。效果取决于边框的颜色值

ridge: 定义3D脊边框。效果取决于边框的颜色值

inset:定义一个3D的嵌入边框。效果取决于边框的颜色值

outset: 定义一个3D突出边框。 效果取决于边框的颜色值

定位

position 属性的五个值：

- static

​        HTML 元素的默认值，即没有定位，遵循正常的文档流对象。静态定位的元素不会受到 top, bottom, left, right影响。

​         div.static {    position: static;    border: 3px solid #73AD21; }

- relative

​       相对定位元素的定位是相对其正常位置。

- fixed

​      元素的位置相对于浏览器窗口是固定位置。即使窗口是滚动的它也不会移动：

​       p.pos_fixed {    position:fixed;    top:30px;    right:5px; }

- absolute

​        绝对定位的元素的位置相对于最近的已定位父元素，如果元素没有已定位的父元素，那么它的位置相对于<html>

- sticky

​       粘性定位的元素是依赖于用户的滚动，在 **position:relative** 与 **position:fixed** 定位之间切换。

伪类

CSS伪类是用来添加一些选择器的特殊效果。

伪类的语法：

selector:pseudo-class {property:value;}

CSS类也可以使用伪类：

selector.class:pseudo-class {property:value;}

传统网页布局方式

1. 标准流
2. 浮动
3. 定位
4. flexbox
5. grid

