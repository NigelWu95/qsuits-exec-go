package manual

import "fmt"

func AccountUsage()  {

	fmt.Println("1. 设置 account：")
	fmt.Println("命令格式：-account=<source>-<name> -<source>-id= -<source>-secret= [-d]，如：")
	fmt.Println()
	fmt.Println("-account=test/qiniu-test -ak= -sk= 设置七牛账号，账号名为 test，没有数据源标识时默认设置七牛账号")
	fmt.Println("-account=ten-test -ten-id= -ten-secret= 设置腾讯云账号，账号名为 test")
	fmt.Println("-account=ali-test -ali-id= -ali-secret= 设置阿里云账号，账号名为 test")
	fmt.Println("-account=s3-test -s3-id= -s3-secret= 设置 AWS/S3 账号，账号名为 test")
	fmt.Println("-account=up-test -up-id= -up-secret= 设置又拍云账号，账号名为 test")
	fmt.Println("-account=hua-test -hua-id= -hua-secret= 设置华为云账号，账号名为 test")
	fmt.Println("-account=bai-test -bai-id= -bai-secret= 设置百度云账号，账号名为 test")
	fmt.Println("-d 表示默认账号选项，此时设置的账号将会成为全局默认账号，执行操作时 -d 选项将调取该默认账号")
	fmt.Println()
	fmt.Println("2. 使用 account 账号：")
	fmt.Println("-a=test 表示使用 test 账号，数据源会自动根据 path 参数判断")
	fmt.Println("-d 表示使用默认的账号，数据源会自动根据 path 参数判断")
	fmt.Println()
	fmt.Println("3. 查询 account 账号：")
	fmt.Println("命令格式：-getaccount=<name>-<source> [-dis] [-d]，默认只显示 id 的明文而隐藏 secret，-dis 参数表示选择明文显示 secret，如：")
	fmt.Println()
	fmt.Println("-getaccount -d 表示查询设置的默认账号的密钥")
	fmt.Println("-getaccount=test -dis 表示查询设置的所有账号名为 test 的密钥，并显示 secret 的明文")
	fmt.Println("-getaccount=test-s3 表示查询设置的 S3 账号名为 test 的密钥")
	fmt.Println("-getaccount=test-qiniu 表示查询设置的七牛账号名为 test 的密钥")
	fmt.Println("-getaccount=test-tencent 表示查询设置的腾讯账号名为 test 的密钥")
}
