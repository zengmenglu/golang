# 功能
gin提供了go http web框架。主要是封装了go http库，并额外提供路由组的封装，以及前缀树的路由匹配，当需要处理的http请求的路由url较多时，路由匹配性能较好。

# 框架核心技术
* 责任链模式：采用责任链模式，将http的处理函数（中间件）加入到责任链，当匹配到http请求时，执行责任链中的方法。
* 前缀树 & 最长公共子串: 路由树上采用最长公共子串匹配路径，每一种http方法维护一个根节点，当前缀相同时，如果后缀不同，则在树上增加子节点。路由匹配时按照前缀树的方式进行匹配。

# API
## 初始化

新建一个路由有下面两种方式
```
r := gin.New() // 返回一个gin实例
r := gin.Default() // 等价于New() + 2个middleware(Logger+Recovery).
```
两种方式的区别在于Default()在New()的基础上使用gin的Logger()和Recovery()中间件。
* Logger()中间件将日志写入gin.DefaultWriter. gin.DefaultWriter的默认值是os.Stdout。
* Recovery()中间级会recover任何panic，如果有panic，会写入500.

示例：
```
func setRoute() *gin.Engine {
	r := gin.Default() // 创建一个gin实例
	r.GET("/ping", func(c *gin.Context) { // 处理get消息,定义response内容
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r
}
func main() {
	r := setRoute()
	r.Run(":8080") // 可变长参数，0～1个参数，默认监听8080端口，监听并在 localhost:8080 上启动服务
}
```
## 测试
借助httptest包构造http请求，接收http响应，进行测试。
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

## 路由匹配
路由匹配根据http方法和uri相对路径进行匹配，并设定相应的路由处理方法（中间件）。其中Any()接口可以匹配任意http方法的路由。
```
router.GET(path,handle) //path是uri相对路径，handle是路由处理中间件
router.POST(path,handle)
router.DELETE(path,handle)
router.PATCH(path,handle)
router.PUT(path,handle)
router.OPTIONS(path,handle)
router.HEAD(path,handle)

// 上面的方法等价于
router.Handle("GET", path, handle)
router.Handle("POST", path, handle)
...

// Any匹配任何http请求
router.Any(path,handle)
```

## 端口监听
```
router.Run(port...) // 可变长参数，只接收0～1个参数。如果入参为空，默认为8080端口

// 等效于
http.ListenAndServe(port, router)
```

## 中间件
中间件为格式满足下列示例的函数，用于路由处理。
```
type HandlerFunc func(*gin.Context)
```
使用Use()可以为路由添加任意数量的中间件。
```
r.Use(midware1,midware2,...)
```

### gin.Logger()
gin.Logger()中间件默认会将日志写入gin.DefaultWriter. gin.DefaultWriter的默认值是os.Stdout。
```
// 修改gin.Logger记录到文件
f, _ := os.Create("gin.log")
gin.DefaultWriter = io.MultiWriter(f)

// 同时将日志写入文件和控制台
gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
```

### c.Next()

c.Next只能在中间件内部被调用，用于执行责任链中被挂起的后续的处理程序，处理完成之后再返回前一个未执行完成的处理程序。

例：
```
func globalMid1(c *gin.Context) {
	fmt.Println("global-1-A")
	c.Next()
	fmt.Println("global-1-B")
}
func globalMid2(c *gin.Context) {
	fmt.Println("global-2-A")
	c.Next()
	fmt.Println("global-2-B")
}
func mid1(c *gin.Context) {
	fmt.Println("mid-1-A")
	c.Next()
	fmt.Println("mid-1-B")
}

func main() {
	r := gin.Default()
	rGrp := r.Group("", globalMid1)
	rGrp.Use(globalMid2)
	{
		rGrp.GET("/abc", mid1)
	}
	r.Run(":8080")
}
```
上面例子，输入uri "localhost:8080/abc" 执行结果为
```
global-1-A
global-2-A
mid-1-A
mid-1-B
global-2-B
global-1-B
```

## 路由组
可以将有相同处理方法或相同路由前缀的路由放在一个路由组内。为路由组定义的中间件，路由组中每一个路由都会执行。
```
grp := router.Group(relativePath, handle...)

//等价于
grp := router.Group(relativePath)
grp.Use(handle...)
```

