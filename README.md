# qsuits-exec-go  
qsuits-java 代理执行工具 by golang  

|操作系统|程序名|地址|
|---|-----|---|
|windows 32 位|qsuits_windows_386.exe|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_windows_386.exe)|
|windows 64 位|qsuits_windows_amd64.exe|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_windows_amd64.exe)|
|linux 32 位|qsuits_linux_386|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_linux_386)|
|linux 64 位|qsuits_linux_amd64|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_linux_amd64)|
|mac 32 位|qsuits_darwin_386|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_darwin_386)|
|mac 64 位|qsuits_darwin_amd64|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_darwin_amd64)|

**Tips：**  
1. 下载完成后 Linux 和 Mac 下需要使用 `chmod +x qsuits_xxx_xxx` 为其授予可执行权限【建议将该执行文件改名为 qsuits 后移动到系统环境变量的可
执行文件路径下，便于从任意位置调用 qsuits】。  
2. 如果运行出现 `line 7: syntax error near unexpected token 'newline'`, `line 7: '<!DOCTYPE html>'` 等这样的错误，表明下载的文件
不对，请重新从上述链接中下载。   
3. 由于需要 java 环境进行运行，如果系统环境变量中已安装 java 8 或以上，则直接会调用系统下的 java。如果您的登录账号没有安装系统程序的权限，或者系
统下的 java 版本在 8 以下，那么需要单独下载一个 8 以上的 jdk，安装或解压在自己的路径下面，然后可以通过 setjdk <jdkpath> 来设置该 jdk 环境，
设置后可以通过 -j/--java 来使用该 jdk，也可以直接 -j/--java <jdkpath> 来使用。  
4. java 安装简易[指南](https://blog.csdn.net/wubinghengajw/article/details/102612267): https://blog.csdn.net/wubinghengajw/article/details/102612267  

Usage of qsuits（用法）:  
&ensp;&ensp;&ensp;&ensp;this tool is a agent program for qsuits, your local environment need java8 or above. In 
default mode, this tool will use latest java qsuits to exec, you only need use qsuits-java's parameters to run. If you 
use local mode with "-L/--Local" it mean you dont want to update latest qsuits automatically.  
&ensp;&ensp;&ensp;&ensp;这个工具是 qsuits-java 的一个代理程序，本地必须有 java8 或者 java8 以上的环境。在默认模式下，qsuits 运行时会去使
用 qsuits-java 的最新版本来执行操作，只需要传递 qsuits-java 所规定的参数即可。如果您使用 "-L/--Local" 参数表示您只想使用本地设置的默认版本，
而不自动更新 qsuits-java 的最新版本。 

Options（选项）:  
&ensp;&ensp;&ensp;&ensp;&ensp; -h/--help &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; Print usage. 打印用法说明  
&ensp;&ensp;&ensp;&ensp;&ensp; -L/--Local &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; Use current default qsuits version to exec. 
使用当前的默认 qsuits-java 版本来运行  
&ensp;&ensp;&ensp;&ensp;&ensp; -j/--java \[<jdkpath>\] &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; Use custom jdk by existing setting or assigned <jdkpath>."
使用自定义 jdk 通过已存在的设置（指 setjdk 操作的设置）或者通过命令行 <jdkpath> 的设置，运行时指定的 <jdkpath> 会自动更新到 setjdk 中   

Commands:  
&ensp;&ensp;&ensp;&ensp;&ensp; help &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; Print usage. 打印使用帮助  
&ensp;&ensp;&ensp;&ensp;&ensp; selfupdate &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; Update this own executable program by
 itself. 升级 qsuits 自身可执行程序  
&ensp;&ensp;&ensp;&ensp;&ensp; versions &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; List all qsuits versions from local.
 列举本地的所有 qsuits-java 的版本  
&ensp;&ensp;&ensp;&ensp;&ensp; clear &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; Remove all old qsuits 
versions from local. 清除本地 qsuits-java 的所有旧版本  
&ensp;&ensp;&ensp;&ensp;&ensp; current &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; Query local default qsuits version.
查询 qsuits-java 本地当前的默认版本  
&ensp;&ensp;&ensp;&ensp;&ensp; chgver <no.> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp; Set local default qsuits version. 设置本地的默认 qsuits-java 版本  
&ensp;&ensp;&ensp;&ensp;&ensp; download <no.> &ensp;&ensp;&ensp;&ensp; Download qsuits with specified version. 下载 qsuits-java 的指定版本  
&ensp;&ensp;&ensp;&ensp;&ensp; update <no.> &ensp;&ensp;&ensp;&ensp;&ensp;&ensp; Update qsuits with specified version.
更新本地的 qsuits-java 版本，结合 download 和 chgver  
&ensp;&ensp;&ensp;&ensp;&ensp; setjdk \<jdkpath\> &ensp;&ensp;&ensp; Set jdk path as default.
设置默认 jdk 路径，后续操作通过 -j/--java 可以使用该 jdk 执行  

qsuits 完整文档见：[七牛开发者中心](https://developer.qiniu.io/kodo/tools/6263/the-command-line-tools-qsuits#1) / [github qsuit 文档](qsuits.md)  
Usage of qsuits-java（qsuits-java 的完整用法请参考）:  https://github.com/NigelWu95/qiniu-suits-java  
