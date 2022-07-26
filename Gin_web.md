# Gin_web

- Gin框架基本使用
- GORM基本使用
- WEB开发项目实战 

## 一、关于web

- Web是使用http协议交互的应用网络
- Web是通过浏览器/app访问各种资源

<img src="/Users/xujiaxin/Desktop/picture/web.png" alt="web" style="zoom:50%;" />

一个请求对应一个响应：输入一个url，返回一个页面。

### 1、使用`net/http`包：

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)
func sayhello(w http.ResponseWriter,r *http.Request){
	_,err:=fmt.Fprintln(w,"dudu")
	if err != nil {
		log.Println(err)
		return
	}
}
func main(){
	http.HandleFunc("/hello",sayhello)
	err:=http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Println(err)
		return
	}
}
```

在浏览器上输入：`127.0.0.1:8080/hello`会显示：dudu

### 2、创建html文件

```html
<!DOCTYPE html>
<title>dudu</title>
<body>
<h1 style='color:red'>
  hello Golang!
</h1>
<h1>
  hello gin!
</h1>
<img src="https://timgsa.baidu.com/timg?image&quality=80&size=b9999_10000&sec=1600011052622&di=9aeee5de695a40c8d469f0c3980c2d48&imgtype=0&src=http%3A%2F%2Fa4.att.hudong.com%2F22%2F59%2F19300001325156131228593878903.jpg">
</body>
```

```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)
func sayhello(w http.ResponseWriter,r *http.Request){
	html,err:=ioutil.ReadFile("./hello.html")
	if err != nil {
		log.Println(err)
		return
	}
	_,_=w.Write(html)
}
func main(){
	http.HandleFunc("/hello",sayhello)
	err:=http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Println(err)
		return
	}
}
```

在浏览器上输入：`127.0.0.1:8080/hello`会显示：

<img src="/Users/xujiaxin/Desktop/picture/截屏2022-07-12 12.15.16.png" alt="截屏2022-07-12 12.15.16" style="zoom:50%;" />

图片挂了，更换对应的图片地址即可。

## 二、Gin框架介绍

框架是为了让我们更高效的开发，专注于开发设计。

`Gin`是使用Go语言编写的web框架。Go世界里最流行的Web框架，[Github](https://gitee.com/link?target=https%3A%2F%2Fgithub.com%2Fgin-gonic%2Fgin)上有`32K+`star。 基于[httprouter](https://gitee.com/link?target=https%3A%2F%2Fgithub.com%2Fjulienschmidt%2Fhttprouter)开发的Web框架。 [中文文档](https://gitee.com/link?target=https%3A%2F%2Fgin-gonic.com%2Fzh-cn%2Fdocs%2F)齐全，简单易用的轻量级框架。

### 1、安装

```shell
go get -u github.com/gin-gonic/gin
```

### 2、基本示例

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	//返回默认的路由引擎
	r :=gin.Default()
	//GET:请求方式；/hello：请求的路径
	//客户端浏览器等访问/hello路径时，会执行后面的匿名函数
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"msg":"GET",
		})
	})
	
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"POST",
		})
	})
	
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"DELETE",
		})
	})
	//启动http服务，默认端口8080
	err:=r.Run()
	if err != nil {
		log.Println(err)
		return
	}
}
```

从上面的示例可知，gin框架支持RESTful API，简单来说就是：

**客户端和web服务器交互的时候，使用HTTP协议中的四个请求方法代表不同的动作。**

- `GET`用来获取资源

- `POST`用来新建资源

- `PUT`用来更新资源

- `DELETE`用来删除资源

  只要API程序遵循了REST风格，就称其为RESTful API。

  举例：现在编写一个图书管理系统，需要对书籍进行增删改查，编写程序时设计客户端和web服务器交互的方式，通常会这么设计：

  | 请求方法 | URL          | 含义         |
  | :------- | ------------ | ------------ |
  | GET      | /book        | 查询书籍信息 |
  | POST     | /create_book | 创建书籍记录 |
  | POST     | /updata_book | 更新书籍信息 |
  | POST     | /delete_book | 删除书籍信息 |

  现在使用RESTful 风格：

  | 请求方法 | URL   | 含义         |
  | -------- | ----- | ------------ |
  | GET      | /book | 查询书籍信息 |
  | POST     | /book | 创建书籍信息 |
  | PUT      | /book | 更新书籍信息 |
  | DELETE   | /book | 删除书籍信息 |

## 三、模板渲染

`html/template`包实现了数据驱动的模板，用于生成可防止代码注入的安全的HTML内容。它提供了和`text/template`包相同的接口，Go语言中输出HTML的场景都应使用`html/template`这个包。

### 1、go语言模版引擎

Go语言内置了文本模板引擎`text/template`和用于HTML文档的`html/template`。它们的作用机制可以简单归纳如下：

