# myclient

## build
run follow command
```shell script
$ make build
go build -o mycli -v

```
you will get `mycli` binary on current platform

## build all platform
run follow command
```shell script
$ make buildall
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mycli_unix -v
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o mycli_mac -v
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mycli_win -v

$ ll
total 54056
-rw-r--r--  1 nagatyase  staff   742B Oct 18 19:32 Makefile
-rw-r--r--  1 nagatyase  staff   1.2K Oct 18 19:33 README.md
-rwxr-xr-x  1 nagatyase  staff   250B Oct 18 19:19 dev.sh
-rw-r--r--  1 nagatyase  staff   4.8K Oct 18 19:18 main.go
-rwxr-xr-x  1 nagatyase  staff   6.6M Oct 18 19:33 mycli
-rwxr-xr-x  1 nagatyase  staff   6.7M Oct 18 19:33 mycli_mac
-rwxr-xr-x  1 nagatyase  staff   6.6M Oct 18 19:33 mycli_unix
-rwxr-xr-x  1 nagatyase  staff   6.4M Oct 18 19:33 mycli_win

```
you will get `mycli` binary

## run --url
```shell script
$ ./mycli --url https://api.muxiaoguo.cn/api/dujitang             
Url : https://api.muxiaoguo.cn/api/dujitang
StatusCode : 200
Duratime : 453535972ns
ResponseSize : 127
Response : {"code":"200","msg":"success","data":{"comment":"你现在的生活，也许不是你想要的，但绝对是你自找的。"}}

```

## run --profile
```shell script
$ ./mycli --url https://api.muxiaoguo.cn/api/dujitang --profile 20
Url : https://api.muxiaoguo.cn/api/dujitang
================
TotalRequestNum : 20
FastestTime : 58968046ns
SlowestTime : 298615547ns
MeanTime : 77617633.000000ns
MedianTime : 61352383.000000ns
SuccessRequest : 20
ErrorRequest : 0
Percentage of requests : 100%
ErrorCode : []
ErrorUrl : []
SmallestRes(bytes) : 85
LargestRes(bytes) : 175
================
Execute Duration : 1552445442ns



```

## run --help
```shell script
$ ./mycli --help
Usage of ./mycli:
  -profile int
        number of requests(a positive integer)
  -url string
        any web sites url

```