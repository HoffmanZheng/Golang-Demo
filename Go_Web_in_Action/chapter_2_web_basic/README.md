# Go Web Development Fundamentals

### Introduction to Web Application Principle

1. It's quite easy to develop a web application in Golang, with the help of `HandleFunc()` and `ListenAndServe()` function in net/http package, See: [1.helloWorldWeb.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_2_web_basic/1.helloWorldWeb.go)

2. Web fundamental, See: [Web：浅析 URL](https://hoffmanzheng.github.io/2020/web-url/), [深入 Web 请求过程](https://hoffmanzheng.github.io/2020/dns-cdn/), [Net：计算机网络读书笔记](https://hoffmanzheng.github.io/2020/net-http-tcp/)

3. MVC is a common framework in software engineering:

   ![](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/images/mvc_role_diagram.png)

4. Template engine enable the separation of user interface and business data, which could generate documents in specific format.

### net/http Package in Golang

1. `http.ListenAndServe()` has two parameters, listenning port and event handler. `Handler` is an interface, which is triggered before `http.HandleFunc()` like an interceptor. A custom handler could written by implementing the interface, which only returns the request whose header carries specific Referer, see: [2.myWebHandler.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_2_web_basic/2.myWebHandler.go)
2. HTTPS could be easily implemented using `ListenAndServeTLS`, see: [3.httpsServer.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_2_web_basic/3.httpsServer.go)
3. Client struct is also provided in net/http package, which could mock http client and send http request, see: [4.getRequest.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_2_web_basic/4.getRequest.go), [5.postRequest.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_2_web_basic/5.postRequest.go), [6.putRequest.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_2_web_basic/6.putRequest.go), [7.deleteRequest.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_2_web_basic/7.deleteRequest.go)

### html/template Package in Golang

1. Template engine is commonly used in MVC framework to render dynamic web pages, which extract the unchanging part, the variable part is provided by the back-end program.
2. Golang built-in html/template package uses `{{data.FieldName}}` to get the incoming data. The template object could be parsed and built by `ParseFiles()`, the rendering part is fulfilled by `Execute()`, see: [8.htmlTemplate.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_2_web_basic/8.htmlTemplate.go)
3. Template deal with  iterable data type, see: [9.template-range.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_2_web_basic/9.template-range.go); template with custom function, see: [10.template-funcs.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_2_web_basic/10.template-funcs.go); template with nested reference, see: [11.template-multi.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_2_web_basic/11.template-multi.go)