1. 模板文件通常定义为`.tmpl`和`.tpl`为后缀（也可以使用其他的后缀），必须使用`UTF8`编码。
2. 模板文件中使用`{{`和`}}`包裹和标识需要传入的数据。
3. 传给模板这样的数据就可以通过点号（`.`）来访问，如果数据是复杂类型的数据，可以通过{ { .FieldName }}来访问它的字段。
4. 除`{{`和`}}`包裹的内容外，其他内容均不做修改原样输出。

### 2、模版引擎的使用

- 定义模板文件

  根据相关衣服啊规则去编写。

- 解析模板文件

```go
  func (t *Template) Parse(src string) (*Template, error)
func ParseFiles(filenames ...string) (*Template, error)
func ParseGlob(pattern string) (*Template, error)
```

- 模板渲染

​		简单来说就是使用数据去填充模板

```go
func (t *Template) Execute(wr io.Writer, data interface{}) error
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
```

### 3、基本示例

模板语法都包含在`{{`和`}}`中间，其中`{{.}}`中的点表示当前对象。

当我们传入一个结构体对象时，我们可以根据`.`来访问结构体的对应字段。例如：

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<p>Hello {{ . }}</p>
</body>
</html>
```

代码示例：

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

func sayhello(w http.ResponseWriter,r *http.Request){
	//1.定义模板
	//2.解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Println("Parse template failed,err:",err)
		return
	}
	//3.渲染模板
	name:="dudu"
	err=t.Execute(w,name)
	if err != nil {
		log.Println("exec failed,err:",err)
		return
	}
}

func main() {
	http.HandleFunc("/hello",sayhello)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

### 4、Template模板

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<p>u1</p>
<p>Hello {{- .u1.Name -}}</p>
<p>性别： {{- .u1.Gender -}}</p>
<p>年龄 ：{{- .u1.Age -}}</p>
{{/* 遇事不决写注释 */}}
<p>m1</p>
<p>hello {{ .m1.name }}</p>
<p>性别：{{ .m1.Gender }}</p>
<p>年龄：{{ .m1.Age }}</p>
<hr>
{{/*变量*/}}
{{ $v1:=100 }}
{{ $age:=.m1.Age }}
<p>{{ $age }}</p>

<hr>
{{/*条件判断*/}}
{{ if $v1 }}
{{ $v1 }}
{{else}}
无
{{end}}
<hr>
{{/*比较函数*/}}
{{ if lt .m1.Age 22 }}
好好上学
{{else}}
好好工作
{{end}}
<hr>
{{/*range*/}}
{{ range $index,$value :=  .hobby }}
    <p>{{$index}} - {{$value}}</p>
{{else}}
    无爱好
{{end}}
<hr>
{{/*with*/}}
<p>m1</p>
{{ with .m1}}
<p>hello {{ .name }}</p>
<p>性别：{{ .Gender }}</p>
<p>年龄：{{ .Age }}</p>
{{end}}
<hr>
{{/*index*/}}
{{index .hobby 2}}
</body>
</html>
```

代码：

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)
type User struct {
	Name string
	Gender string
	Age int
}
func sayhello(w http.ResponseWriter,r *http.Request){
	//1.定义模板
	//2.解析模板
	t,err:=template.ParseFiles("./hello.tmpl")
	if err != nil {
		log.Println("Parse failed,err:",err)
		return
	}
	//3.渲染模板
	u1:=User{
		Name: "嘟嘟",
		Gender: "女",
		Age: 3,
	}
	m1:=map[string]interface{}{
		"name":"嘟嘟",
		"Gender":"女",
		"Age":18,
	}
	hobbyList:=[]string{
		"篮球",
		"足球",
		"双色球",
	}
	err=t.Execute(w,map[string]interface{}{
		"u1":u1,
		"m1":m1,
		"hobby":hobbyList,
	})
	if err != nil {
		log.Println("exec failed,err:",err)
		return
	}
}

