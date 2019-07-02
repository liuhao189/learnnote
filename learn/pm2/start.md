# PM2

PM2是一个进程守护程序，可以帮助你管理和保持你的应用在线。作为简单直观的CLI工具提供。

## 安装 && 使用

最简单的方法是使用pm2的start命令启动应用程序。

--name <app_name>，--watch，--max-memory-restart ,--log, --arg1 arg2 arg3,--restart-delay, --time,
--no-autorestart, --corn, --no-daemon

很多选项可用，你可以根据不同的场景使用不同的选项。

```bash
npm install pm2@latest -g

pm2 start app.js
```

## Managing processes

管理应用状态一般使用以下命令，restart，reload，stop，delete。

除了app_name，还可以传递all，或id。

```bash
pm2 restart app_name
pm2 reload app_name
pm2 stop app_name
pm2 delete app_name
```

## Check Status, logs, metrics

你已经启动了应用，你可以检查应用状态，日志，应用指标。甚至可以使用pm2.io在网页端查看。

```bash
#list managed applications
pm2 list ls status
#display logs
pm2 logs
pm2 logs --lines 200
#terminal based daashboard
pm2 monit
#Monitoring && diagnostic web interface
pm2 plus
```

## Cluster mode

对于Nodejs应用，PM2包含一个自动的负载均衡功能。负载均衡会在衍生的进程中分享网络连接。

```bash
pm2 start app.js -i max
```

## Ecosystem File

你可以创建一个配置文件，配置文件叫做Ecosystem文件。

Ecosystem配置文件中可以设置name，scripts，args，instances，autorestart，env等变量。

```bash
pm2 ecosystem

pm2 start ecosystem.config.js
```

## setup satrtup script

启动PM2管理的进程们，在你的服务器重启后是比较关键的。

```bash
pm2 startup
pm2 save
```

## Restart application on changes

--watch选项。可以用--ignore-watch忽略部分文件夹。

```bash
pm2 start app.js --watch --ignore-watch="node_modules"
```

## Gracefule Stop

在pm2重启，重新加载，关闭进程时，确保你监听了SIGINT信号并清除了多余资源。
