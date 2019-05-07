# 通用脚本语言

UINX脚本程序解释执行用户的命令。命令可以让用户输入，也可以读取一个文件。Shell脚本是非编译的，解释执行的。除了向内核传递命令外，Shell主要任务是提供一个用户执行环节。

## Shell类型

UINX系统一般提供几种shell类型。

sh，最初的，基础的shell，现在一直在UNIX系统类UNIX系统中使用。虽然不是标准shell，liunx为了兼容unix，一般也提供该shell。

bash，标准的GNU shell。在linux上，bash是默认的shell。bash是sh的超集，增加了一些插件。bash可以兼容sh脚本。

csh，C Shell，脚本语法和C类似。

tcsh，TENEX C，csh的超集。提高了使用友好度和速度。

ksh，korn Shell，受到一些有unix背景的人赏识，sh的超集。配置复杂。

/etc/shells可以查看当前系统安装的shell。
```bash
cat /etc/shells
```
/etc/passwd文件记录着默认的shell。
```bash
cat /etc/passwd
```
输入shell的名字即可切换shell。
```bash
sh
bash
zsh
```

## bash的优点

GUN项目提供了免费的和UNIX系统类似的系统管理工具。

Bash是一个兼容sh的，符合IEEE POSIX P1003.2/ISO 9945.2 Shell标准的Shell。相比sh，提供了一系列增强的功能。

### 独有功能

Startup文件在脚本开始执行时会被执行。

login shell，login shell一般是用用户名和密码登录系统后得到的。
会先后执行/etc/profile，~/.bash_profile||~/.bash_login||~/.profile文件，~/.bash_logout会在logout之后执行。

non-login shell，non-login shell意味着你不必登录到系统。会执行~/.bashrc

bash为了兼容sh和符合POSIX规范，会读取/etc/profile和~/.profile文件。

interactive shells，一个交互的shell通常从用户终端中读取命令，写入输出。
```bash
 echo $-
 # 包含i则是交互模式，$-为使用set命名设置的flag
```
交互模式：会执行startup文件，工作管理默认开启，命名历史记录会开启(保存在$HISTFILE文件)，alias功能会开启，SIGTERM会被忽略，SIGINT会被捕捉和处理(不会退出交互shell)，SIGUP会停止交互式shell，