func main() {
	http.HandleFunc("/hello",sayhello)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

#### 4.1自定义模板

Html:

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>自定义模板函数</title>
</head>
<body>
    {{ kua . }}
</body>
</html>
```

代码

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

func f1(w http.ResponseWriter,r *http.Request){
	//定义一个函数kua
	kua:= func(name string)(string,error){
		//要么只有一个返回值；有两个返回值第二个必须是error类型
		return name+"真可爱",nil
	}
	//1.定义模板

	t:=template.New("f.tmpl") //创建一个名字是f.tmpl的模板对象
	//告诉模板引擎，现在多了一个自定义函数kua
	t.Funcs(template.FuncMap{
		"kua":kua,
	})

	//2.解析模板
	_,err:=t.ParseFiles("./f.tmpl")  //要与上面的f.tmpl一致
	if err != nil {
		log.Println("parse failed,err:",err)
		return
	}
	//3.渲染模板
	name:="嘟嘟"
	err=t.Execute(w,name)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/hello",f1)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

#### 4.2模板嵌套

先定义两个tmpl文件

`t .tmpl`

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>tmpl test</title>
</head>
<body>

<h1>测试嵌套template语法</h1>
<hr>
{{/*嵌套了另外一个单独的模板文件*/}}
{{template "ul.tmpl"}}
<hr>
{{/*嵌套一个define定义的模板文件*/}}
{{template "ol.tmpl"}}
<div>你好,{{ . }}</div>
</body>
</html>
{{/*通过define定义一个模板*/}}
{{ define "ol.tmpl"}}
    <ol>
        <li>吃饭</li>
        <li>睡觉</li>
        <li>打豆豆</li>
    </ol>
{{end}}
```

`ul.tmpl`

```html
<ul>
  <li>注释</li>
  <li>日志</li>
  <li>测试</li>
</ul>
```

代码：

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)


func demo1(w http.ResponseWriter,r *http.Request){
	t,err:=template.ParseFiles("./t.tmpl","./ul.tmpl")
	if err != nil {
		log.Println("parse2 failed,err:",err)
		return
	}
	name:="dudu"
	err=t.Execute(w,name)
	if err != nil {
		log.Println(err)
		return
	}
}

func main() {
	http.HandleFunc("/tmplDemo",demo1)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

#### 4.3模板继承block

引入：

先定义两个模板：

`index.tmpl`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>模板继承</title>
    <style>
        * {
            margin: 0;
        }
        .nav {
            width: 100%;
            height: 50px;
            position: fixed;
            top: 0;
            background-color: burlywood;
        }
        .main{
            margin-top:50px;
        }
        .menu {
            width: 20%;
            height: 100%;
            position: fixed;
            left: 0;
            background-color: blueviolet;
        }
        .center {
            text-align: center;
        }
    </style>
</head>
<body>
<div class="nav"></div>
<div class="main">
    <div class="menu"></div>
    <div class="content center">
        <h1>这是index页面</h1>
        <p>Hello {{.}}</p>
    </div>
</div>
</body>
</html>
```

`home.index`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>模板继承</title>
    <style>
        * {
            margin: 0;
        }
        .nav {
            width: 100%;
            height: 50px;
            position: fixed;
            top: 0;
            background-color: burlywood;
        }
        .main{
            margin-top:50px;
        }
        .menu {
            width: 20%;
            height: 100%;
            position: fixed;
            left: 0;
            background-color: blueviolet;
        }
        .center {
            text-align: center;
        }
    </style>
</head>
<body>
<div class="nav"></div>
<div class="main">
    <div class="menu"></div>
    <div class="content center">
        <h1>这是home页面</h1>
        <p>Hello {{.}}</p>
    </div>
</div>
</body>
</html>
```

代码：

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter,r*http.Request){
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./index.tmpl")
	if err != nil {
		log.Println("parse tmpl1 failed,err:",err)
		return
	}
	msg:="dudu"
	//渲染模板
	err=t.Execute(w,msg)
	if err != nil {
		log.Println("exec index failed,err:",err)
		return
	}
}

func home(w http.ResponseWriter,r*http.Request){
	//定义模板
	//解析模板
	t,err:=template.ParseFiles("./home.tmpl")
	if err != nil {
		log.Println("parse tmpl2 failed,err:",err)
		return
	}
	msg:="dudu"
	//渲染模板
	err=t.Execute(w,msg)
	if err != nil {
		log.Println("exec home failed,err:",err)
		return
	}
}

