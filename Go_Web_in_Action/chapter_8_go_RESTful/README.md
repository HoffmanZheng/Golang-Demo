# Go RESTful API Interface Development

1. In the team development, repetitive functions could be anywhere if there is no unified standard. Because there is no unified calling specification, it's alse difficult for us to understand the code written by other people, and it's hard to start secondary development and maintenance. 

2. A mature framework helps developers realize many basic functions. Developers only need to concentrate on implementing the required business logic.

### Usage of Beego framework

1. Beego is a high-performance HTTP framework developed by Golang, which is commonly used to develop API, web application, backend service, etc. 

2. A Beego project could easily be initiaized by using bee command-line tool, as below:
```powershell
go get github.com/beego/beego
go get github.com/beego/bee
bee new <projectName>
cd ./<projectName>  && bee run
```

3. [conf/app.conf](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_8_go_RESTful/beego-project/conf/app.conf) is the default configuration file of project, which could `include` other conf files. These config could be retrieved by using `beego.AppConfig.String("mysql_host")`.

### Usage of Gin framework

1. The router of Gin guides the request to its corresponding handler, see: [1.ginHello.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_8_go_RESTful/1.ginHello.go) On the other hand, different format of request and response could be handled, see: [2.ginParseRequest.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_8_go_RESTful/2.ginParseRequest.go)

2. Gin also provides middleware to intercept HTTP　requests and perform some actions like: permission check, error handle, etc. See: [3.ginMiddleware.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_8_go_RESTful/3.ginMiddleware.go). Global middleware and middleware for routing grouping are both supported in Gin.

3. A RESTful API developed by Gin is displayed in [4.ginRESTFul.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_8_go_RESTful/4.ginRESTFul.go).