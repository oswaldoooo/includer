# Includer 使用教程
### **简介**
includer是用于c++的头文件管理的命令行工具。前提是包(包含头文件的目录)需要发布在github公开仓库上，当然你也可以拉取别人发布的开源包。
### **配置文件展示**

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
```
### 环境配置
请将下列配置加入你的相关profile文件，以确保其生效。你也可使用快速安装[脚本](https://brotherhoodhk.org/products/shell/includer_installer.sh)
```shell
export INCLUDER_HOME=你的includer存放地址
export PATH=$PATH:$INCLUDER_HOME
```
快速安装脚本
```shell
curl https://brotherhoodhk.org/products/shells/includer_installer.sh
```
验证是否配置成功
```shell
includer version
```