package web

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"lzy/framework/session"
	"lzy/framework/util"
	"net/http"
	"path"
	"strings"
)

const (
	MAIN_PAGE = "main.html"
)

//缓存html模板
var templates = make(map[string]*template.Template)

//存放路由
var routeHandle = make(map[string]http.HandlerFunc)

//Session管理器
var sessionManager = session.NewSessionManager()

//当前会话session
var currentSession session.Session = nil

//登录路由
var loginRoute string

//端口
var port string

var sessionMaxAge int64

//框架自定义路径
//全局错误页面路径
var errFileName string = "error.html"

//全局css路径
var cssDirPath string = "../view/static/css/"

//全局js路径
var jsDirPath string = "../view/static/js/"

//全局image路径
var imagPath string = "../view/static/image/"

var fontsPath string = "../view/static/fonts/"

//全局html路径
var HtmlDirPath string = "../view/html/"

var dbConfig string = "../conf/db.properties"

//初始化当前session
func initCurrentSession(w http.ResponseWriter, r *http.Request) {

	if util.IsNotZero(sessionMaxAge) {
		sessionManager.SetMaxAge(sessionMaxAge)
	}

	currentSession = sessionManager.BeginSession(w, r)
	//启动一条协程定时判断session是否超时
	//该协程生命周期跟随应用的生命周期
	fmt.Println("current session ", currentSession)
	//go sessionManager.GC(currentSession.GetId())
}

func SetSessionMaxAge(t int64) {
	sessionMaxAge = t
}

func GetSession() session.Session {
	return currentSession
}

func SetDbConfig(s string) {
	dbConfig = s
}

//添加路由
func AddHandler(route string, h http.HandlerFunc) {
	routeHandle[route] = h
}

//设置错误页面
func SetErrFileName(fileName string) {
	errFileName = fileName
}

//端口
func SetPort(p string) {
	port = p
}

//登录路由
func SetLoginRoute(route string) {
	loginRoute = route
}

//设置css目录路径
func SetCssDirPath(path string) {
	cssDirPath = path
}

//设置js目录路径
func SetJsDirPath(path string) {
	jsDirPath = path
}

//设置image目录路径
func SetImagePath(path string) {
	imagPath = path
}

func Logout(w http.ResponseWriter, r *http.Request) {
	//sessionManager.Destroy(w, r)
	http.Redirect(w, r, loginRoute, 302)
}

//处理所有静态文件路由
func HandleStaticResource(mux *http.ServeMux) {
	mux.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) { //框架规定静态路由"/static/"
		filePath := r.URL.Path
		filePath = filePath[len("/static/"):]
		ext := filePath[strings.LastIndex(filePath, "."):] //取后缀
		fmt.Println("{{{{{{{{{{{{", ext)
		switch ext {
		case ".css":
			w.Header().Set("Content-Type", "text/css")
			filePath = cssDirPath + filePath
		case ".map":
			w.Header().Set("Content-Type", "text/css")
			filePath = cssDirPath + filePath
		case ".js":
			w.Header().Set("Content-Type", "text/javascript")
			filePath = jsDirPath + filePath
		case ".jpg":
			filePath = imagPath + filePath
		case ".png":
			filePath = imagPath + filePath
		default:
			w.Header().Set("Content-Type", "text/html")
			filePath = fontsPath + filePath
		}

		if util.IsExistsFile(filePath) {
			//将服务端的一个文件内容读写到 http.Response-Writer
			//并返回给请求来源的 *http.Request客户端
			http.ServeFile(w, r, filePath)
		}
	})
}

//安全分发路由
func SecurityHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() { //处理全局错误
			if err, ok := recover().(error); ok { //判断recover函数返回的类型
				//				http.Error(w, err.Error(), http.StatusInternalServerError)
				fmt.Println(err.Error())

				//				w.WriteHeader(http.StatusInternalServerError)
				//				ExecuteHtml(w, errFileName, err)
				//				log.Println(err.Error())

				//TODO  正式环境写日志文件中
				//				log.Println(string(debug.Stack()))
			}
		}()

		urlPath := r.URL.Path
		if urlPath == loginRoute && r.Method == "POST" { //提交登录表单后，根据判断开始会话
			initCurrentSession(w, r) //确保当前会话是该请求的cookie

		} else if urlPath != loginRoute { //其它路由
			fmt.Println("current ssssss", currentSession)
			initCurrentSession(w, r) //确保当前会话是该请求的cookie,如果超时，则重新初始化一个，空的session
			fmt.Println(currentSession)
			value := currentSession.Get("userId")
			fmt.Println("session get suerId value ", value)
			if value == nil || value == "" {
				requestType := r.Header.Get("X-Requested-With")
				fmt.Println("requestType ---------> ", requestType)

				if requestType == "XMLHttpRequest" {
					io.WriteString(w, "error")
					return
				} else {
					http.Redirect(w, r, loginRoute, 302)
				}
			} else {
				sessionManager.Update(w, r)
			}
		}
		handler(w, r)
	}
}

//初始化缓存模板
func InitCacheTemplate(filePath string) {
	//fmt.Println("path==========>", filePath)
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		util.HandleErr(err)
	}

	for _, f := range files { //遍历目录下所有文件
		//fmt.Println("fileName===========>", f.Name())
		if f.IsDir() { //如果是目录，则递归它
			InitCacheTemplate(filePath + f.Name() + "/")

		}
		fn := f.Name()
		if ext := path.Ext(fn); ext != ".html" {
			continue
		}
		//文件路径
		t := template.Must(template.ParseFiles(filePath + fn)) //缓存模板
		templates[fn] = t
	}
}

//执行文件，返回给客户端(生产环境使用)
func ExecuteHtml(w http.ResponseWriter, fileName string, rows interface{}) {
	w.Header().Set("Content-Type", "text/html")
	t := templates[fileName]
	if t == nil {
		panic("not registered " + fileName + "template...")
		return
	}
	err := t.Execute(w, rows)
	util.HandleErr(err)
}

//执行文件，返回给客户端(开发环境使用)
func DevEnvExecuteHtml(w http.ResponseWriter, fileName string, rows interface{}) {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles(HtmlDirPath + fileName)
	util.HandleErr(err)
	t.Execute(w, rows)
}

//启动路由
func StartRoute() {
	if !util.IsNotEmpty(port) {
		port = ":8689" //框架默认端口
	}
	if !util.IsNotEmpty(loginRoute) {
		loginRoute = "/system/toLogin"
	}
	mux := http.NewServeMux()
	if l := len(routeHandle); l > 0 {
		for k, v := range routeHandle {
			mux.HandleFunc(k, SecurityHandler(v))
			delete(routeHandle, k)
		}
	} else {
		log.Println("not registered anny route handle...")
	}

	//处理静态资源
	HandleStaticResource(mux)
	err := http.ListenAndServe(port, mux)
	util.HandleErr(err)
}
