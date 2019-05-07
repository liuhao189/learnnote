# 使用dataset管理数据

## 默认的映射

ECharts针对最常见的直角坐标系中的图表（折线图、柱状图、散点图、K线图等）、饼图、漏斗图。给出了简单的默认的映射，从而不需要配置encode也可以出现图表（一旦给出了encode，那么就不会采用默认映射）。

默认映射的基本规则大体是：

1、在坐标系中（直角坐标系，极坐标系等）：如果有类目轴：axis.type为Category，则将第一行或列映射到这个轴上，后续每一行或列对应一个系列。如果没有类目轴，例如直角坐标系的XY轴，则每两列对应一个系列，这两列分别映射到这两个轴上。

2、如果没有坐标系，取第一列或行为名字，第二列或行为数值。

默认的规则不能满足要求时，就可以自己来配置encode，也并不复杂。

## 数据的各种格式

多数常见图表中，数据适用于二维表的形式描述。广为使用的数据表格软件(Excel、Numbers)或关系数据数据库都是二维表。它们的数据可以导出成JSON格式，输入到dataset.source中，在不少情况下可以免去一些数据处理的步骤。

    假如数据导出成CSV文件，那么可以使用一些CSV工具如DSV或PapaParse将CSV转成JSON。

在JS常用的数据传输格式中，二维数组可以比较直观的存储二维表。除了二维数组以外，dataset也支持key-value方式的数据格式。

## 多个dataset和他们的引用

可以同时定义多个dataset，系列可以通过series.datasetIndex来指定引用哪个dataset。

## Echarts3的数据设置方式仍正常使用

ECharts4之前一直以来的数据声明方式仍然被正常支持，如果系列已经声明了series.data，那么就会使用series.data而非dataset。

其实，series.data也是种会一直存在的重要设置方式。一些特殊的非table格式的图表，treemap，graph，lines等，现在仍不支持在dataset中设置，仍然需要使用series.data。另外，对于巨大数据量的渲染，需要使用appendData进行增量加载，这种情况下不支持使用dataset。