func main() {
	http.HandleFunc("/index",index)
	http.HandleFunc("/home",home)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

可以看到两个模板的内容基本上是一模一样，这时候就可以使用`block`:典型的用法是定义一组根模板，然后通过在其中重新定义块模板进行自定义。

定义三个tmpl文件

`base.tmpl`

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>模板继承</title>
    <style>
        * {
            margin: 0;
        }
        .nav {
            width: 100%;
            height: 50px;
            position: fixed;
            top: 0;
            background-color: burlywood;
        }
        .main{
            margin-top:50px;
        }
        .menu {
            width: 20%;
            height: 100%;
            position: fixed;
            left: 0;
            background-color: blueviolet;
        }
        .center {
            text-align: center;
        }
    </style>
</head>
<body>
<div class="nav"></div>
<div class="main">
    <div class="menu"></div>
    <div class="content center">
       {{block "content" .}}{{end}}
    </div>
</div>
</body>
</html>
```

`index.tmpl`

```html
{{/*继承根模板*/}}
{{template "base.tmpl" .}}
{{/*重新定义块模板*/}}
{{define "content"}}
    <h1>index2页面</h1>
        <p>Hello {{ . }}</p>
{{end}}
```

`home.tmpl`

```html
{{/*继承根模板*/}}
{{template "base.tmpl" .}}
{{/*重新定义块模板*/}}
{{define "content"}}
    <h1>home2页面</h1>
    <p>Hello {{.}}</p>
{{end}}
```

代码：

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter,r*http.Request){
	//定义模板(模板继承的方式)
	//解析模板
	t,err:=template.ParseFiles("./template/base.tmpl","./template/index.tmpl")
	if err != nil {
		log.Println("parse failed,err:",err)
		return
	}
	//渲染模板
	name:="嘟嘟"
	err=t.ExecuteTemplate(w,"index.tmpl",name)
	if err != nil {
		log.Println("exec failed,err:",err)
		return
	}
}

func home(w http.ResponseWriter,r*http.Request){
	//定义模板(模板继承的方式)
	//解析模板
	t,err:=template.ParseFiles("./template/base.tmpl","./template/home.tmpl")
	if err != nil {
		log.Println("parse failed,err:",err)
		return
	}
	//渲染模板
	name:="一一"
	err=t.ExecuteTemplate(w,"home.tmpl",name)
	if err != nil {
		log.Println("exec failed,err:",err)
		return
	}
}

func main() {
	http.HandleFunc("/index",index)
	http.HandleFunc("/home",home)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

#### 4.4模板补充

创建一个`xss.tmpl`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>修改模板引擎标识符号</title>
</head>
<body>
Hello {{ .str1 }}
Hello {{ .str2 |safe}}
</body>
</html>
```

主函数：

```go
package main

import (
	"html/template"
	"log"
	"net/http"
)

func xss(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板
	t,err:=template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML{
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		log.Println("parse xss failed,err:",err)
		return
	}
	//渲染模板
	str1:="<script>alert(123)</script>"
	str2 := "<a href='https://liwenzhou.com'>liwenzhou的博客</a>"
	err=t.Execute(w,map[string]interface{}{
		"str1":str1,
		"str2":str2,

	})
	if err != nil {
		log.Println("exec xss failed,err:",err)
		return
	}

}

func main() {
	http.HandleFunc("/xss",xss)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

```

执行后发现`liwenzhou的博客`是可以点击链接的，是因为我们在`Hello {{ .str2 |safe}}`加了`safe`,需要提前在主函数中定义，且必须在解析函数前定义：

```go
t,err:=template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(s string) template.HTML{
			return template.HTML(s)
		},
	}).ParseFiles("./xss.tmpl")
```

## 四、Gin渲染

### 1、HTML渲染

在`templates`下定义两个目录`posts/index.tmpl`和`users/index.tmpl`内容如下：

`posts/index.tmpl`：

```html
{{define "posts/index.tmpl"}}
    <!DOCTYPE html>
    <html lang="en">

    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>posts/index</title>
    </head>
    <body>
    {{.title}}
    </body>
    </html>
{{end}}
```

`users/index.tmpl`:

```html
{{define "users/index.tmpl"}}
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>users/index</title>
    </head>
    <body>
    {{.title}}
    </body>
    </html>
{{end}}
```

代码：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r:=gin.Default()
	//定义模板
	//解析模板
	//r.LoadHTMLFiles("templates/posts/index.tmpl","templates/users/index.tmpl")
	r.LoadHTMLGlob("templates/**/*") //**表示目录；*表示文件

	//渲染模板
	r.GET("/posts/index", func(c *gin.Context) {
		//http请求
		//gin.H就是一个map[string]interface{}
		c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
			"title":"posts/index.tmpl",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"users/index.tmpl",gin.H{
			"title":"users/index.tmpl",
		})
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}


}

