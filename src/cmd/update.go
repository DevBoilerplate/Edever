package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
)

var (
	repo string
	list bool
	down bool
	tag  string
)

func getLatest(repo string) (name string, created string, backDown map[string]string) {
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
			assets := make(map[string]string)
			for _, value := range dataMap["assets"].([]interface{}) {
				item := value.(map[string]interface{})
				if item["name"] != nil {
					assets[item["name"].(string)] = item["browser_download_url"].(string)
				}
			}
			return dataMap["tag_name"].(string), dataMap["created_at"].(string), assets
		}
	}
	if repo == "github" {
		u, _ := url.Parse("https://api.github.com/repos/HerbertHe/Edever/releases/latest")
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
			assets := make(map[string]string)
			for _, value := range dataMap["assets"].([]interface{}) {
				item := value.(map[string]interface{})
				if item["name"] != nil {
					assets[item["name"].(string)] = item["browser_download_url"].(string)
				}
			}
			return dataMap["tag_name"].(string), dataMap["created_at"].(string), assets
		}
	}
	return "", "", nil
}

// 获取所有可下载版本
func listAll(repo string) {
	// Gitee
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

	// Github
	if repo == "github" {
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
//func downloadFile() {
//
//}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新Edever",
	Long: `
edever update -d 即可自动更新(默认为Gitee)
edever update -d -r [github/gitee](默认Gitee) 选择目标仓库
edever update -l 列出仓库所有的发行版(默认为Gitee)
edever update -d -t (tag) 更新到指定的版本，默认为最新(默认为Gitee)
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("使用edever version即可查看本地版本号\n更新方法请参考edever update -h")
		name, created, assets := getLatest(repo)
		fmt.Printf("最新版本: %v\n创建时间: %v\n", name, created)
		if list {
			listAll(repo)
		}
		if down {
			switch runtime.GOOS {
			case "windows":
				{
					if tag == "latest" {
						fmt.Println(assets["edever-win.zip"])
					}
				}
			case "linux":
				{
					if tag == "latest" {
						fmt.Println(assets["edever-linux.zip"])
					}
				}
			case "darwin":
				{
					if tag == "latest" {
						fmt.Println(assets["edever-darwin.zip"])
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&repo, "repo", "r", "gitee", "edever update -d -r [github/gitee]")
	updateCmd.Flags().BoolVarP(&list, "list", "l", false, "edever update -l")
	updateCmd.Flags().StringVarP(&tag, "tag", "t", "latest", "edever update -t (tag)")
	updateCmd.Flags().BoolVarP(&down, "down", "d", false, "edever update -d [-r/-t]")
}
