package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
)

func SendGet(link string) []byte {
	u, _ := url.Parse(link)
	res, err := http.Get(u.String())
	if err != nil {
		fmt.Println(err)
	}
	result, err := ioutil.ReadAll(res.Body)
	_ = res.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func GitTemp(link string, projName string) {
	gitShell := exec.Command("git", "clone", link, projName)
	_, err := gitShell.Output()
	if err != nil {
		fmt.Printf("模板拉取报错:\n%v", err)
	} else {
		fmt.Println("模板拉取成功！使用请参考工程目录下的README.md")
		fmt.Printf("工程位于: %v", projName)
	}
}

var commands = map[string]string{
	"windows": "start",
	"darwin": "open",
	"linux": "xdg-open",
}

// 打开浏览发送请求
func GetByBrowser(url string) {
	fmt.Println("正在跳转: \t" + url)
	cmd := exec.Command(commands[runtime.GOOS] + url)
	_ = cmd.Start()
}