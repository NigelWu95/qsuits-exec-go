package manual

import "fmt"

func FileUsage() {

	fmt.Println("path=")
	fmt.Println("unit-len=")
	fmt.Println("threads=")
	fmt.Println("indexes=key,etag,size")
	fmt.Println()
	fmt.Println("参数值     含义")
	fmt.Println("path      选择从[云存储空间列举]还是从[本地路径中读取]资源，本地数据源时填写本地文件或者目录路径，云存储数据源时可填写 \"qiniu://<bucket>\"、\"tencent://<bucket>\" 等")
	fmt.Println("unit-len  表示一次读取的文件个数（读取或列举长度，不同数据源有不同默认值），对应与读取文件时每次处理的行数或者列举请求时设置的 limit 参数")
	fmt.Println("threads   表示预期最大线程数，若实际得到的文件数或列举前缀数小于该值时以实际数目为准")
	fmt.Println("indexes   资源元信息字段索引（下标），设置输入行对应的元信息字段下标")

	fmt.Println()
	fmt.Println("parse=tab/json/file")
	fmt.Println("separator=\t")
	fmt.Println("# 文件内容读取资源列表时一般可能需要设置 indexes 参数（默认只包含 key 字段的解析）")
	fmt.Println("indexes=")
	fmt.Println("add-keyPrefix=")
	fmt.Println("rm-keyPrefix=")
	fmt.Println("......")
	fmt.Println("参数名 参数值及类型  含义")
	fmt.Println("parse           json/tab/csv，数据行格式，json 表示使用 json 解析，tab 表示使用分隔符（默认 \t）分割解析，csv 表示使用 \",\" 分割解析")
	fmt.Println("separator       当 parse=tab 时，可另行指定该参数为格式分隔符来分析字段")
	fmt.Println("add-keyPrefix   将解析出的 key 字段加上指定前缀再进行后续操作，用于输入 key 可能比实际空间的 key 少了前缀的情况，补上前缀才能获取到资源")
	fmt.Println("rm-keyPrefix    将解析出的 key 字段去除指定前缀再进行后续操作，用于输入 key 可能比实际空间的 key 多了前缀的情况，如输入行中的文件名多了 / 前缀")

	fmt.Println()
	fmt.Println("数据源详细参数配置和说明及可能涉及的高级用法见：数据源配置 https://github.com/NigelWu95/qiniu-suits-java/blob/master/docs/datasource.md#2-file-本地文件读取")
}