示例：
```
func authRoute() *gin.Engine {
	r := gin.New()

	// way 1
	authGrp := r.Group("/way1")
	authGrp.Use(middlewareAuth)
	{
		authGrp.GET("/ping", middleware1Ping)
	}

	// way 2
	authGrp2 := r.Group("/way2", middlewareAuth)
	{
		authGrp2.GET("/ping", middleware1Ping)
	}
	return r
}
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
### uri 参数校验
```
type Uri struct {
	User     string `uri:"user" binding:"required"`
	Password string `uri:"password" binding:"required"`
}
func setUriRoute() *gin.Engine {
	router := gin.Default()
	router.GET("/:user/:password", func(c *gin.Context) {
		var uri Uri
		err := c.ShouldBindUri(&uri)
		if err != nil || uri.User != "myUser" || uri.Password != "myPassword" {
			// 参数校验不通过
			return
		}
	})
	return router
}
```

### POST data 校验
```
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}
func setLoginRoute() *gin.Engine {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		var form LoginForm
		err := c.ShouldBind(&form)
		if err != nil || form.User != "myUser" || form.Password != "myPassword" {
			// 参数校验不通过
			return
		}
	})
	return router
}
```

## 获取POST请求表单中数据、查询字段
api:
```
// c is *gin.Context
// 获取查询字段值
c.Query(key) // 获取不到返回零值
c.Default(key, defaultValue) // 获取不到返回默认值

// 获取POST表单数据
c.PostForm(key) // 获取不到返回零值
c.DefaultPostForm(key, defaultValue) // 获取不到返回默认值
```

示例：

(1) 例如uri：localhost:8080/form_post?id=123&page=1
``` 
id := c.Query("id") // 获取查询字段中id的值, 如果没有这个字段，返回零值""。
page := c.DefaultQuery("page", "0") // 获取查询字段中page的值，如果没有这个字段，返回设定的默认值"0"
```

(2) 例如post表单中包含{"message":"this is message", "name": "this is name"}
``` 
message := c.PostForm("message")
name := c.DefaultPostForm("name", "anonymous") // 如果没有这个字段，返回设定的默认值"anonymous"
```

## 返回数据渲染
返回数据支持json, HTML, xml, yaml, protobuf, string格式。

```
r := gin.Default()
r.GET("/json", func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{ // gin.H is a shortcut for map[string]interface{}
		"message": "json",
	})
})

r.GET("/morejson", func(c *gin.Context) {
	var jsonMsg struct {
		Name    string
		Message string
	}
	jsonMsg.Name = "Nick"
	jsonMsg.Message = "this is message"
	c.JSON(http.StatusOK, jsonMsg)
})

r.GET("/xml", func(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{
		"message": "xml",
	})
})

r.GET("/yaml", func(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{
		"message": "yaml",
	})
})
```

## 重路由

### 重路由
```
c.Redirect(httpCode, newLocation)
```
* http code取值范围201,3XX.
* newLocation 可以为相对路径，则内部路由；也可以是网址，则外部重定向。

### 消息转发
通过http包发出http请求，再通过c.DataFromReader()将结果回写。例：
```
r.GET("/forward", func(c *gin.Context) {
	rsp, err := http.Get("https://www.baidu.com/")
	if err != nil || rsp == nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.DataFromReader(rsp.StatusCode, rsp.ContentLength, rsp.Header.Get("Content-Type"), rsp.Body, nil)
})
```

## 多服务
需要引入```"golang.org/x/sync/errgroup"```包。
* ```errgroup.Group```定义了一个goroutine集合。
* ```g.Go(func())```会调用一个goroutine执行func()。
* ```g.Wait()```阻塞进程直至所有调用```g.Go(func())```的函数返回，并返回第一个不为nil的error。


```
import "golang.org/x/sync/errgroup"
var g errgroup.Group
func main() {
	for serv := range servers { // serv类型*http.Server
		g.Go(func() error {
			return serv.ListenAndServe()
		})
	}
	if err := g.Wait(); err != nil { // 阻塞进程直至所有goroutine返回
		log.Fatal(err)
	}
}
```


> 参考资料
>* gin源码：https://github.com/gin-gonic/gin
>* gin官方文档https://gin-gonic.com/zh-cn/docs/
>* gin中文翻译：https://github.com/skyhee/gin-doc-cn#request
>* validator官方文档: https://godoc.org/github.com/go-playground/validator


