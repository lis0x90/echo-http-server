# echo-http-server
Simple echo http server (with slow response ability)

# Run
Run http server on `localhost:9999` and emulate slow response with timeout `2s`:

```
echo_http_server -bind localhost:9999 -sleep-time 2s
```

Test server with `curl`:
```
$ curl -w "\n\nTiming: %{time_connect}:%{time_starttransfer}:%{time_total}" localhost:9999/abc/cde
Path: abc/cde

Timing: 0,203:2,203:2,203
```
