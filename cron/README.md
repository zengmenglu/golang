# 功能
开启定时任务。

# API

## 创建定时器管理器

```
c := cron.New(opts ...Option)
```
* 入参可以为空，则默认解析。
* opts 可以设置下面三个值：
    * 时区：默认值 time.Local
    * 时间解析器
    * chain: 包裹提交的任务，默认值是会recover panic 并且打印错误。

### 设置时区

```
//
cron.New(cron.WithLocation(time.UTC))

// 
nyc, _ := time.LoadLocation("America/New_York")
c := cron.New(cron.WithLocation(nyc))

//
cron.New().AddFunc("CRON_TZ=Asia/Tokyo 0 6 * * ?", ...)

//
c := cron.New()
c.SetLocation("America/New_York")
```

### 设置时间解析器

设置支持解析秒字段。
```
cron.New(
	cron.WithParser(
		cron.NewParser(
			cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
```

下面设置秒字段是必选。
```
cron.New(cron.WithSeconds())
```

### 设置chain



## 添加定时任务

```
id, err := c.AddFunc(spec string, cmd func())
```
* 入参spec是定时器的时间设置，cmd是需要被执行的定时任务。
* 返回值id是该定时任务的ID。


## 定时器时间格式

spec := "min hour day month week"

字段名	|是否必填|	允许的值	|允许的特殊字符
---|---|---|---
分（Minutes）	|Yes|	0-59	|* / , -
时（Hours）	|Yes|	0-23|	* / , -
一个月中的某天（Day of month）	|Yes|	1-31|	* / , - ?
月（Month）	|Yes|	1-12 or JAN-DEC|	* / , -
星期几（Day of week）|	Yes|	0-6 or SUN-SAT  (大小写不敏感)|	* / , - ?

### 支持的符号

* 星号 ( * ) ：匹配字段的所有值
* 斜线 ( / ) ：描述范围的增量，表现为 “N-MAX/x”，first-last/x 的形式，例如 3-59/15 表示此时的第三分钟和此后的每 15 分钟，到59分钟为止。即从 N 开始，使用增量直到该特定范围结束。它不会重复。
* 逗号 ( , ) ：分隔列表中的项目。例如，在 Day of week 使用“MON，WED，FRI”将意味着星期一，星期三和星期五
* 连字符 ( - ) ：连字符用于定义范围。例如，9 - 17 表示从上午 9 点到下午 5 点的每个小时
* 问号 ( ? ) ： 不指定值，用于代替 “ * ”，类似 “ _ ” 的存在。

### 预定义的特殊时间表

输入|	简述|	相当于
---|---|---
@yearly (or @annually)	|一年运行一次，1月1日零时运行	|0 0 1 1 *
@monthly|	每个月运行一次，每个月的第一天零时运行	|0 0 1 * *
@weekly	|每周运行一次，周六/日零时运行	|0 0 * * 0
@daily (or @midnight)	|每天运行一次，零时运行	|0 0 * * *
@hourly	|每小时运行一次，每小时第一个时刻运行	|0 * * * *
@every <duration>|自定义时间段|

### 一些例子

```
c.AddFunc("30 * * * *", func() { fmt.Println("Every hour on the half hour") })
c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
c.AddFunc("@hourly",      func() { fmt.Println("Every hour, starting an hour from now") })
c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
c.AddFunc("@daily", func() { fmt.Println("Every day") })
```


## 启动定时器
```
c.Start()
```
## 停止定时器

```
ctx := c.Stop()
```


> 参考文档
> * 源码：https://github.com/robfig/cron
> * 官方文档： https://godoc.org/github.com/robfig/cron