# Go Web Development Fundamentals

### Introduction to Web Application Principle

1. It's quite easy to develop a web application in Golang, with the help of `HandleFunc()` and `ListenAndServe()` function in net/http package, See: [1.helloWorldWeb.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go%20Web%20in%20Action/chapter_2_web_basic/1.helloWorldWeb.go)

2. Web fundamental, See: [Web：浅析 URL](https://hoffmanzheng.github.io/2020/web-url/), [深入 Web 请求过程](https://hoffmanzheng.github.io/2020/dns-cdn/), [Net：计算机网络读书笔记](https://hoffmanzheng.github.io/2020/net-http-tcp/)

3. MVC is a common framework in software engineering:

   [](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go%20Web%20in%20Action/images/mvc_role_diagram.png)

4. Template engine enable the separation of user interface and business data, which could generate documents in specific format.

### net/http Package in Golang

1. `http.ListenAndServe()` has two parameters, listenning port and event handler. `Handler` is an interface, which is triggered before `http.HandleFunc()` like an interceptor. A custom handler could written by implementing the interface, which only returns the request whose header with specific Referer, see: [2.myWebHandler.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go%20Web%20in%20Action/chapter_2_web_basic/2.myWebHandler.go)
2. HTTPS could be easily implemented using `ListenAndServeTLS`, See: [3.httpsServer.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go%20Web%20in%20Action/chapter_2_web_basic/httpsServer.go)
3. 

### html/template Package in Golang

