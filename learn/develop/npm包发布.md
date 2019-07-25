# 注册登录

在https://www.npmjs.com 官网注册用户后。

使用npm login登录。登录后使用npm whoami测试登录是否成功。

# package and modules

package是一个文件或一个文件夹。必须有package.json文件，package.json文件描述包信息。

一个package应该符合：

1、一个包含package.json的文件夹。

2、解压后可得到1的一个gzipped网络包。

3、一个gzipped网络包的URL。

4、registry中存在的name@version。

5、name@tag,latest

6、git url。git://github.com/user/project.git#commit-ish。


Module是在node_modules文件夹下，可以被require方法加载。

为了被Node.js使用，module应该符合：

1、包含package.json的文件夹且含有main字段。

2、包含index.js的文件夹。

3、一个JS文件。

# Creating a package.json file

发布到registry的package必须包含package.json文件。

package.json文件：

1、列出package的依赖。

2、指明package的版本，定力版本最好依据版本规则。

3、方便其他开发者使用你的代码。

## Required name and version fileds

name，必须是小写字母，可以使用连字符或下换线。

version，必须是x.x.x，遵循软件版本规则。

## Other fields

author，Your Name<email@example.com> (http://example.com)

# Creating a new package.json file

## Running a CLI question

```bash
npm init
```
## Customizing the package.json questionnaire

npm init中可以添加自定义的问题和字段。

1、home文件下，创建一个.npm-init.js文件。

2、导出一个自定义字段的对象即可。

```bash
module.exports = {
    'customField':'Example custom field',
    'otherCustomField':'This example field is rally cool'
}
```

## creating a default package.json file

为了创建一个默认的package.json，使用npm init -y或--yes标记。

## Setting config options for the init command

```bash
npm set init.author.email "example-user@example.com"
npm set init.author.name  "example_user"
```