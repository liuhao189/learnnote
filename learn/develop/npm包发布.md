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

2、指明package的版本，定义版本最好依据版本规则。

3、方便其他开发者使用你的代码。

## Required name and version fileds

name，必须是小写字母，可以使用连字符或下划线。

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

# Creating Node.js modules

## main field

其它人加载你的模块时引入的文件是main属性指向的文件。

# README.md

为了给使用你发布的包的人一个良好的体验，建议书写readme文件在你的包根目录。

readme文件会在你的包页面进行展示。

# Publish unscoped public packages

unscoped的包都是public的，只用包名就可引用。

## reviewving package contents for sensitive or unnecessary information

注意避免敏感信息发布到npm包上。

低度敏感信息，可以使用.npmignore或.gitignore文件避免发布到npm上。

```bash
npm publish

https://npmjs.com/packages/package-name
```

# Keeping files out of your package

使用.npmignore文件使部分文件和文件夹不会包含到npm包中。

如果没有.npmignore文件，npm会使用.gitignore文件。

如果想使用.gitignore里的文件，需要创建.gitignore文件。

1、空白行或以#开头的行会忽略

2、支持标准的glob模式

3、/结尾可以指定一个文件夹。

4、!开始可以取反。

默认的，以下模式会被忽略：

1、.*.swp，._*，.DB_Store,.git,.hg,.npmrc,.lock-wsscript,.svn,.wafpickle-*,.config.gypi,CVS,.npm-delog.log。

node_modules也会被忽略。

以下文件永远不会被忽略。

1、package.json

2、README

3、CHNAGELOG

4、LICENSE/LICENCE

## Testing whether your .npmignore or files config works

npm pack会产生一个tarball，在当前目录。


## DevDependencies and Dependencies

当执行npm install，npm会下载dependencies和devDependencies。

可以使用smver calculator来查看要下载的具体版本。

1、dependencies，线上环境需要使用的依赖。

2、devDependencies，开发环境需要使用的依赖。

## Adding dist-tags to packages

发布标签是更易读易懂的文本。

npm publish会默认将你的包打上latest标签。如果要使用其他tag，需要使用--tag符号。

```bash
npm publish --tag beta
npm dist-tag add <package-name>@<version> tag
```

## 更新版本

可以使用npm version update_type 命名更新版本。

update_type必须是patch，major或minor。

## 从registry中删除包

在你发布包之后的72小时以内，你可以使用npm命令删除该npm包。

72小时以后，需要联系npm技术支持。

一旦删除npm包，npm包在24小时以内不能再次发布。需要换npm包名才可再次发布。

```bash
npm unpublish <package-name> -f
npm unpublish <package_name>@<version>
```

## deprecate and undeprecating

如果你不想再维护npm包的特定版本，你可以将该版本标记为deprecate。用户安装该版本时，会在命令行打印一条deprecate的消息。

### deprecate npm包

deprecate npm包，会让该包不会被搜索到。包主页会显示一条deprecate消息。

undeprecate npm包，使用npm deprecate <package_name> ""命令，message为空字符串即可。

    powershell好像不可以，cmd可以。

```bash
npm deprecate <package_name> "message"

# undeprecate

npm deprecate <package_name> ""
```

