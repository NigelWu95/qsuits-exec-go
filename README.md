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
1. 下载完成后使用 `chmod +x qsuits_xxx_xxx ` 为其授予可执行权限【建议将该执行文件改名为 qsuits 后移动到系统环境变量的可执行文件路径下，便于从
任意位置调用 qsuits】。  
2. 如果运行出现 `line 7: syntax error near unexpected token 'newline'`, `line 7: '<!DOCTYPE html>'` 等这样的错误，表明下载的文件
不对，请重新从上述链接中下载。   

Usage（用法）:  
&ensp;&ensp;this tool is a agent program for qsuits, your local environment need java8 or above. In 
default mode, this tool will use latest java qsuits to exec, you only need use qsuits-java's parameters to run. If you 
use local mode with "-L" it mean you dont want to update latest qsuits automatically.  
&ensp;&ensp;这个工具是 qsuits-java 的一个代理程序，本地必须有 java8 或者 java8 以上的环境。在默认模式下，qsuits 运行时会去使用 qsuits-java 的最新版本来
执行操作，只需要传递 qsuits-java 所规定的参数即可。如果您使用 "-L" 参数表示您只想使用本地设置的默认版本，而不自动更新 qsuits-java 的最新版本。 

Options（选项）:  
&ensp;&ensp;&ensp;&ensp;&ensp; --Local/-L &ensp;&ensp;&ensp;&ensp;&ensp; use current default qsuits version to exec.
使用当前的默认 qsuits-java 版本来运行  
&ensp;&ensp;&ensp;&ensp;&ensp; --help/-h &ensp;&ensp;&ensp;&ensp;&ensp;&ensp; print usage. 打印用法说明  
Commands:  
&ensp;&ensp;&ensp;&ensp;&ensp; help &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; print usage. 打印使用帮助  
&ensp;&ensp;&ensp;&ensp;&ensp; upgrade &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; upgrade this own executable program by
 itself. 升级 qsuits 自身可执行程序  
&ensp;&ensp;&ensp;&ensp;&ensp; versions &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; list all qsuits versions from local.
 列举本地的所有 qsuits-java 的版本  
&ensp;&ensp;&ensp;&ensp;&ensp; clear &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; remove all old qsuits 
versions from local. 清除本地 qsuits-java 的所有旧版本  
&ensp;&ensp;&ensp;&ensp;&ensp; current &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp; query local default qsuits version.
查询 qsuits-java 本地当前的默认版本  
&ensp;&ensp;&ensp;&ensp;&ensp; chgver <no.> &ensp;&ensp;&ensp; set local default qsuits version. 设置本地的默认 qsuits-java 版本  
&ensp;&ensp;&ensp;&ensp;&ensp; download <no.> &ensp; download qsuits with specified version. 下载 qsuits-java 的指定版本  

Usage of qsuits-java（qsuits-java 的完整用法请参考）:  https://github.com/NigelWu95/qiniu-suits-java  
