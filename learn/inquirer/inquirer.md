# inquirer

## 简单介绍

Inquirer.js可以为Node.js提供一个简单易用的命令行。

Inquirer.js可以简化：

1、用户输入错误反馈

2、向用户提问

3、格式化用户输入

4、校验用户输入

5、分层提问

## 安装

```bash
npm i inquirer
```

```js
inquirer.prompt([{
    name: 'name',
    message: "What is your name?"
}]).then(answer => {
    // answer = { name : 'liuhao' };
    console.log('Your name is ' + answer.name + '!');
}).catch(err => {
    console.error(err);
});
```

## 方法

inquirer.prompt(questions) -> promise

questions: 问题数组

inquirer.registerPrompt(name,promt)

name，prompt的名字，一般用于问题type

prompt，prompt插件对象

inquirer.createPromptModule() -> prompt function

创建一个自己的inquirer module。避免影响其它lib。

## 类型

Question

type: string, 默认为input，可以包含input, confirm, list, rawlist, expand, checkbox, password,editor

