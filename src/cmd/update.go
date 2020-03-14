package cmd

import (
	"cobra.new/utils"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var (
	repo string
	list bool
	down bool
	tag  string
)

type Source struct {
	tag     string
	created string
	assets  map[string]string
}

func getLatest(repo string) Source {
	if repo == "gitee" {
		result := utils.SendGet("https://gitee.com/api/v5/repos/HerbertHe/Edever/releases/latest")
		var dataMap map[string]interface{}
		if err := json.Unmarshal([]byte(result), &dataMap); err == nil {
			assets := make(map[string]string)
			var source Source
			for _, value := range dataMap["assets"].([]interface{}) {
				item := value.(map[string]interface{})
				if item["name"] != nil {
					assets[item["name"].(string)] = item["browser_download_url"].(string)
				}
			}
			source.tag = dataMap["tag_name"].(string)
			source.created = dataMap["created_at"].(string)
			source.assets = assets
			return source
		}
		return Source{
			tag:     "",
			created: "",
			assets:  nil,
		}
	}
	return Source{
		tag:     "",
		created: "",
		assets:  nil,
	}
}

// 获取所有的tag
func listTags(repo string) []string {
	// Gitee
	if repo == "gitee" {
		result := utils.SendGet("https://gitee.com/api/v5/repos/HerbertHe/Edever/tags")
		var res []map[string]interface{}
		if err := json.Unmarshal(result, &res); err == nil {
			var backTags []string
			for _, value := range res {
				backTags = append(backTags, value["name"].(string))
			}
			return backTags
		}
		return nil
	}
	return nil
}

// 获取指定tag版本
func listTagVersion(repo, tag string) Source {
	//	Gitee
	if repo == "gitee" {
		result := utils.SendGet("https://gitee.com/api/v5/repos/HerbertHe/Edever/releases/tags/" + tag)
		var dataMap map[string]interface{}
		if err := json.Unmarshal([]byte(result), &dataMap); err == nil {
			assets := make(map[string]string)
			var source Source
			for _, value := range dataMap["assets"].([]interface{}) {
				item := value.(map[string]interface{})
				if item["name"] != nil {
					assets[item["name"].(string)] = item["browser_download_url"].(string)
				}
			}
			source.tag = dataMap["tag_name"].(string)
			source.created = dataMap["created_at"].(string)
			source.assets = assets
			return source
		}
		return Source{
			tag:     "",
			created: "",
			assets:  nil,
		}
	}
	return Source{
		tag:     "",
		created: "",
		assets:  nil,
	}
}

// 获取所有可下载版本
func listAll(repo string) []Source {
	// Gitee
	if repo == "gitee" {
		result := utils.SendGet("https://gitee.com/api/v5/repos/HerbertHe/Edever/releases?page=1&per_page=20")
		var dataS []map[string]interface{}
		if err := json.Unmarshal(result, &dataS); err == nil {
			var backArray []Source
			for _, value := range dataS {
				var backMap Source
				cache := make(map[string]string)
				for _, v := range value["assets"].([]interface{}) {
					item := v.(map[string]interface{})
					if item["name"] != nil {
						cache[item["name"].(string)] = item["browser_download_url"].(string)
					}
				}
				backMap.tag = value["tag_name"].(string)
				backMap.created = value["created_at"].(string)
				backMap.assets = cache
				backArray = append(backArray, backMap)
			}
			return backArray
		}
		return nil
	}
	return nil
}

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
		fmt.Printf("使用edever version即可查看本地版本号\n更新方法请参考edever update -h\n\n")
		latestVersion := getLatest(repo)
		fmt.Printf("最新版本: %v\n创建时间: %v\n", latestVersion.tag, latestVersion.created)
		if list {
			sources := listAll(repo)
			fmt.Println("\n可更新的发行版如下")
			for _, value := range sources {
				fmt.Printf("%v\t\t\t%v\n", value.tag, value.created)
				fmt.Println(value.assets)
			}
		}

		if down {
			switch runtime.GOOS {
			case "windows":
				{
					if tag == "latest" {
						fmt.Println(latestVersion.assets["edever-win.zip"])
					}

					if tag != "latest" {
						fmt.Println(listTagVersion(repo, tag))
					}
				}
			case "linux":
				{
					if tag == "latest" {
						fmt.Println(latestVersion.assets["edever-linux.zip"])
					}
					if tag != "latest" {
						fmt.Println(listTagVersion(repo, tag))
					}
				}
			case "darwin":
				{
					if tag == "latest" {
						fmt.Println(latestVersion.assets["edever-darwin.zip"])
					}
					if tag != "latest" {
						fmt.Println(listTagVersion(repo, tag))
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
