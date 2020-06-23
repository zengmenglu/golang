# 功能
提供go http web框架。主要是封装了go http库，并额外提供路由组的封装，以及前缀树的路由匹配，当需要处理的http请求的路由url较多时，可以增加性能。

# 框架核心技术
* 责任链模式：采用责任链模式，将http的处理函数加入到责任链，当匹配到http请求时，执行责任链中的方法。
* 前缀树 & 最长公共子串: 路由树上采用最长公共子串匹配路径，每一种http方法维护一个根节点，当前缀相同时，如果后缀不同，则在树上增加子节点。路由匹配时按照前缀树的方式进行匹配。

# API
## 初始化
```
router:= gin.New() // 返回一个gin实例
router:= gin.Default() // 等价于New() + 2个middleware(Logger+Recovery).logger打印上下文信息，recover处理panic信息并且返回500
```

## 路由分组
```
grp := router.Group(relativePath, handle...)

//等价于
grp := router.Group(relativePath)
grp.Use(handle...)
```

## 匹配HTTP请求
```
router.GET(path,handle)
router.POST(path,handle)
router.DELETE(path,handle)
router.PATCH(path,handle)
router.PUT(path,handle)
router.OPTIONS(path,handle)
router.HEAD(path,handle)

// 等价于
router.Handle("GET", path, handle)

// Any 匹配任何http GET, POST, PUT, PATCH, HEAD, OPTIONS, DELETE, CONNECT, TRACE 请求
router.Any(path,handle)
```

## 监听HTTP请求
```
router.Run(port...) // 可变长参数，0～1个参数。如果入参为空，默认为8080端口

// 等效于
http.ListenAndServe(port, router)
```

## Binding

Binding的功能是将请求中携带的数据与目标数据结构进行匹配的过程。可以通过Binding获取请求中的数据，同时gin框架结合了[validator](https://godoc.org/github.com/go-playground/validator)
的功能，可以进行数据验证。
常用：
```
var data Data
err := c.ShouldBind(&data) // for json or xml, c is *gin.Context
err := c.ShouldBindJSON(&data) // for json
err := c.ShouldBindXML(&data) // for xml
err := c.ShouldBindUri(&data) // for uri
err := c.ShouldBindQuery(&data) // for query
```

## 获取POST请求表单以及查询字段
```
// 例如uri：localhost:8080/form_post?id=123&page=1
id := c.Query("id") // 获取查询字段中id的值, 如果没有这个字段，返回零值。c is *gin.Context
page := c.DefaultQuery("page", "0") // 获取查询字段中page的值，如果没有这个字段，返回设定的默认值

// 例如post表单中包含{"message":"this is message", "name": "this is name"}
message := c.PostForm("message") // 获取post表单数据
name := c.DefaultPostForm("name", "anonymous") // if no data,return default value
```

## 返回数据渲染

返回数据支持json, HTML, xml, yaml, protobuf, string格式。

# Get Start

## basic http request handle
```
import (
   "github.com/gin-gonic/gin"
   "net/http"
)
func setRoute() *gin.Engine {
   r := gin.Default() // 创建一个gin实例
   r.GET("/ping", func(c *gin.Context) { // 处理get消息,定义response body的内容
      c.JSON(http.StatusOK, gin.H{
         "message": "pong",
      })
   })
   return r
}
func main() {
   r := setRoute()
   r.Run(":8080”) // 可变长参数，0～1个参数，默认监听8080端口，监听并在 localhost:8080 上启动服务
}
```
借助httptest包进行测试：
```
import (
   "net/http"
   "net/http/httptest"
   "testing"
)
func TestPingRoute(t *testing.T) {
   route := setRoute()
   rsp := httptest.NewRecorder()
   req := httptest.NewRequest(http.MethodGet, "/ping", nil)
   route.ServeHTTP(rsp, req)
   if rsp.Code != http.StatusOK || rsp.Body.String() != "{\"message\":\"pong\"}" {
      t.Error("fail")
   }
}
```

> 参考资料
>* gin源码：https://github.com/gin-gonic/gin
>* gin官方文档https://gin-gonic.com/zh-cn/docs/
>* gin中文翻译：https://github.com/skyhee/gin-doc-cn#request
>* validator官方文档: https://godoc.org/github.com/go-playground/validator


