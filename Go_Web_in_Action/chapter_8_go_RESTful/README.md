# Go RESTful API Interface Development

1. In the team development, repetitive functions could be anywhere if there is no unified standard. Because there is no unified calling specification, it's alse difficult for us to understand the code written by other people, and it's hard to start secondary development and maintenance. 

2. A mature framework helps developers realize many basic functions. Developers only need to concentrate on implementing the required business logic.

### Usage of Beego framework



### Usage of Gin framework

1. The router of Gin guides the request to its corresponding handler, see: [1.ginHello.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_8_go_RESTful/1.ginHello.go) On the other hand, different format of request and response could be handled, see: [2.ginParseRequest.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_8_go_RESTful/2.ginParseRequest.go)

2. Gin also provides middleware to intercept HTTP　requests and perform some actions like: permission check, error handle, etc. See: [3.ginMiddleware.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_8_go_RESTful/3.ginMiddleware.go)

### OAuth 2.0 developed by Golang