```

### 2、自定义函数模板

`index.tmpl`：

```html
{{define "users/index.tmpl"}}
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <title>users/index</title>
    </head>
    <body>
    {{ .title | safe }}
    </body>
    </html>
{{end}}
```

代码：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)

func main() {
	r:=gin.Default()
	//定义模板
	//gin框架给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	//解析模板
	//r.LoadHTMLFiles("templates/posts/index.tmpl","templates/users/index.tmpl")
	r.LoadHTMLGlob("templates/**/*")     //**表示目录；*表示文件

	//渲染模板

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"users/index.tmpl",gin.H{
			"title":"<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

编译后可以出现可点击网页链接

### 3、静态文件处理

创建一个`statics`下面创建一个`index.css`

`index.css`:

```css
body{
    background-color: cadetblue;
}
```

然后在`/templates/users/index.tmpl`:

```html
{{define "users/index.tmpl"}}
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <link rel="stylesheet" href="/xxx/index.css">
        <title>users/index</title>
    </head>
    <body>
    {{ .title | safe }}
    </body>
    </html>
{{end}}
```

主函数：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
)
//静态文件：
//html页面上用到的样式文件：.css js文件 图片
func main() {
	r:=gin.Default()
	//定义模板
	//加载静态文件
	r.Static("/xxx","./statics")      //在/xxx下找到文件 
	//gin框架给模板添加自定义函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML{
			return template.HTML(str)
		},
	})
	//解析模板
	//r.LoadHTMLFiles("templates/posts/index.tmpl","templates/users/index.tmpl")
	r.LoadHTMLGlob("templates/**/*") //**表示目录；*表示文件

	//渲染模板

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"users/index.tmpl",gin.H{
			"title":"<a href='https://liwenzhou.com'>李文周的博客</a>",
		})
	})
	
	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

### 4、JSON渲染

代码：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func main() {
	r:=gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//1.使用map
		data1 := map[string]interface{}{
			"name": "dudu",
			"msg":  "hello",
			"age":  3,
		}
		c.JSON(http.StatusOK,data1)
	})

	r.GET("/json1", func(c *gin.Context) {
		//等同于map
		data2:=gin.H{
			"name":"yiyi",
			"msg":"hello",
			"age":2,
		}
		c.JSON(http.StatusOK,data2)
	})

	r.GET("/json2", func(c *gin.Context) {
		//2.使用结构体
		type data struct{
			Name string   `json:"name"`     //灵活使用tag来对结构字字段做定制化操作
			Age int        `json:"xxx"`
			Msg string
		}

		d:=data{
			Name:"8hao",
			Age:1,
			Msg:"hello",
		}
		c.JSON(http.StatusOK,d)      //json的序列化，字段名必须大写才能导出。
	})
	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

<u>总结</u>：一般有两种方法一种是map，一种是结构体；

gin内置了一种`gin.H`等同于map。

### 5、XML渲染

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	r := gin.Default()
	// gin.H 是map[string]interface{}的缩写
	r.GET("/someXML", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.XML(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/moreXML", func(c *gin.Context) {
		// 方法二：使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)
	})
	r.Run(":8080")
}
```

### 6、YAML渲染

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func main() {
	r := gin.Default()

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok", "status": http.StatusOK})
	})
	r.Run(":8080")
}
```

### 7、protobuf渲染

```go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)


func main() {
	r := gin.Default()
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8080")
}
```

## 五、获取参数

### 1、获取querystring参数

代码：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r:=gin.Default()
	//querystring

	//GET请求URL ？后面的是querystring等参数
	//key=value，多个key-value用 & 符号连接
	//eq: /web?query=yiyi&age=3
	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的query string 参数
		//name:=c.Query("query")                             //1.通过Query获取请求中携带的querystring参数

		name:=c.DefaultQuery("query","嘟嘟")  //2.取不到就用指定的默认值
		age:=c.DefaultQuery("age","3")

		//name,ok:=c.GetQuery("query")                       //3.取到返回（值，true）取不到第二个参数就返回（""，false）
		//if !ok{
		//	//取不到
		//	name="嘟嘟"
		//}
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
		})
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

### 2、获取form参数

新建两个hmtl文件

`login.html`：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>login</title>
</head>
<body>
<form action="/login" method="post" novalidate autocomplete="off">
    <div>
        <label for="username">username:</label>
        <input type="text" name="username" id="username">
    </div>

    <div>
        <label for="password">password:</label>
        <input type="password" name="password" id="password">
    </div>

    <div>
        <input type="submit" value="登陆">
    </div>

</form>
</body>
</html>
```

`index.html`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>index</title>
</head>
<body>
<h1>Hello {{ .Name }}!</h1>
<p>你的密码是：{{ .Password }}</p>
</body>
</html>
```

代码：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//获取form表单提交的参数

func main(){
	r:=gin.Default()
	r.LoadHTMLFiles("./login.html","index.html")

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK,"login.html",nil)
	})

	//login posts请求
	r.POST("/login", func(c *gin.Context) {
		//获取form表单提交的数据
		//方法1
		//username:=c.PostForm("username")   //取到返回值，取不到返回空字符串
		//password:=c.PostForm("password")

		//方法2
		//username:=c.DefaultPostForm("username","somebody")
		//password:=c.DefaultPostForm("zzz","***")

		//方法3
		username,ok:=c.GetPostForm("username")
		if !ok{
			username="sb"
		}

		password,ok:=c.GetPostForm("password")
		if !ok{
			password="123566"
		}


		c.HTML(http.StatusOK,"index.html",gin.H{
			"Name":username,
			"Password":password,
		})
	})

	
	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

<u>切记：</u>一个请求对应一个响应。

第一次返回get方法；点击登陆后返回post方法。

### 3、获取path/URL参数

代码：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//获取请求的path(URL)参数，返回的都是字符串类型
//注意URL的匹配不要冲突

func main() {
	r:=gin.Default()

	r.GET("/user/:name/:age/:ds", func(c *gin.Context) {
		//获取参数路径
		name:=c.Param("name")
		age:=c.Param("age")   //string类型
		ds:=c.Param("ds")
		c.JSON(http.StatusOK,gin.H{
			"name":name,
			"age":age,
			"ds":ds,
		})
	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year:=c.Param("year")
		month:=c.Param("month")
		c.JSON(http.StatusOK,gin.H{
			"year":year,
			"month":month,
		})
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

### 4、参数绑定

之前的方法：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
type UserInfo struct {
	username string
	password string

}
func main() {
	r:=gin.Default()
	r.GET("/user", func(c *gin.Context) {
		username:=c.Query("username")
		password:=c.Query("password")
		u:=UserInfo{
			username,
			password,
		}
		fmt.Printf("%#v\n",u)
		c.JSON(http.StatusOK,gin.H{
			"msg":"ok",
		})
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

}
```

