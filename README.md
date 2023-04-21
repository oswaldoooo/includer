# Includer 使用教程
### **简介**
includer是用于c++的头文件管理的命令行工具。前提是包(包含头文件的目录)需要发布在github公开仓库上，当然你也可以拉取别人发布的开源包。
### **配置文件展示**
beta1.0 改动:配置文件中的包配置下可不包含头文件搜寻，在不包含头文件设置下，默认会将整个包链接到当前目录的lib中
```xml
<?xml version="1.0" encoding="UTF-8"?>
<includer>
    <!-- 包地址，可以有多个包，这里是单个包的实例，导入多个包则以此类推 -->
    <include_config packagename="github.com/oswaldoooo/ctools">
        <!-- 需要在包中搜索的头文件header 不止一个，因为你搜寻的头文件不止一个 -->
        <header name="arraylist.h"/>
        
    </include_config>
</includer>
```
### 命令行使用展示
```shell
#不加指定配置文件，则默认配置文件名为当前{目录名}.xml
includer init
#指定当前目录下配置文件,这里的exmaple.xml是示例，仅做参考
includer init example.xml
#在当前目录下生成配置文件模版,不指定文件名则默认以目录名为名生成
includer generate
#查看配置文件中的头文件是否覆盖(测试功能)
includer reload
#通过命令行给配置文件添加包,（必须指定配置文件名)
includer package add packagename -c filename
#加载全部包,统一其中的第三方头文件引入路径不统一的问题
includer load
#加载指定的包(测试功能)
includer load packagename
```
### 环境配置
请将下列配置加入你的相关profile文件，以确保其生效。你也可使用快速安装[脚本](https://brotherhoodhk.org/products/shells/includer_installer.sh)
```shell
export INCLUDER_HOME=你的includer存放地址
export PATH=$PATH:$INCLUDER_HOME
```
快速安装脚本
```shell
curl https://brotherhoodhk.org/products/shells/includer_installer.sh bash
```
验证是否配置成功
```shell
includer version
```