# DllPlugin && DllReferencePlugin

DllPlugin和DllReferencePlugin结合起来使用可以大大提高打包构建速度。

## DllPlugin

DllPlugin使用一个单独的webpack配置生成dll包，同时会生成一个manifest.json文件。

DllPlugin参数：

1、name，dll的命名规则；

2、path，manifest.json文件路径；

3、context(可选参数)。

manifest.json文件包含require和import的请求与module ids之间的关系。

## DllReferencePlugin

DllReferencePlugin用在基本的webpack配置中。

DllReferencePlugin参数：

1、context，绝对路径

2、manifest，manifest文件路径

3、content，可选

4、name，dll名字

5、scope，可选

6、sourceType，可选，dll的类型

## Modes

DllPlugin可以以两种不同方式使用。

### Scoped Mode

Dll中的内容通过模块前缀访问。

    scope='xyz' 
    // a file 'abc' in the dll 
    require('xyz/abc')

### Mapped Mode

Dll中的内容映射到文件系统目录。

## Demo


