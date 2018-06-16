package util

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"

	"os"
	"os/exec"
	"reflect"
	"strings"
)

//------------------ util_struct --------------------------
//响应消息工具
type Message struct {
	Code    int
	Message string
}

type Err struct {
}

func NewErr() *Err {
	return &Err{}
}
func NewMessage() *Message {
	return &Message{Code: 200, Message: "success"}
}

func (e *Err) Error(err string) string {
	return fmt.Sprintf("now err is ------>>> %s", err)
}

//----------------- util_handle_string ------------------------------

func IsExistsMap(m map[string]string, k string) bool {
	s := m[k]

	if IsEmpty(s) {
		return false
	}

	return true
}

//字符串非空判断
func IsNotEmpty(s string) bool {
	if len(s) == 0 || s == "" {
		return false
	}

	return true
}

//字符串非空判断
func IsEmpty(s string) bool {
	if s == "" || len(s) == 0 {
		return true
	}

	return false
}

//数字非零判断
func IsNotZero(i int64) bool {
	var j int64 = 0
	if i == j {
		return false
	}
	return true
}

//利用反射判断是否是string类型
func IsStringT(t reflect.Type) bool {
	if t.Name() == "string" {
		return true
	}
	return false
}

//判断切片是否为nil 长度是否小于或等于0
func CheckSlick(sli []string) bool {
	if sli == nil || len(sli) <= 0 {
		return true
	}
	return false

}

//利用反射判断是否是int64类型
func IsIntT(t reflect.Type) bool {

	if t.Name() == "int64" {
		return true
	}
	return false
}

func SubString(str string, begin, length int) (substr string) {
	lth := len(str)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(str[begin:end])
}

//------------------------ util_read_file ------------------------------
//一次性读取
func ReadAll(filePath string) ([]byte, error) {

	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

//分块读取并传入处理函数，边读边处理
func ReadBlock(filePath string, bufSize int, handle func([]byte)) error {

	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}

	buf := make([]byte, bufSize)
	bufR := bufio.NewReader(f)

	for {
		n, err := bufR.Read(buf)
		handle(buf[:n])

		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

//每次读一行
func ReadLine(filePath string, lineN int) ([]string, error) {

	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		fmt.Printf("open %s error", filePath)
		return nil, err
	}

	bufR := bufio.NewReader(f)

	b := make([]string, lineN)
	tag := 0

	for {

		if tag == lineN {
			break
		}
		line, _, err := bufR.ReadLine()

		if err != nil {
			if err == io.EOF {
				return b, nil
			}
			fmt.Printf("read %s error %s in %d line", filePath, err.Error(), tag)
			return nil, err
		}
		b[tag] = string(line)

		tag++
	}

	return b, nil
}

//---------util_checkerr-------------
func CheckErr(e error, msg *Message) {
	if e != nil {
		fmt.Println(e.Error()) //测试和开发环境直接打印，正式环境将写入日志文件
		//		msg.Code = 500
		//		msg.Message = e.Error()
		//		panic(e)
	}

}
func HandleErr(e error) {
	if e != nil {
		fmt.Println(e.Error())
		//		log.Fatal(e) //测试和开发环境直接打印，正式环境将写入日志文件
		//		panic(e)
	}
}

//--------------util_file_path
//获得当前路径
func GetCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//获的上一级目录
func GetParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

//判断文件是否存在
func IsExistsFile(path string) bool {
	//os.Stat()这个是获取文件的信息
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	//判断一个文件是否已经存在的错误
	return os.IsExist(err)
}

//生成uuid
func GetUuid() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	//加密
	return base64.URLEncoding.EncodeToString(b)
}

//冒泡排序切片
func BubbleSortSlick(values []string) {

	for i := 0; i < len(values)-1; i++ {
		flag := true
		for j := 0; j < len(values)-i-1; j++ {

			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]

				flag = false
			}
		}

		if flag == true {
			break
		}
	}
}

//冒泡排序map
func BubbleSortMap(values map[int]map[string]string) {

	var str, str2 string = "", ""
	for i := 0; i < len(values)-1; i++ {
		flag := true
		for j := 0; j < len(values)-i-1; j++ {

			str = values[j]["Id"]
			str2 = values[j+1]["Id"]
			if str > str2 {
				values[j], values[j+1] = values[j+1], values[j]

				flag = false
			}
		}

		if flag == true {
			break
		}
	}
}
