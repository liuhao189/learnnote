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

name: string, answers中的属性名，如果包含“.”，会创建子对象。

message:string|Function问题名称，如果是一个函数，第一个参数会是当前会话的回答。

default:string|number|boolean|array|function, 默认值，函数第一个参数会是当前会话的回答。

choices:array|function，选择项数组或返回选择项数组的函数，选项可以是字符串，包含name、short和value的对象。

validate:function，接受用户的输入和会话内回答对象，校验通过返回true，校验失败返回message。

filter:function，接受用户输入，返回转换后的值。

transformer:Function, 接受用户输入，会话回答对象和选项，返回展示信息给用户。

when:function,boolean，接受用户的输入，返回是否避过这个问题。

pageSize: 展示的条数。

prefix: message前缀

suffix: message后缀

default，choices，validate，filter，when都可以是异步函数，返回Promise或者使用this.async。

new inquirer.Separator() 可以加到choices对象中，作为项的分隔符。

## 用户界面和布局

除了上文的内容，inquirer提供一些基本文本UI。



