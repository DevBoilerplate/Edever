package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"net/http"

	"github.com/spf13/cobra"
	"runtime"
)

var (
	repo string
	list bool
	tag string
)

// 获取最新版
func getLatest(repo string) {
	if repo == "gitee" {
		u, _ := url.Parse("https://gitee.com/api/v5/repos/HerbertHe/Edever/releases/latest")
		res, err := http.Get(u.String())
		if err != nil {
			fmt.Println(err)
		}
		result, err := ioutil.ReadAll(res.Body)
		_ = res.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
		var dataMap map[string]interface{}
		if err := json.Unmarshal([]byte(result), &dataMap); err == nil {
			fmt.Printf("最新版本: %v\n创建时间: %v\n", dataMap["tag_name"], dataMap["created_at"])
			//fmt.Println(dataMap["assets"])
		}
	}
}

// 获取所有可下载版本
func listAll(repo string) {
	if repo == "gitee" {
		u, _ := url.Parse("https://gitee.com/api/v5/repos/HerbertHe/Edever/releases?page=1&per_page=20")
		res, err := http.Get(u.String())
		if err != nil {
			fmt.Println(err)
		}
		result, err := ioutil.ReadAll(res.Body)
		_ = res.Body.Close()
		if err != nil {
			fmt.Println(err)
		}

		var dataS []map[string]interface{}
		if err := json.Unmarshal(result, &dataS); err == nil {
			fmt.Println("可供更新发行版如下:")
			for _, value := range dataS {
				fmt.Printf("%v\t\t\t%v", value["tag_name"], value["created_at"])
			}
		}

	}
}

// 下载
func downloadFile() {

}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新Edever",
	Long: `
执行命令默认将自动执行更新！
edever update -r [github/gitee](默认GitHub) 选择目标仓库
edever update -l 列出所有的发行版
edever update -t (tag) 更新到指定的版本，默认为最新
`,
	Run: func(cmd *cobra.Command, args []string) {
		getLatest("gitee")
		if list {
			listAll("gitee")
		}
		switch runtime.GOOS {
		case "windows": {
			getLatest(repo)
		}
		case "linux": fmt.Print()
		case "darwin":fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&repo, "repo", "r", "github", "edever update -r [github/gitee]")
	updateCmd.Flags().BoolVarP(&list, "list", "l", false, "edever update -l")
	updateCmd.Flags().StringVarP(&tag, "tag", "t", "latest", "edever update -t (tag)")
}