现在用结构体：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
type UserInfo struct {
	Username string  `form:"username" json:"user"`
	Password string	 `form:"password" json:"pwd"`

}
func main() {
	r:=gin.Default()

	r.GET("/user", func(c *gin.Context) {

		var u UserInfo  //声明一个UserInfo类型的变量
		err:=c.ShouldBind(&u)    //传指针
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})

	r.POST("/form", func(c *gin.Context) {

		var u UserInfo  //声明一个UserInfo类型的变量
		err:=c.ShouldBind(&u)    //传指针
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})

	r.POST("/json", func(c *gin.Context) {

		var u UserInfo  //声明一个UserInfo类型的变量
		err:=c.ShouldBind(&u)    //传指针
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else {
			fmt.Printf("%#v\n",u)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

}
```

`.ShouldBind()`函数功能强大，能够基于请求自动提取`JSON`、`form表单`和`QueryString`类型的数据，并把值绑定到指定的结构体对象。

### 5、文件上传

#### 5.1单个文件上传

新建一个`index.html`:

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>index</title>
</head>
<body>

<form action="/upload" method="post" enctype="multipart/form-data">
    <input type="file" name="f1">
    <input type="submit" value="上传">
</form>

</body>
</html>
```

代码：

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
)

func main() {
	r:=gin.Default()

	r.LoadHTMLGlob("./index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		f,err:=c.FormFile("f1") //从请求中获取携带的参数
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else {
			//从请求中读取文件，将读取到的文件保存在本地（服务端本地）
			//filePath:=fmt.Sprintf("./%s",f.Filename)
			filePath:=path.Join("./",f.Filename)
			_=c.SaveUploadedFile(f,filePath)
			c.JSON(http.StatusOK,gin.H{
				"status":"ok",
			})
		}

	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}

```

#### 5.2多个文件上传

新建一个`index.html`:

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
<form action="http://localhost:8000/upload" method="post" enctype="multipart/form-data">
    上传文件:<input type="file" name="files" multiple>
    <input type="submit" value="提交">
</form>
</body>
</html>
```

代码：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// gin的helloWorld

func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 限制表单上传大小 8MB，默认为32MB
	r.LoadHTMLGlob("./index.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"index.html",nil)
	})
  
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})
	//默认端口号是8080
	r.Run(":8000")
}
```

### 6、重定向

#### 6.1HTTP重定向

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r:=gin.Default()

	r.GET("/index", func(c *gin.Context) {
		//c.JSON(http.StatusOK,gin.H{
		//	"status":"ok",
		//})
		//跳转到sogo.com
		c.Redirect(http.StatusMovedPermanently,"http://www.sogo.com")
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

}

```

#### 6.2路由重定向

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r:=gin.Default()

	r.GET("/a", func(c *gin.Context) {
		//跳转到/b对应的处理路由函数
		c.Request.URL.Path="/b"    //把请求的URL修改
		r.HandleContext(c)     //继续后续的处理

	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"b",
		})
	})

	err:=r.Run(":9090")
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

}
```

## 六、Gin路由

### 1、普通路由

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r:=gin.Default()
	//访问/index的GET请求会走这一条处理逻辑
	//路由
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"POST",
		})
	})

	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"DELETE",
		})
	})

	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"PUT",
		})
	})

	
	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

}
```

这种方法是`GET`、`POST`、`PUT`、`DELETE`分开写的，比较易理解但是繁琐。

