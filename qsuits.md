# 简介

qsuits 是一个围绕七牛云服务 API 设计开发的高效便捷的多线程工具，能够高效**并发列举云存储空间**的资源列表，支持包括**七牛云/阿里云/腾讯云/AWS S3/又拍云/华为云/百度云等**不同数据源的云存储空间列举和[数据备份](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datamigration.md)，支持 S3 接口的数据源均可以通过 S3 数据源的方式来导出。同时支持对本地提供的资源列表并发进行批量处理，API 操作主要包括本地文件批量上传和对七牛云存储资源进行增/删/改/查/迁移/转码/内容审核等。该工具非常适合于存储空间文件数量较多场景下的列举，以及针对空间文件直接做过滤和 API 操作，也可以方便地进行不同数据源公开或私有空间的数据备份和迁移。除了批处理之外，该工具同时提供了针对即时输入的交互式操作和单次操作模式。该工具实际功能基于 Java8 编写，可基于 jdk8 环境在命令行或 ide 中运行，但是为了给开发者提供更好的使用方式，同样提供了基于 go 编写的命令行[执行器](https://github.com/NigelWu95/qsuits-exec-go)，如果您有使用上的问题或者建议，以及任何被认为适合加到该工具中的新需求，都可以联系七牛技术支持或在[ ISSUE 列表](https://github.com/NigelWu95/qiniu-suits-java/issues)里进行反馈，我们会协助您更好地使用它。  

源码地址：  
GitHub 项目 [qiniu-suits-java](https://github.com/NigelWu95/qiniu-suits-java)
GitHub 项目 [qsuits-exec-go](https://github.com/NigelWu95/qsuits-exec-go)

# 执行器下载

qsuits 执行器使用 Go 语言编写而成，提供了预先编译好的各主流操作系统平台的二进制可执行文件供下载，便于命令行直接使用，由于 java 也是一门跨平台的语言，您可以在 Linux/Mac 或是 Windows 先安装好 java 环境，如果您的本地环境中尚未包含 jdk，该工具将会提示或帮助您去安装它，在以下的文档中，我们统一使用 `qsuits` 这个命令来做介绍。

> 原 java 项目的发布历史 [查看](https://github.com/NigelWu95/qiniu-suits-java/releases)

|操作系统|程序名|地址|
|---|-----|---|
|windows 32 位|qsuits_windows_386.exe|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_windows_386.exe)|
|windows 64 位|qsuits_windows_amd64.exe|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_windows_amd64.exe)|
|linux 32 位|qsuits_linux_386|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_linux_386)|
|linux 64 位|qsuits_linux_amd64|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_linux_amd64)|
|mac 32 位|qsuits_darwin_386|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_darwin_386)|
|mac 64 位|qsuits_darwin_amd64|[下载](https://github.com/NigelWu95/qsuits-exec-go/raw/master/bin/qsuits_darwin_amd64)|


# 安装使用

qsuits 执行器是一个命令行工具，下载后即可命令行运行，建议重命名为 qsuits 并加入环境变量中使用，java 需要作为平台软件进行安装，以下进行相关的安装说明和指导：


**Linux和Mac平台**

（1）权限
如果在 Linux 或者 Mac 系统上遇到 `Permission Denied` 的错误，请使用命令 `chmod +x qsuits` 来为文件添加可执行权限。

（2）任何位置运行
对于 Linux 或者 Mac，如果希望能够在任何位置都可以执行，那么可以把 `qsuits` 所在的目录加入到环境变量 `$PATH` 中去。假设 `qsuits` 命令被解压到路径 `/home/nigel/Work/Tools` 目录下面，那么我们可以把如下的命令写入到你所使用的 bash 所对应的配置文件中，如果是 `/bin/bash`，那么可以是 `~/.bashrc` 文件，如果是 `/bin/zsh`，那么可以是 `~/.zshrc` 文件中。写入的内容为：

```
export PATH=$PATH:/home/nigel/Work/Tools
```
保存完毕之后，可以通过两种方式立即生效，其一为输入 `source ~/.zshrc` 或者 `source ~/.bashrc` 来使配置立即生效，或者完全关闭命令行，然后重新打开一个即可，接下来就可以在任何位置使用 `qsuits` 命令了。

**Windows平台**

（1）闪退问题
本工具是一个命令行工具，在 Windows 下面请先打开命令行终端，然后输入工具名称执行，不要双击打开，否则会出现运行后直接退出现象。

（2）任何位置运行
如果你希望可以在任意目录下使用 `qsuits`，请将 `qsuits` 工具可执行文件所在目录添加到系统的环境变量中。由于 Windows 系统是图形界面，假设 `qsuits.exe` 命令被解压到路径 `D:\Work\Tools` 目录下面，那么我们把这个目录放到系统的环境变量 `PATH` 里面，操作如下所示。

![](http://qsuits.nigel.net.cn/windows_user_path.png)

（3）文件路径和字符编码问题

- Windows 平台下的文件路径需要写为 `\\` 的写法，如 `C:\\Users\\nigel\\Downloads`。而 Linux 和 Mac 平台下的文件路径为 `/` 的写法，如 `/Users/nigel/Downloads`。
 
- 在使用命令和配置文件时，尽量使用 `""` 双引号，配置文件使用 `utf-8` 编码。

(4) java 安装

qsuits 首次执行在无 java 环境的系统上时会提示您做下载或安装，类似：

![](http://qsuits.nigel.net.cn/choose_install_java_message.jpg)  

输入 "yes" 后会下载适合您系统的 jdk8，然后可根据下载的文件执行安装，由于 java 的安装在不同平台下有差异，具体安装指南可参考 [java 安装指南](https://blog.csdn.net/wubinghengajw/article/details/102612267)

# 账号设置
该工具的一些操作需要使用到密钥，为了不在每次使用时明文输入密钥，该工具提供了账号的设置（经过加密）。因为该工具支持不同数据源的读取，因此也支持设置不同数据源的账号，设置之后直接通过账号名 account name 来进行操作即可，定义不同的 account name 则可设置多对密钥，亦可设置不同数据源的账号密钥，同一数据源的账号名相同时会覆盖该数据源账号的历史密钥，命令行操作如下所示（采用配置文件也可以进行账户设置和使用，命令行参数名前需要加上 `-` 符号而配置文件中参数设置不需要，每项参数一行即可，可参考后面的配置文件使用方式），密钥参数名参考[各存储数据源配置参数](#各存储数据源配置参数)。  

## 1. 设置 account：  
命令格式：`-account=<source>-<name> -<source>-id= -<source>-secret= [-d]`，如：  

`-account=test/qiniu-test -ak= -sk=` 设置七牛账号，账号名为 test，没有数据源标识时默认设置七牛账号  
`-account=ten-test -ten-id= -ten-secret=` 设置腾讯云账号，账号名为 test  
`-account=ali-test -ali-id= -ali-secret=` 设置阿里云账号，账号名为 test  
`-account=s3-test -s3-id= -s3-secret=` 设置 AWS/S3 账号，账号名为 test  
`-account=up-test -up-id= -up-secret=` 设置又拍云账号，账号名为 test  
`-account=hua-test -hua-id= -hua-secret=` 设置华为云账号，账号名为 test  
`-account=bai-test -bai-id= -bai-secret=` 设置百度云账号，账号名为 test  
`-d` 表示默认账号选项，此时设置的账号将会成为全局默认账号，执行操作时 -d 选项将调取该默认账号  

## 2. 使用 account 账号：  
`-a=test` 表示使用 test 账号，数据源会自动根据 path 参数判断  
`-d` 表示使用默认的账号，数据源会自动根据 path 参数判断  

## 3. 查询 account 账号：
命令格式：`-getaccount=<source>-<name> [-dis] [-d]`，默认只显示 id 的明文而隐藏 secret，`-dis` 参数表示选择明文显示 secret，如：  

`-getaccount -d` 表示查询设置的默认账号的密钥  
`-getaccount=test -dis` 表示查询设置的所有账号名为 test 的密钥，并显示 secret 的明文  
`-getaccount=s3-test` 表示查询设置的 S3 账号名为 test 的密钥  
`-getaccount=ten-test` 表示查询设置的腾讯账号名为 test 的密钥  
`-getaccount=qiniu-test` 表示查询设置的七牛账号名为 test 的密钥  

## 4. 删除 account 账号：  
命令格式：`-delaccount=<source>-<name>`，删除账号只允许一次删除一条，如：  
`-delaccount=s3-test` 表示删除设置的 S3 账号名为 test 的密钥  
`-delaccount=ten-test` 表示删除设置的腾讯账号名为 test 的密钥  
`-delaccount=test/qiniu-test` 表示删除设置的七牛账号名为 test 的密钥  

# 命令选项

执行器为了更好的管理 qsuits 的使用，还提供了一些执行选项和自定义命令：

## Options（选项）:

    -h/--help            打印用法说明
    -L/--Local           使用当前的默认 qsuits-java 版本来运行，即不从网上更新最新版本
    -j/--java jdkpath    使用自定义 jdk 通过已存在的设置（指 setjdk 操作的设置）或者通过命令行的设置，运行时指定的 会自动更新到 setjdk 中

## Commands（自定义命令）：

|命令名称         |作用描述 |
|----------------|-----|
|help            |打印使用帮助|
|selfupdate      |升级 qsuits 自身可执行程序|
|versions        |列举本地的所有 qsuits-java 的版本|
|clear           |清除本地 qsuits-java 的所有旧版本|
|current         |查询 qsuits-java 本地当前的默认版本|
|download <no.>  |下载 qsuits-java 的指定版本，<no.> 为版本号，如 8.1.0|
|chgver <no.>    |设置本地的默认 qsuits-java 版本，<no.> 为版本号，如 8.1.0|
|update <no.>    |更新本地的 qsuits-java 版本（结合 download 和 chgver），<no.> 为版本号，如 8.1.0|
|setjdk <jdkpath>|设置默认 jdk 路径，后续操作通过 -j/--java 可以使用该 jdk 执行，没有该项设置的情况下，会自动使用系统环境变量中的默认 jdk|

# 帮助手册

|命令名称         |作用描述 |
|----------------|-----|
|accounthelp     |打印 account 设置帮助|
|storagehelp     |打印 storage 数据源使用帮助|
|filehelp        |打印 file 数据源使用帮助|
|filterhelp      |打印 filter 数据过滤器使用帮助|
|processhelp     |打印 process 数据操作使用帮助|
	
# 配置文件

配置文件中可设置形如<参数名>=<参数值>，每行一个参数，如下所示表示从七牛的空间列举全部文件，并默认将结果列表保存在当前路径下创建的 \<bucket\> 目录下：  
```
path=qiniu://<bucket>
ak=
sk=
```  
**备注1**：支持通过默认配置文件来设置运行参数，默认配置文件路径为 `resources/application.config` 或 `resources/application.properties`，properties 方式需要遵循 java 的转义规则，两个文件存在任意一个均可作为默认配置文件来设置参数（优先使用 resources/application.properties），此时则不需要通过 `-config=` 指定配置文件路径，指定 `-config=` 时则默认文件路径无效。  
**备注2**：直接使用命令行传入参数（较繁琐），不使用配置文件的情况下全部所需参数可以完全从命令行指定，形式为：**`-<key>=<value>`**，**请务必在参数前加上 `-`**，如果参数值中间包含空格，请使用 `-<key>="<value>"` 或者 `-<key>='<value>'` 如  
```
qsuits -path=qiniu://<bucket> -ak=<ak> -sk=<sk> -f-date-scale="[2019-09-01, 2019-10-01 10:00:00]"
```  
**备注3**：命令行参数与配置文件参数可同时使用，参数名相同时命令行参数值会覆盖配置文件参数值，且为默认原则。**【推荐使用配置文件方式，
一是安全性，二是参数历史可保留且修改方便；推荐使用 -account 提前设置好账号，安全性更高，使用时 -a=\<account-name\> 即可，不必再暴露密钥】**，使用配置文件来设置密钥每次设置一个账号，如配置文件：
```
account=test
ak=
sk=
```
命令行执行 `qsuits -config=<config-path>` 即可设置账号完成，`account` 参数也可以放在命令行，执行 `qsuits -account=test -config=<config-path>` 即可。 因为默认数据源为 `qiniu`，所以设置时 account 时没有任何数据源信息则设置为七牛的账号。

# 运行模式
### （1）批处理模式：[读取[数据源](#数据源)] => [选择[过滤器](#数据过滤)] => [数据[处理过程](#数据操作)] => [[结果持久化](#结果持久化)]   
### （2）交互模式：从命令行输入数据时，process 支持[交互模式](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/interactive.md)运行，一次启动，可无限次命令行输入 data，输入一次处理一次并返回结果。  
### （3）单行模式：从命令行输入数据时，process 支持[单行模式](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/single.md)运行，一次启动，指定 data 参数，直接一次处理并返回结果。  

# 数据源
数据源分为三种类型：**云存储列举(storage)**、**文本文件行读取(file)**、**文件路径和属性读取(filepath)**，可以通过 `path=` 来指定数据源地址：  
`path=qiniu://<bucket>` 表示从七牛存储空间列举出资源列表，参考[七牛数据源示例](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#1-七牛云存储)  
`path=tencent://<bucket>` 表示从腾讯存储空间列举出资源列表，参考[腾讯数据源示例](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#2-腾讯云存储)  
`path=aliyun://<bucket>` 表示从阿里存储空间列举出资源列表，参考[阿里数据源示例](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#3-阿里云存储)  
`path=s3://<bucket>` 表示从 aws/s3 存储空间列举出资源列表，参考[S3数据源示例](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#4-aws-s3)  
`path=upyun://<bucket>` 表示从又拍云存储空间列举出资源列表，参考[又拍数据源示例](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#5-又拍云存储)  
`path=huawei://<bucket>` 表示从华为云存储空间列举出资源列表，参考[华为数据源示例](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#6-华为云存储)  
`path=baidu://<bucket>` 表示从百度云存储空间列举出资源列表，参考[百度数据源示例](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#7-百度云存储)  
`path=<path>` 表示从本地目录（或文件）中读取资源列表，参考[本地文件数据源示例](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#8-local-files)  
未设置数据源时则默认从七牛空间进行列举，数据源详细参数配置和说明及可能涉及的高级用法见：[数据源配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md)，配置文件示例可参考
[配置模板](https://github.com/NigelWu95/qiniu-suits-java/blob/master/resources/application.config)  

## 1. 公共参数
```
path=
unit-len=
threads=
indexes=key,etag,fsize
```  
|参数名|参数值及类型 |含义|  
|-----|-------|-----|  
|path| 数据源字符串| 选择从[云存储空间列举]还是从[本地路径文件中读取]资源，本地数据源时填写本地文件或者目录路径，云存储数据源时可填写 "qiniu://<bucket\>"、"tencent://<bucket\>" 等|  
|unit-len| 整型数字| 表示一次读取的文件个数（读取或列举长度，不同数据源有不同默认值），对应与读取文件时每次处理的行数或者列举请求时设置的 limit 参数|  
|threads| 整型数| 表示预期最大线程数，若实际得到的文件数或列举前缀数小于该值时以实际数目为准|  
|indexes| 字符串列表| 资源元信息字段索引（下标），设置输入行对应的元信息字段下标|  

**备注：** indexes、unit-len、threads 均有默认值非必填，indexes 说明及默认值参考下述[ indexes 索引](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#%E5%85%B3%E4%BA%8E-indexes-%E7%B4%A2%E5%BC%95)，unit-len 和 threads 说明及默认值参考下述[并发处理](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#%E5%85%B3%E4%BA%8E%E5%B9%B6%E5%8F%91%E5%A4%84%E7%90%86)，建议根据需要优化参数配置。  

## 2. storage 云存储列举  
```
<密钥配置>
region=
bucket=
marker=
start=
end=
prefixes=
anti-prefixes=
prefix-left=
prefix-right=
```  

|参数名|参数值及类型 |含义|  
|-----|-------|-----|  
|<密钥配置>|字符串|密钥对字符串|  
|region|字符串|存储区域|
|bucket|字符串| 需要列举的空间名称，通过 "path=qiniu://<bucket>" 来设置的话此参数可不设置，设置则会覆盖 path 中指定的 bucket 值|  
|prefixes| 字符串| 表示只列举某些文件名前缀的资源，，支持以 `,` 分隔的列表，如果前缀本身包含 `,\=` 等特殊字符则需要加转义符，如 `\,`|  
|prefix-config| 字符串| 该选项用于设置列举前缀的配置文件路径，配置文件格式为 json，参考[ prefix-config 配置文件](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#%E6%95%B0%E6%8D%AE%E6%BA%90%E5%AE%8C%E5%A4%87%E6%80%A7%E5%92%8C%E5%A4%9A%E5%89%8D%E7%BC%80%E5%88%97%E4%B8%BE)|
|anti-prefixes| 字符串| 表示列举时排除某些文件名前缀的资源，支持以 `,` 分隔的列表，特殊字符同样需要转义符|  
|prefix-left| true/false| 当设置多个前缀时，可选择是否列举所有前缀 ASCII 顺序之前的文件|  
|prefix-right| true/false| 当设置多个前缀时，可选择是否列举所有前缀 ASCII 顺序之后的文件|  

支持从不同的云存储上列举出空间文件，默认线程数(threads 参数)为 50，千万以内文件数量可以不增加线程，建议阅读[并发列举](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#并发列举)参考参数设置优化列举效率，通常云存储空间列举的必须参数包括密钥（可通过 account 方式来使用，下述为了直接表示密钥参数名所以没有体现 account 用法）、空间名（通过 path 或 bucket 设置）及空间所在区域(通过 region 设置，允许不设置的情况下表明支持自动查询)。

### 各存储数据源配置参数：  

|storage 源|             密钥和 region 字段         |                  对应关系和描述               |  
|---------|---------------------------------------|---------------------------------------------|  
|qiniu    |`ak=`<br>`sk=`<br>`region=z0/z1/z2/...`|密钥对为七牛云账号的 AccessKey 和 SecretKey<br>region使用简称(可不设置)，参考[七牛 Region](https://developer.qiniu.com/kodo/manual/1671/region-endpoint)|  
|tencent  |`ten-id=`<br>`ten-secret=`<br>`region=ap-beijing/...`| 密钥对为腾讯云账号的 SecretId 和 SecretKey<br>region使用简称(可不设置)，参考[腾讯 Region](https://cloud.tencent.com/document/product/436/6224)|  
|aliyun   |`ali-id=`<br>`ali-secret=`<br>`region=oss-cn-hangzhou/...`| 密钥对为阿里云账号的 AccessKeyId 和 AccessKeySecret<br>region使用简称(可不设置)，参考[阿里 Region](https://help.aliyun.com/document_detail/31837.html)|  
|aws/s3   |`s3-id=`<br>`s3-secret=`<br>`region=ap-east-1/...`| 密钥对为 aws/s3 api 账号的 AccessKeyId 和 SecretKey<br>region使用简称(可不设置)，参考[ AWS Region](https://docs.aws.amazon.com/zh_cn/general/latest/gr/rande.html)|  
|upyun    |`up-id=`<br>`up-secret=`<br>| 密钥对为又拍云存储空间授权的[操作员](https://help.upyun.com/knowledge-base/quick_start/#e6938de4bd9ce59198)和其密码，又拍云存储目前没有 region 概念|  
|huawei   |`hua-id=`<br>`hua-secret=`<br>`region=cn-north-1/...`| 密钥对为华为云账号的 AccessKeyId 和 SecretAccessKey<br>region(可不设置)使用简称，参考[华为 Region](https://support.huaweicloud.com/devg-obs/zh-cn_topic_0105713153.html)|  
|baidu    |`bai-id=`<br>`bai-secret=`<br>`region=bj/gz/su...`| 密钥对为百度云账号的 AccessKeyId 和 SecretAccessKey<br>region(可不设置)使用简称，参考[百度 Region](https://cloud.baidu.com/doc/BOS/s/Ojwvyrpgd#%E7%A1%AE%E8%AE%A4endpoint)|  

## 3 file 文本文件行读取
```
parse=tab/json
separator=\t
# 文件内容读取资源列表时一般可能需要设置 indexes 参数（默认只包含 key 字段的解析）
indexes=
add-keyPrefix=
rm-keyPrefix=
line-config=
```
|参数名|参数值及类型 |含义|  
|-----|-------|-----|  
|parse| 字符串 json/tab/csv| 数据行格式，json 表示使用 json 解析，tab 表示使用分隔符（默认 `\t`）分割解析，csv 表示使用 `,` 分割解析|  
|separator| 字符串| 当 parse=tab 时，可另行指定该参数为格式分隔符来分析字段|  
|add-keyPrefix| 字符串|将解析出的 key 字段加上指定前缀再进行后续操作，用于输入 key 可能比实际空间的 key 少了前缀的情况，补上前缀才能获取到资源|  
|rm-keyPrefix| 字符串|将解析出的 key 字段去除指定前缀再进行后续操作，用于输入 key 可能比实际空间的 key 多了前缀的情况，如输入行中的文件名多了 `/` 前缀|  
|line-config| 配置文件路径|表示从该配置中读取文件名作为 file 数据源，同时文件名对应的值表示读取该文件的起始位置，配置文件格式为 json，参考[ line-config 配置文件](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#line-config-%E9%85%8D%E7%BD%AE)|  
**数据源详细参数配置和说明及可能涉及的高级用法见：[数据源配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md)**  

## 4. filepath 文件路径和属性读取
该数据源用于上传文件的操作，设置 `process=qupload` 时自动生效，从 `path` 中读取所有文件（除隐藏文件外）执行上传操作，具体配置可参考[ qupload 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/uploadfile.md)。

# 数据过滤  
从数据源输入的数据通常可能存在过滤需求，如过滤指定规则的文件名、过滤时间点或者过滤存储类型等，可通过配置选项设置一些过滤条件，目前支持两种过滤条件：**基本字段过滤**和**特殊特征匹配过滤**。

## 1. 基本字段过滤  
根据设置的字段条件进行筛选，多个条件时需同时满足才保留，若存在记录不包该字段信息时则正向规则下不保留，反正规则下保留，字段包含：  
`f-prefix=` &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;表示**选择**文件名符合该前缀的文件  
`f-suffix=` &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;表示**选择**文件名符合该后缀的文件  
`f-inner=` &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;表示**选择**文件名包含该部分字符的文件  
`f-regex=` &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;表示**选择**文件名符合该正则表达式的文件，所填内容必须为正则表达式  
`f-mime=` &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;表示**选择**符合该 mime 类型（也即 content-type）的文件  
`f-type=` &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;表示**选择**符合该存储类型的文件，参下述[关于 f-type](#（1）关于-f-type)|  
`f-status=` &ensp;&ensp;&ensp;&ensp;&ensp;&ensp;&ensp;表示**选择**符合该存储状态的文件, 为 0（启用） 或 1（禁用）  
`f-date-scale` &ensp;&ensp;&ensp;&ensp;设置过滤的时间范围，格式为 [\<date1\>,\<date2\>]，\<date\> 格式为：yyyy-MM-DD HH:MM:SS，可参考：[特殊规则](#（3）f-date-scale)  
`f-anti-prefix=` &ensp;&ensp;表示**排除**文件名符合该前缀的文件  
`f-anti-suffix=` &ensp;&ensp;表示**排除**文件名符合该后缀的文件  
`f-anti-inner=` &ensp;&ensp;&ensp;表示**排除**文件名包含该部分字符的文件  
`f-anti-regex=` &ensp;&ensp;&ensp;表示**排除**文件名符合该正则表达式的文件，所填内容必须为正则表达式  
`f-anti-mime=` &ensp;&ensp;&ensp;&ensp;表示**排除**该 mime 类型的文件  

### （1）关于 f-type
|存储源|type 参数类型|具体值                   |
|-----|-----------|------------------------|
|七牛  | 整型      |0 表示标准存储；1 表示低频存储|
|其他  | 字符串     |如：Standard 表示标准存储  |  

### （2）特殊字符
特殊字符包括: `, \ =` 如有参数值本身包含特殊字符需要进行转义：`\, \\ \=`  

### （3）f-date-scale
\<date\> 中的 00:00:00 为默认值可省略，无起始时间则可填 [0,\<date2\>]，结束时间支持 now 和 max，分别表示到当前时间为结束或无结束时间。如果使用命令行来设置，注意日期值包含空格的情况（date 日期和时刻中间含有空格分隔符），故在设置时需要使用引号 `'` 或者 `"`，如 `-f-date-scale="[0,2018-08-01 12:30:00]"`，配置文件则不需要引号。  

## 2. 特殊特征匹配过滤 f-check[-x]  
根据资源的字段关系选择某个特征下的文件，目前支持 `ext-mime` 检查，程序内置的默认特征配置见：[check 默认配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/resources/check.json)，运行
参数选项如下：  
`f-check=ext-mime` 表示进行**后缀名 ext** 和 **mimeType**（即 content-type）匹配性检查，不符合规范的疑似异常文件将被筛选出来  
`f-check-config` 自定义资源字段规范对应关系列表的配置文件，格式为 json，自定义规范配置 key 字段必填，其元素类型为列表 [], 否则无效，如
`ext-mime` 配置时后缀名和 mimeType 用 `:` 组合成字符串成为一组对应关系，写法如下：  
```
{
  "ext-mime": [
    "mp5:video/mp5"
  ]
}
```  
配置举例：[check-config 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/resources/check-config.json)  
`f-check-rewrite` 是否覆盖默认的特征配置，为 false（默认）表示将自定义的规范对应关系列表和默认的列表进行叠加，否则程序内置的规范对应关系将失效，只检查自定义的规范列表。设置了过滤条件的情况下，后续的处理过程会选择满足过滤条件的记录来进行，或者对于数据源的输入进行过滤后的记录可以直接持久化保存结果，如通过 qiniu 源获取文件列表过滤后进行保存，可设置 `save-total=true/false` 来选择是否将列举到的完整记录进行保存。  
filter 详细配置可见[ filter 配置说明](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/filter.md)  

## 过滤示例
1、过滤 `test` 目录下 mime（content-type）为 video 类型的文件：  
```
f-prefix=test/
f-mime=video/
```
2、过滤 2019 年 10 月 1 日 10 点之前的 .jpg 文件：  
```
f-suffix=.jpg
f-date-scale=0,2019-10-01 10:00:00
```

# 数据操作
process 处理过程表示对数据源输入的每一条记录进行处理的类型，所有处理结果保存在 save-path 或默认路径下，具体处理过程由处理类型参数指定，如 **process=type/status/lifecycle/copy** (命令行方式则指定为 **-process=xxx**) 等，同时 process 操作支持设置公共参数：  
`retry-times=` 操作失败（可重试的异常情况下，如请求超时）需要进行的重试次数，默认为 5 次  
`batch-size=` 支持 batch 操作时设置的一次批量操作的文件个数（支持 batch 操作：type/status/lifecycle/delete/copy/move/rename/stat，其他操作请勿设置 batchSize 或者设置为 0），当响应结果较多 429/573 状态码时需要降低 batch-size，或者直接使用非 batch 方式：batch-size=0/1  

**处理操作类型：**  

|处理类型配置          |含义                                         |文档链接                                    |
|--------------------|--------------------------------------------|-------------------------------------------|
|process=qupload      | 表示上传文件到存储空间                        | [qupload 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/uploadfile.md)         |
|process=delete       | 表示删除空间资源                             | [delete 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/delete.md)              |
|process=copy         | 表示复制资源到指定空间                        | [copy 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/copy.md)                  |
|process=move         | 表示移动资源到指定空间                        | [move 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/move.md)                  |
|process=rename       | 表示对指定空间的资源进行重命名                 | [rename 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/rename.md)              |
|process=stat         | 表示查询空间资源的元信息                      | [stat 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/stat.md)                  |
|process=type         | 表示修改空间资源的存储类型（低频/标准）         | [type 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/type.md)                  |
|process=status       | 表示修改空间资源的状态（启用/禁用）             | [status 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/status.md)              |
|process=lifecycle    | 表示修改空间资源的生命周期                     | [lifecycle 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/lifecycle.md)        |
|process=mirror       | 表示对设置了镜像源的空间资源进行镜像更新         | [mirror 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/mirror.md)              |
|process=asyncfetch   | 表示异步抓取资源到指定空间                     | [asyncfetch 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/asyncfetch.md)      |
|process=qhash        | 表示查询资源的 qhash                         | [qhash 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/qhash.md)                |
|process=avinfo       | 表示查询空间资源的视频元信息                    | [avinfo 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/avinfo.md)             |
|process=pfopcmd      | 表示根据音视频资源的 avinfo 信息来生成转码指令   | [pfopcmd 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/pfopcmd.md)           |
|process=pfop         | 表示对空间资源执行 pfop 请求                   | [pfop 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/pfop.md)                 |
|process=pfopresult   | 表示通过 persistentId 查询 pfop 的结果        | [pfopresult 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/pfopresult.md)     |
|process=privateurl   | 表示对私有空间资源进行私有签名                  | [privateurl 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/privateurl.md)     |
|process=exportts     | 表示对 m3u8 的资源进行读取导出其中的 ts 文件列表 | [exportts 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/exportts.md)         |
|process=download     | 表示通过 http 下载资源到本地                   | [download 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/downloadfile.md)     |
|process=imagecensor  | 表示图片类型资源内容审核                       | [imagecensor 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/censor.md#图片审核)|
|process=videocensor  | 表示视频类型资源内容审核                       | [videocensor 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/censor.md#视频审核)|
|process=censorresult | 表示内容审核结果查询                          | [censorresult 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/censorresult.md)  |
|process=mime         | 修改资源的 mimeType                          | [mime 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/mime.md)                 |
|process=metadata     | 修改资源的 metadata                          | [metadata 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/metadata.md)         |
**注意**：
1、云存储数据源 + process 操作的情况下通常会涉及两对密钥，数据源一对，process 操作一对，如果是 delete、status 等操作则这两对密钥相同，使用一个密钥设置或者一个 account (`-a=<account-name>`) 即可，copy、move 要求针对同一个账号操作或者采用空间授权，因此也只需要一堆密钥，但如果是其他存储数据源的数据备份操作 asyncfetch，就需要两对不同的密钥，而 account 只支持设置一个，这时第二对的七牛密钥可以通过同一个 account-name 的设置来获得，因为同一个 account-name 可以为不同数据源做密钥设置，如：`-account=ali-test -ali-id= -ali-secret=` 设置了阿里云 test 名称的账号，同时 `-account=qiniu-test -ak= -sk=` 设置了七牛 test 名称的账号，则通过 `-a=test` 可以同时拿到阿里云和七牛云的 test 账号，因此可以直接通过同一个 account-name 来进行操作。但是如果明确指定了另外的 ak，sk，则会使用您设置的这一对七牛密钥。
2、也真是因为不同数据源的 account-name 可同名特性，以及支持主动设置密钥来覆盖 account 的密钥，在具体操作时需要注意账号和密钥的使用，以免对另外一个账号执行了操作。  

# 结果持久化
对数据源输出（列举）结果进行持久化操作（目前支持写入到本地文件），持久化选项：  
```
save-total=
save-path=
save-format=
save-separator=
rm-fields=
```  
|参数名|参数值及类型 | 含义|  
|-----|-------|-----|  
|save-total| true/false| 是否直接保存数据源完整输出结果，针对存在下一步处理过程时是否需要保存原始数据|  
|save-path| local file 相对路径字符串| 表示保存结果的文件路径|  
|save-format| json/tab/csv| 结果保存格式，将每一条结果记录格式化为对应格式，默认为 tab 格式（减小输出结果的体积）|  
|save-separator| 字符串| 结果保存为 tab 格式时使用的分隔符，结合 save-format=tab 默认为使用 "\t"|  
|rm-fields| 字符串列表| 保存结果中去除的字段，为输入行中的实际字段选项，用 "," 做分隔，如 key,hash，表明从结果中去除 key 和 hash 字段再进行保存，不填表示所有字段均保留|  

**关于save-total**  
（1）用于选择是否直接保存数据源完整输出结果，针对存在过滤条件或下一步处理过程时是否需要保存原始数据，如 bucket 的 list 操作需要在列举出结果之后再针对字段进行过滤或者做删除，save-total=true 则表示保存列举出来的完整数据，而过滤的结果会单独保存，如果只需要过滤之后的数据，则设置为 false，如果是删除等操作，通常删除结果会直接保存文件名和删除结果，原始数据也不需要保存。
（1）本地文件数据源时默认如果存在 process 或者 filter 则设置 save-total=false，反之则设置 save-total=true（说明可能是单纯格式转换）。  
（2）云存储数据源时默认设置 save-total=true。  
（3）保存结果的路径 **默认（save-path）使用 <bucket\>（云存储数据源情况下）名称或者 <path\>-result 来创建目录**。  

**关于持久化文件名** 
（1）持数据源久化结果的文件名为 "<source-name\>\_success_<order\>.txt"，如 qiniu 存储数据源结果为 "qiniu_success_<order\>.txt"，local 数据源结果为 "local_success_<order\>.txt"。  
（2）如果设置了过滤选项或者处理过程，则过滤到的结果文件名为 "filter_success/error_<order\>.txt"。
（3）process 过程保存的结果为文件为 "<process\>\_success/error\_<order\>.txt"，<process\>\_success/error\_<order\>.txt 表明无法成功处理的结果，<process\>\_need_retry\_<order\>.txt，表明为需要重试的记录，可能需要确认所有错误数据和记录的错误信息。  

**关于 rm-fields** 
rm-fields 可选择持久化结果中去除某些字段，未设置的情况下保留所有原始字段，数据源导出的每一行信息以目标格式 save-format 保存在 save-path 的文件中。file 数据源输入字段完全取决于 indexes 和其他的一些 index 设置，可参考 [indexes 索引](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#%E5%85%B3%E4%BA%8E-indexes-%E7%B4%A2%E5%BC%95)，而其他 index 设置与数据处理类型有关，比如 url-index 来输入 url 信息。对于云储存数据源，不使用 indexes 规定输入字段的话默认是保留所有字段，字段定义可参考[关于文件信息字段](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#%E5%85%B3%E4%BA%8E%E6%96%87%E4%BB%B6%E4%BF%A1%E6%81%AF%E5%AD%97%E6%AE%B5)    

结果持久化详细配置说明见 [持久化配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/resultsave.md)。  

# 超时设置
多数数据源或者操作涉及网络请求，因此提供超时时间设置，默认的超时时间一般能够满足要求，特殊需要的情况下可以修改各超时时间：  
`connect-timeout=60` 网络连接超时时间，程序默认 60s  
`read-timeout=120` socket 读取超时时间，程序默认 120s  
`request-timeout=60` 网络请求超时时间，程序默认 60s  

# 错误及异常
1、一般情况下，终端输出异常信息如 socket timeout 超时为正常现象，如：
```
list prefix:<prefix> retrying...
...
java.net.SocketTimeoutException: timeout
```
程序会自动重试，如果比较频繁则可以修改[超时配置](#超时设置)重新运行程序，超过重试次数或者其他非预期异常发生时程序会退出，可以将异常信息反馈在[ ISSUE 列表](https://github.com/NigelWu95/qiniu-suits-java/issues) 中。  

2、常见错误信息：  
（1）java.lang.UnsupportedClassVersionError: Unsupported major.minor version ...  
请使用 java 8 或以上版本的 jdk（jre） 环境来运行该程序。  
（2）java.lang.OutOfMemoryError: GC overhead limit exceeded  
表示可能是内存中加载了过多的资源导致 java 的 gc 内存溢出，需要关闭程序重新运行，更换高配置机器或者降低线程数 threads 或者 unit-len 重新运行，如果已运行较长时间可采用[断点续操作](#断点续操作)。  
（3）java.lang.OutOfMemoryError: unable to create new native thread   
与（1）类似，内存溢出导致无法继续创建更多线程或对象，降低线程数 threads 重新运行。  
（4）java.lang.OutOfMemoryError: Java heap space   
运行过程中 jvm 堆内存（一般默认为系统内存的 1/4）不足，降低线程数 threads 或者 unit-len 重新运行，如果已运行较长时间可采用[断点续操作](#断点续操作)。  

# 程序日志
qsuits-java 使用 slf4j+log4j2 来记录运行日志，日志产生在当前路径的 logs 目录下，说明如下：  
1、数据源位置记录信息 =\> procedure.log，记录行格式为 json，数据源读取位置打点数据，每一行都是一次数据源位置记录，最后一行即为最后记录下的位置信息，如果信息为 `{}` 表明程序运行完整，没有断点需要再次运行，如果信息中包含具体的字符串，说明这是程序留下的断点，则该行信息可以取出作为断点操作的配置内容，具体参考：[断点操作](#断点续操作)  
2、程序运行过程输出及异常信息，通过终端 Console 和 qsuits.info、qsuits.error 输出。  
3、日志输出的默认文件名为 procedure.log、qsuits.info 和 qsuits.error，每次运行前会检查当前路径下是否存在历史日志文件，如果存在则会将文件名加上数字，如 procedure0.log、qsuits0.info、qsuits0.error 或 procedure1.log、qsuits1.info、qsuits1.error。  

# 断点续操作
支持断点记录，在程序运行后出现异常导致终止或部分数据源路径错误或者是 INT 信号(命令行 Ctrl+C 中断执行)终止程序时，会记录数据导出中断的位置，记录的信息可用于下次直接从未完成处继续导出数据，而不需要全部重新开始。尤其在对云存储空间列举文件列表时，特大量文件列表导出耗时可能会比较长，可能存在断点续操作的需求，续操作说明：  
1、如果存在续操作的需要，终止程序时会在 save-path 同路径下生成一个 json 文件，存储数据源生成的默认文件名为 <bucket\>-prefixes.json，file 数据源生成的文件名为 <path\>-lines.json，该 json 文件中记录了断点信息。如：  
```
➜ ~ cat ../temp-prefixes.json 
{"0":null,"2":null,"3":null,"7":null,"C":null,"D":null,"E":null,"F":null,"I":null,"O":null,"U":null,"W":null,"a":null,"b":null,"c":null,"f":null,"g":null,"k":null,"l":null,"m":null,"q":null,"r":null,"t":null}
```
2、对于云存储文件列表列举操作记录的断点可以直接作为下次续操作的操作来使用完成后续列举，如断点文件为 <filename\>.json，则在下次列举时使用断点文件作为前缀配置文件: prefix-config=<breakpoint_filepath\> 即可，参见：[prefix-config 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#prefix-config-配置)。  
3、对于 file 数据源产生的断点文件记录了读取的文本行，亦可以直接作为下次续操作的操作来使用完成后续列举，如断点文件为 <filename\>.json，则在下次继续读 file 数据源操作时使用断点文件作为行配置文件: line-config=<breakpoint_filepath\> 即可，参见：[line-config 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#line-config-配置)。  
4、断点续操作时建议修改下 save-path，便于和上一次保存的结果做区分，否则可能会覆盖上次的结果，且文件名难以区分。  

**注意：如果是系统宕机、断电或者强制关机或者进程强行 kill 等情况，无法得到输出的断点提示，因此只能通过[<位置记录日志>](#程序日志)来查看最后的断点信息，取出 procedure.log 日志的最后一行并创建断点配置文件，从而按照上述方式进行断点运行。**  

# 分布式任务方案
对于不同账号或空间可以直接在不同的机器上执行任务，对于单个空间资源数量太大无法在合适条件下使用单台机器完成作业时，可分机器进行作业，如对一个空间列举完整文件列表时，可以按照连续的前缀字符分割成多段分别执行各个机器的任务，建议的前缀列表为:  
```!"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz```，将该列表任意分成 n 段，如：
```
prefixes=!,",#,$,%,&,',(,),*,+,\,,-,.,/,0,1
prefixes=2,3,4,5,6,7,8,9,:,;
prefixes=<,=,>,?,@,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O
prefixes=P,Q,R,S,T,U,V,W,X,Y,Z,[,\\,],^,_,`
prefixes=a,b,c,d,e,f,g,h,i,j,k,l,m
prefixes=n,o,p,q,r,s,t,u,v,w,x,y,z
```
（**`,`，`\` 需要转义**）将前缀分为上述几段后，设置 prefixes 参数可以分做六台机器执行，同时因为需要列举空间全部文件，需要分别在第一段 prefixes 设置 `prefix-left=true`，在最后一段 prefixes 设置 `prefix-right=true`（其他段 prefixes 不能同时设置 prefix-left 或 prefix-right，且仅能第一段设置 prefix-left 和最后一段设置 prefix-right，参数描述见[数据源完备性](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#数据源完备性和多前缀列举)  

# 时间计划
支持时间计划参数，控制任务按照时间调度来执行，包含延迟执行和周期性（周期为 1 天）暂停策略，参数如下（不设置则无时间计划）：  
```
start-time=2019-10-07 08:00:00
pause-delay=36000
pause-duration=50400
```  
|参数           |含义                    |  
|--------------|-----------------------|  
|start-time    |任务正式开始执行的时间点，时间格式为：yyyy-MM-DD HH:MM:SS，HH:MM:SS 为 00:00:00（缺省值）时可省略|  
|pause-delay   |任务开始执行后经过多少秒（默认 0s）进行第一次暂停，单位 s(秒)，之后每天的该时间点执行暂停|  
|pause-duration|任务每天的暂停持续时间，单位 s(秒)，每次暂停经过该段时间后恢复任务的执行|  

比如上述参数表示在 2019-10-07 08:00:00 开始运行，运行 10 小时后在 2019-10-07 18:00:00 对任务进行暂停，暂停 14 个小时，第二天的 08:00:00 恢复运行，因此该配置的含义表示任务在每天的 08:00:00-18:00:00 期间运行（18:00:00-第二天08:00:00 期间暂停）。start-time 不允许超出到一周之后，pause-duration 的最小暂停时间为 1800s(0.5 小时) 最大暂停时间为 84600s(23.5 小时)。pause-delay 的默认值为 0，小于 0 时表示不执行时间计划，或者 pause-duration 小于 0 时同样表示不执行时间计划。  

# 暂停和恢复
暂停和恢复是操作系统特性，如果对系统熟悉也可以基于此来制作时间计划。Linux/Mac 下支持以下操作来暂停和恢复进程：  
暂停（Ctrl + Z 命令）：  
![](http://qsuits.nigel.net.cn/qsuits_ctrlz_to_pause.jpg)
恢复（fg 命令，注意该命令需要在暂停时的同个 terminal 下执行，同时建议不要做路径切换）：  
![](http://qsuits.nigel.net.cn/qsuits_fg_to_continue.jpg)

# 问题反馈

如果您有任何问题，可以在 [ISSUE 列表](https://github.com/NigelWu95/qiniu-suits-java/issues)里反馈，我们会尽快回复您，感谢您的支持与理解。