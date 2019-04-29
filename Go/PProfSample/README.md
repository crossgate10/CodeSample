#### Refeference: 
[https://blog.yeeef.com/post/golang-concurrency/ (only web cache)](https://webcache.googleusercontent.com/search?q=cache:UnFGnwYNbq4J:https://blog.yeeef.com/post/golang-concurrency/+&cd=1&hl=zh-TW&ct=clnk&gl=tw&lr=lang_zh-CN%7Clang_zh-TW)

[Using the Go execution tracer to speed up fractal rendering](https://medium.com/justforfunc/using-the-go-execution-tracer-to-speed-up-fractal-rendering-c06bb3760507)

[[Go] pprof 效能分析](https://zwindr.blogspot.com/2018/10/go-go-pprof.html)

[使用 pprof 和火焰图调试 golang 应用](https://cizixs.com/2017/09/11/profiling-golang-program/)



##### ＠Commands
<pre>
go run main.go
</pre>
In folder, you'll see **cpu.prof**, **mem.mprof** & **trc.trace**

then you could run command below to get into interaction mode of pprof:
<pre>
// To profile cpu
<b>go tool pprof cpu.prof</b>

// To profile mem
<b>go tool pprof mem.mprof</b>

// To trace in browser
// Prerequisite: In mac, you should run <b>brew install GraphViz</b>
<b>go tool trace trc.trace</b>
</pre>

##### @Cases
1. Directly, without goroutine
2. With goroutine each pixel
3. With goroutine each row
4. With worker pattern
5. With buffered channel
6. With worker each row
