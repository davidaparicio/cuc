## Examples

With Go (Build or Run)

```
david@mac:~$ CGO_ENABLED=1 go run ./main.go loop -v -s 1 -c 200 -f assets/mp3/ubuntu_dialog_info.mp3
{"level":"info","ts":1678738578.205648,"caller":"cmd/loop.go:18","msg":"It's a match!","attempt":1,"statuscode":200,"backoff":1,"url":"https://www.example.com/"}
{"level":"info","ts":1678738578.855443,"caller":"cmd/loop.go:18","msg":"It's a match!","attempt":2,"statuscode":200,"backoff":1,"url":"https://www.example.com/"}
{"level":"info","ts":1678738579.8583958,"caller":"cmd/loop.go:18","msg":"It's a match!","attempt":3,"statuscode":200,"backoff":1,"url":"https://www.example.com/"}
^C{"level":"info","ts":1678738580.211995,"caller":"runtime/asm_amd64.s:1598","msg":"Caught the following signal","signal":"interrupt"}
{"level":"info","ts":1678738580.212107,"caller":"internal/core.go:118","msg":"Graceful shutdown..","ctx.err":"context canceled"}
```