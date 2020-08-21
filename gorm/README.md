# 简介
ORM是通过实例对象的语法，完成关系型数据库的操作的技术，是"对象-关系映射"（Object/Relational Mapping）的缩写。
* 数据库的表（table） --> 类（class）
* 记录（record，行数据）--> 对象（object）
* 字段（field）--> 对象的属性（attribute）

而 gorm 是 Go 语言实现的 ORM。ORM封装了SQL语句，开发者只使用面向对象编程，与数据对象交互，而不用关心底层数据库。

# CRUD 操作
数据库的基本操作有四种：create（新建）、read（读取）、update（更新）和delete（删除），简称 CRUD。



> 参考
> * gorm 源码：https://github.com/go-gorm/gorm
> * orm实例教程：http://www.ruanyifeng.com/blog/2019/02/orm-tutorial.html


