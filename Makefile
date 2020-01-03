build:
	GOOS="darwin" GOARCH="amd64" go build -o "bin/qsuits_darwin_amd64" src/entry.go
# 	GOOS="darwin" GOARCH="386" go build -o "bin/qsuits_darwin_386" src/entry.go
	GOOS="windows" GOARCH="amd64" go build -o "bin/qsuits_windows_amd64.exe" src/entry.go
	GOOS="windows" GOARCH="386" go build -o "bin/qsuits_windows_386.exe" src/entry.go
	GOOS="linux" GOARCH="amd64" go build -o "bin/qsuits_linux_amd64" src/entry.go
	GOOS="linux" GOARCH="386" go build -o "bin/qsuits_linux_386" src/entry.go

release_tswork:
	rm -rf logs*
	qsuits -path=bin -process=qupload -a=tswork -bucket=qsuits -keep-path=false -save-path=logs
	qsuits -path=logs/qupload_success_1.txt -rm-keyPrefix=bin/ -process=cdnrefresh -a=tswork -domain=qsuits.nigel.net.cn -protocol=http -save-path=logs1
	qsuits -path=logs/qupload_success_1.txt -rm-keyPrefix=bin/ -process=cdnrefresh -a=tswork -domain=qsuits.nigel.net.cn -protocol=https -save-path=logs2
	qsuits -path=logs/qupload_success_1.txt -rm-keyPrefix=bin/ -process=cdnprefetch -a=tswork -domain=qsuits.nigel.net.cn -protocol=http -save-path=logs3
	qsuits -path=logs/qupload_success_1.txt -rm-keyPrefix=bin/ -process=cdnprefetch -a=tswork -domain=qsuits.nigel.net.cn -protocol=https -save-path=logs4

release_devtools:
	rm -rf logs*
	qsuits -path=bin -process=qupload -a=devtools -bucket=devtools -keep-path=false -save-path=logs
	qsuits -path=logs/qupload_success_1.txt -rm-keyPrefix=bin/ -process=cdnrefresh -a=devtools -domain=devtools.qiniu.com -protocol=http -save-path=logs1
	qsuits -path=logs/qupload_success_1.txt -rm-keyPrefix=bin/ -process=cdnrefresh -a=devtools -domain=devtools.qiniu.com -protocol=https -save-path=logs2
	qsuits -path=logs/qupload_success_1.txt -rm-keyPrefix=bin/ -process=cdnprefetch -a=devtools -domain=devtools.qiniu.com -protocol=http -save-path=logs3
	qsuits -path=logs/qupload_success_1.txt -rm-keyPrefix=bin/ -process=cdnprefetch -a=devtools -domain=devtools.qiniu.com -protocol=https -save-path=logs4
