# Handle Go Web requests

### Receive Web Requests

1. Serve Multiplexer receives web requests and forwards them to the corresponding handlers. Database interactions would take place if necessary. The response would be rendered by template engine, the whole process is display as below:

   ![ServerMux](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/images/ServerMux.jpeg)

2. `ServeMux` is a struct contains a mapping, which map URL to corresponding handler. Requests will be redirected according to that map. `http.HandleFunc()` would register a handler for a certain URL. Multiple handlers could be specified in that way, see: [1.useDefaultServeMux.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_3_handle_request/1.useDefaultServeMux.go)

3. Although it's convenient to use `DefaultServeMux`, but it's not recommend to use it in production. Because defaultServeMux is a global variable. It might be used or modified by some third-party components, which leads to conflicts. The recommended approach is to use custom ServerMux, see: [2.customisedServeMux.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_3_handle_request/2.customisedServeMux.go)

4. One restriction of ServeMux is: it can not match the URL pattern. `HttpRouter` could make up for that shortcommings, see: [3.httpRouter.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_3_handle_request/3.httpRouter.go)

5. Custom handler is more flexible and powerful, but it's a little tedious to define a new struct for each handler, see: [1.useDefaultServeMux.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_3_handle_request/1.useDefaultServeMux.go). HandlerFunc could be defined and registered individualy by `HandleFunc()`, see: [2.customisedServeMux.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go_Web_in_Action/chapter_3_handle_request/2.customisedServeMux.go). 

### Handle Web Requests