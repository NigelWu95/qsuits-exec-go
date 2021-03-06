package manual

import "fmt"

func ProcessUsage() {

	fmt.Println()
	fmt.Println("process 处理过程表示对数据源输入的每一条记录进行处理的类型，所有处理结果保存在 save-path 或默认路径下，具体处理过程由处理类型参数指定，如 **process=type/status/lifecycle/copy** (命令行方式则指定为 **-process=xxx**等，同时 process 操作支持设置公共参数：")
	fmt.Println("retry-times= 操作失败（可重试的异常情况下，如请求超时）需要进行的重试次数，默认为 5 次 ")
	fmt.Println("batch-size= 支持 batch 操作时设置的一次批量操作的文件个数（支持 batch 操作：type/status/lifecycle/delete/copy/move/rename/stat/cdnrefresh/cdnprefetch，其他操作请勿设置 batchSize 或者设置为 0），当响应结果较多 429/573 等状态码时可能是超过并发限制需要降低 batch-size，或者直接使用非 batch 方式：batch-size=0/1")
	fmt.Println()
	fmt.Println("**处理操作类型：**")
	fmt.Println()
	fmt.Println("|处理类型配置           |含义                                         |文档链接                                                                             |")
	fmt.Println("|-------------------- |--------------------------------------------|-----------------------------------------------------------------------------------|")
	fmt.Println("|process=qupload      | 表示上传文件到存储空间                         | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/uploadfile.md       |")
	fmt.Println("|process=syncupload   | 表示将 url 的内容直传到存储空间                 | [syncupload 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/syncupload.md)|")
	fmt.Println("|process=delete       | 表示删除空间资源                              | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/delete.md           |")
	fmt.Println("|process=copy         | 表示复制资源到指定空间                         | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/copy.md             |")
	fmt.Println("|process=move         | 表示移动资源到指定空间                         | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/move.md             |")
	fmt.Println("|process=rename       | 表示对指定空间的资源进行重命名                   | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/rename.md           |")
	fmt.Println("|process=stat         | 表示查询空间资源的元信息                        | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/stat.md             |")
	fmt.Println("|process=type         | 表示修改空间资源的存储类型（低频/标准）            | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/type.md             |")
	fmt.Println("|process=status       | 表示修改空间资源的状态（启用/禁用）               | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/status.md           |")
	fmt.Println("|process=lifecycle    | 表示修改空间资源的生命周期                      | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/lifecycle.md        |")
	fmt.Println("|process=mirror       | 表示对设置了镜像源的空间资源进行镜像更新           | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/mirror.md           |")
	fmt.Println("|process=asyncfetch   | 表示异步抓取资源到指定空间                      | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/asyncfetch.md       |")
	fmt.Println("|process=fetch        | 表示同步抓取 url 资源到指定空间                 | [fetch 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/fetch.md)|")
	fmt.Println("|process=qhash        | 表示查询资源的 qhash                          | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/qhash.md            |")
	fmt.Println("|process=avinfo       | 表示查询空间资源的视频元信息                     | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/avinfo.md           |")
	fmt.Println("|process=pfopcmd      | 表示根据音视频资源的 avinfo 信息来生成转码指令     | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/pfopcmd.md          |")
	fmt.Println("|process=pfop         | 表示对空间资源执行 pfop 请求                    | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/pfop.md             |")
	fmt.Println("|process=pfopresult   | 表示通过 persistentId 查询 pfop 的结果         | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/pfopresult.md       |")
	fmt.Println("|process=privateurl   | 表示对私有空间资源进行私有签名                   | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/privateurl.md       |")
	fmt.Println("|process=exportts     | 表示对 m3u8 的资源进行读取导出其中的 ts 文件列表   | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/exportts.md         |")
	fmt.Println("|process=download     | 表示通过 http 下载资源到本地                    | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/downloadfile.md     |")
	fmt.Println("|process=imagecensor  | 表示图片类型资源内容审核                        | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/censor.md#图片审核)   |")
	fmt.Println("|process=videocensor  | 表示视频类型资源内容审核                        | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/censor.md#视频审核)   |")
	fmt.Println("|process=censorresult | 表示内容审核结果查询                           | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/censorresult.md     |")
	fmt.Println("|process=mime         | 修改资源的 mimeType                          | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/mime.md             |")
	fmt.Println("|process=metadata     | 修改资源的 metadata                          | https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/metadata.md         |")
	fmt.Println("|process=cdnrefresh   | 刷新 CDN 刷新                               | [cdnrefresh 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/cdn.md#cdn-刷新缓存)|")
	fmt.Println("|process=cdnprefetch  | CDN 资源预取（预热）                          | [cdnprefetch 配置](https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/cdn.md#cdn-资源预取)|")
	fmt.Println()
	fmt.Println("**注意**：")
	fmt.Println("1、云存储数据源 + process 操作的情况下通常会涉及两对密钥，数据源一对，process 操作一对，如果是 delete、status 等操作则这两对密钥相同，使用一个密钥设置或者一个 account (`-a=<account-name>`) 即可，copy、move 要求针对同一个账号操作或者采用空间授权，因此也只需要一堆密钥，但如果是其他存储数据源的数据备份操作 asyncfetch，就需要两对不同的密钥，而 account 只支持设置一个，这时第二对的七牛密钥可以通过同一个 account-name 的设置来获得，因为同一个 account-name 可以为不同数据源做密钥设置，如：`-account=ali-test -ali-id= -ali-secret=` 设置了阿里云 test 名称的账号，同时 `-account=qiniu-test -ak= -sk=` 设置了七牛 test 名称的账号，则通过 `-a=test` 可以同时拿到阿里云和七牛云的 test 账号，因此可以直接通过同一个 account-name 来进行操作。但是如果明确指定了另外的 ak，sk，则会使用您设置的这一对七牛密钥。")
	fmt.Println("2、也真是因为不同数据源的 account-name 可同名特性，以及支持主动设置密钥来覆盖 account 的密钥，在具体操作时需要注意账号和密钥的使用，以免对另外一个账号执行了操作。")
	fmt.Println()
	fmt.Println("process 操作类型较多，且参数各异，还请查看在线文档来帮助使用，如果有必要参数缺少时程序也会出现提示。")
}