下面使用`r.Any`方法，可以匹配所有情的方法：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r:=gin.Default()
	//访问/index的GET请求会走这一条处理逻辑
	//路由

	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK,gin.H{
			"method":"GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK,gin.H{
				"method":"POST",
			})
		case http.MethodPut:
			c.JSON(http.StatusOK,gin.H{
			"method":"PUT",
			})
		case http.MethodDelete:
			c.JSON(http.StatusOK,gin.H{
			"method":"DELETE",
			})
		default:
			fmt.Println("dudu")

		}

	})
	
	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

}
```

### 2、路由组

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r:=gin.Default()

	//路由组的组
	//把公用的前缀提取出来，创建一个路由组
	
	userGroup:=r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {})
		userGroup.GET("/login", func(c *gin.Context) {})
		userGroup.POST("/login", func(c *gin.Context) {})
	}
	
	shopGroup:=r.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {})
		shopGroup.GET("/cart", func(c *gin.Context) {})
		shopGroup.POST("/checkout", func(c *gin.Context) {})
	}
	
	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}

}
```

路由组：有共同URL前缀的路由划分为一个路由组。用一对`{}`包裹同组的路由

## 七、中间件

### 1、中间件概念

![中间件](/Users/xujiaxin/Desktop/picture/中间件.png)



中间件适合处理一些公共的业务逻辑

### 2、中间件一般写法和处理方法

写法一：`m1`和想要处理函数写在一起

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//
func indexHandler(c *gin.Context){
	fmt.Println("index...")
	c.JSON(http.StatusOK,gin.H{
		"msg":"index",
	})
}

//定义一个中间件m1:统计处理请求函数的耗时
func m1(c *gin.Context){
	fmt.Println("m1 in ...")
	start:=time.Now()
	fmt.Println("start:",start)
	c.Next()
	cost:=time.Since(start)
	fmt.Printf("cost:%v\n",cost)
	fmt.Println("m1 out ...")

}

func main() {
	r:=gin.Default()

	r.GET("/index", m1,indexHandler)

	r.GET("/shop",m1, func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"shop",
		})
	})
	r.GET("/user",m1, func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"user",
		})
	})


	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}

```

写法二：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//
func indexHandler(c *gin.Context){
	fmt.Println("index...")
	c.JSON(http.StatusOK,gin.H{
		"msg":"index",
	})
}

//定义一个中间件m1:统计处理请求函数的耗时
func m1(c *gin.Context){
	fmt.Println("m1 in ...")
	start:=time.Now()
	fmt.Println("start:",start)
	c.Next()
	cost:=time.Since(start)
	fmt.Printf("cost:%v\n",cost)
	fmt.Println("m1 out ...")

}

func main() {
	r:=gin.Default()

	r.Use(m1)

	r.GET("/index", indexHandler)

	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"shop",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"user",
		})
	})


	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}

```

如果有两个中间件分别调用了`c.Next()`执行顺序是怎么样的？

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//
func indexHandler(c *gin.Context){
	fmt.Println("index...")
	c.JSON(http.StatusOK,gin.H{
		"msg":"index",
	})
}

//定义一个中间件m1:统计处理请求函数的耗时
func m1(c *gin.Context){
	fmt.Println("m1 in ...")
	start:=time.Now()
	c.Next()
	cost:=time.Since(start)
	fmt.Printf("cost:%v\n",cost)
	fmt.Println("m1 out ...")

}

func m2(c *gin.Context){
	fmt.Println("m2 in ...")
	c.Next()
	fmt.Println("m2 out ...")
}

func main() {
	r:=gin.Default()

	r.Use(m1,m2)

	r.GET("/index", indexHandler)

	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"shop",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"user",
		})
	})


	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}

```

以下就是执行顺序，大家自己思考下为什么：

```shell
m1 in ...
m2 in ...
index...
m2 out ...
cost:143.208µs
m1 out ...
[GIN] 2022/07/15 - 16:16:15 | 200 |     191.167µs |       127.0.0.1 | GET      "/index"
```

调用`c.Abort`:阻止调用后续函数

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context){
	fmt.Println("index...")
	c.JSON(http.StatusOK,gin.H{
		"msg":"index",
	})
}

//定义一个中间件m1:统计处理请求函数的耗时
func m1(c *gin.Context){
	fmt.Println("m1 in ...")
	start:=time.Now()
	c.Next()   //调用后续的处理函数
	cost:=time.Since(start)
	fmt.Printf("cost:%v\n",cost)
	fmt.Println("m1 out ...")

}

func m2(c *gin.Context){
	fmt.Println("m2 in ...")
	c.Abort()   //阻止调用后续的处理函数
	// return 加上的话，立即结束m2函数
	fmt.Println("m2 out ...")
}

func main() {
	r:=gin.Default()

	r.Use(m1,m2)

	r.GET("/index", indexHandler)

	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"shop",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"user",
		})
	})


	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}

```

执行结果：

