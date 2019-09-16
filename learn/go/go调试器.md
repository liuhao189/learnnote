# dlv

Delve是一个Go编程语言的调试器。

Delve可以让你以交互式的方式控制程序执行，计算变量和表达式，提供线程，goroutine，CPU等信息。

# dlv调试代码

1、执行dlv debug命令。dlv会去编译你的代码，会传一些参数给编译器，好让编译器编译出来更加方面调试的可执行文件，然后启动你的程序，并且attch上去。命令行会停留在debug session。

```bash
dlv debug test-debug.go
```

2、设置一个端点。break main.main，在main包的main函数设置一个断点。输出信息告诉我们断点的ID和断点的位置，函数名和文件名以及所在的行数。

```bash
break main.main
```

3、continue命令让程序运行到我们打断点的地方。

```bash
continue
```

4、next命令让程序运行到下一句话，向继续向下，可以直接按回车。delve会重复上一条命令。

```bash
next or enter key
```

5、print看一下变量或表达式的信息。

```bash
print ctx
print ctx.request
```