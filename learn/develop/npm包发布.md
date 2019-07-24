# 注册登录

在https://www.npmjs.com 官网注册用户后。

使用npm login登录。登录后使用npm whoami测试登录是否成功。

# package and modules

package是一个文件或一个文件夹。必须有package.json文件，package.json文件描述包信息。

一个package应该符合：

1、一个包含package.json的文件夹。

2、解压后可得到1的一个gzipped网络包。

3、一个gzipped网络包的URL。

4、registry中存在的nam@version。

5、name@tag,latest

6、git url。git://github.com/user/project.git#commit-ish。


Module是在node_modules文件夹下，可以被require方法加载。

为了被Node.js使用，module应该符合：

1、包含package.json的文件夹且含有main字段。

2、包含index.js的文件夹。

3、一个JS文件。