```shell
m1 in ...
m2 in ...
m2 out ...
cost:3.583µs
m1 out ...
```

可以看到浏览器也不会返回任何信息，因为`indexHandler`没有被执行。

### 3、业务具体写法

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//
func indexHandler(c *gin.Context){
	fmt.Println("index...")
	c.JSON(http.StatusOK,gin.H{
		"msg":"index",
	})
}

//定义一个中间件m1:统计处理请求函数的耗时
func m1(c *gin.Context){
	fmt.Println("m1 in ...")
	start:=time.Now()
	c.Next()   //调用后续的处理函数
	cost:=time.Since(start)
	fmt.Printf("cost:%v\n",cost)
	fmt.Println("m1 out ...")

}

func  authMiddleWare(doCheck bool)gin.HandlerFunc {
	//查询/连接数据库
	//其他的一些准备参数
	return func(c *gin.Context) {
		if doCheck{
			//存放具体的逻辑
			//是否登陆的判断
			//if 是登陆用户
			//c.Next()
			//else
			//c.Abort()
		}else{
			c.Next()
		}

	}
}
func main() {
	r:=gin.Default()

	r.Use(m1,authMiddleWare(true))

	r.GET("/index", indexHandler)

	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"shop",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"user",
		})
	})


	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

一般写成闭包的形式

### 4、路由组注册中间件

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func  authMiddleWare(doCheck bool)gin.HandlerFunc {
	//查询/连接数据库
	//其他的一些准备参数
	return func(c *gin.Context) {
		if doCheck{
			//存放具体的逻辑
			//是否登陆的判断
			//if 是登陆用户
			//c.Next()
			//else
			//c.Abort()
		}else{
			c.Next()
		}

	}
}
func main() {
	r:=gin.Default()

	//路由组注册中间件方法1：
	xxGroup:=r.Group("/xx",authMiddleWare(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"msg":"xxGroup",
			})
		})
	}

	//路由组注册中间件方法2：
	xx2Group:=r.Group("/xx")
	xx2Group.Use(authMiddleWare(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"msg":"xx2Group",
			})
		})
	}


	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg":"user",
		})
	})


	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

#### 注意事项一、

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context){
	fmt.Println("index...")
	name,ok:=c.Get("name")
	if !ok {
		name="匿名用户"
	}
	c.JSON(http.StatusOK,gin.H{
		"msg":name,
	})
}

//定义一个中间件m1:统计处理请求函数的耗时
func m1(c *gin.Context){
	fmt.Println("m1 in ...")
	start:=time.Now()
	c.Next()   //调用后续的处理函数
	cost:=time.Since(start)
	fmt.Printf("cost:%v\n",cost)
	fmt.Println("m1 out ...")

}

func m2(c *gin.Context){
	fmt.Println("m2 in ...")
	//c.Abort()   //阻止调用后续的处理函数
	c.Set("name","dudu")
	// return 加上的话，立即结束m2函数
	fmt.Println("m2 out ...")
}



func  authMiddleWare(doCheck bool)gin.HandlerFunc {
	//查询/连接数据库
	//其他的一些准备参数
	return func(c *gin.Context) {
		if doCheck{
			//存放具体的逻辑
			//是否登陆的判断
			//if 是登陆用户
			//c.Next()
			c.Next()
			//else
			//c.Abort()
		}else{
			c.Next()
		}

	}
}
func main() {
	r:=gin.Default()

	r.Use(m1,m2,authMiddleWare(true))

	r.GET("/index", indexHandler)

	err:=r.Run()
	if err != nil {
		log.Println("http server start failed,err:",err)
		return
	}
}
```

`m2`中的`c.Set("name","dudu")`在处理函数中调用：

```go
func indexHandler(c *gin.Context){
	fmt.Println("index...")
	name,ok:=c.Get("name")
	if !ok {
		name="匿名用户"
	}
	c.JSON(http.StatusOK,gin.H{
		"msg":name,
	})
}
```

#### 注意事项二、

```go
package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//返回默认的路由引擎
	//r :=gin.Default()   //默认调用了Logger()和Recovery()
	r:=gin.New()
	//GET:请求方式；/hello：请求的路径
	//客户端浏览器等访问/hello路径时，会执行后面的匿名函数
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"msg":"GET",
		})
	})


	//启动http服务，默认端口8080
	err:=r.Run()
	if err != nil {
		log.Println(err)
		return
	}
} 
```

`gin.Default()`默认使用了`Logger`和`Recovery`中间件，其中：

- `Logger`中间件将日志写入`gin.DefaultWriter`，即使配置了`GIN_MODE=release`。
- `Recovery`中间件会recover任何`panic`。如果有panic的话，会写入500响应码。

如果不想使用上面两个默认的中间件，可以使用`gin.New()`新建一个没有任何默认中间件的路由。