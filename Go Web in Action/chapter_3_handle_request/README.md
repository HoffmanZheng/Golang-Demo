# Handle Go Web requests

### Receive Web Requests

1. ServeMux receives web requests and forwards it to the corresponding controller. Database interactions would take place if necessary. The response would be rendered by template engine, the whole process is display as below:

   ![ServerMux](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go%20Web%20in%20Action/images/ServerMux.jpeg)

2. `ServeMux` is a struct contains a mapping, which map URL to corresponding handler. Requests will be redirected according to that map. `http.HandleFunc()` would register a handler for a certain URL. Multiple handlers could be specified in that way, see: [1.useDefaultServeMux.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go%20Web%20in%20Action/chapter_3_handle_request/1.useDefaultServeMux.go)

3. Although it's convenient to use `DefaultServeMux`, but it's not recommend to use in production. Because defaultServeMux is a global variable. It could be used or modified by some third-party components, which might lead to conflicts. The recommended approach is to use custom ServerMux, see: [2.customisedServeMux.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go%20Web%20in%20Action/chapter_3_handle_request/2.customisedServeMux.go)

3. One restriction of ServeMux is: can not match the URL pattern. `HttpRouter` could make up for that shortcommings, see: [3.httpRouter.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go%20Web%20in%20Action/chapter_3_handle_request/3.httpRouter.go)

### Handle Web Requests