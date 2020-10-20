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


$ ./mycli --url http://xxcatchers.com/links
Url : http://xxcatchers.com/links
StatusCode : 200
Duratime : 978446557ns
ResponseSize : 156
Response : {"HTML":"https://gothic-province-290512.uc.r.appspot.com","interview.io":"https://interviewing.io/","coivd tracker":"https://coronavirus.1point3acres.com/"}
```

## run --profile
```shell script


$ ./mycli --url http://xxcatchers.com/links -profile 20
Url : http://xxcatchers.com/links
================
TotalRequestNum : 20
FastestTime : 186749563ns
SlowestTime : 398366555ns
MeanTime : 208114797.000000ns
MedianTime : 196916669.000000ns
SuccessRequest : 20
ErrorRequest : 0
Percentage of requests : 1.000000
ErrorCode : []
ErrorUrl : []
SmallestRes(bytes) : 156
LargestRes(bytes) : 156
================
Execute Duration : 4162382838ns



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
