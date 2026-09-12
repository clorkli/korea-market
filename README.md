# korea-market

开发中的 Go CLI，通过 Bithumb 公共 API 获取 KRW 加密货币行情快照。

目前支持 BTC、ETH、XRP，输出价格和涨跌幅，支持请求超时与部分成功结果。

## Run

需要 Go 1.27.0+，运行时需联网。

```bash
go run .
```

## Test

```bash
go test ./...
```
