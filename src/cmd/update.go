package cmd

import (
	"cobra.new/utils"
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "更新Edever",
	Long: `
默认gitee, tag为latest
edever update -d [-t (tag name)] [-r (github/gitee)]
edever update -l [-r (github/gitee)]
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("使用edever version即可查看本地版本号\n更新方法请参考edever update -h\n")
		latestVersion := utils.GetLatest(repo, "edever")
		fmt.Printf("最新版本: \t%v\n创建时间: \t%v\n\n", latestVersion.Tag, latestVersion.Created)
		if list {
			sources := utils.ListAll(repo, "edever")
			fmt.Println("\n可更新的发行版如下")
			for _, value := range sources {
				fmt.Printf("%v\t\t\t%v\n", value.Tag, value.Created)
				fmt.Println(value.Assets)
			}
		}

		if down {
			switch runtime.GOOS {
			case "windows":
				{
					if tag == "latest" {
						utils.GetByBrowser(latestVersion.Assets["edever-win.zip"])
					}

					if tag != "latest" {
						version := utils.ListTagVersion(repo, tag, "edever")
						if version.Assets == nil {
							fmt.Printf("没有查询到您请求的指定tag版本:\t%v\n", tag)
							fmt.Printf("可以指定获取的版本为:\t%v\n", utils.ListTags(repo, "edever"))
						} else {
							version := utils.ListTagVersion(repo, tag, "edever")
							fmt.Printf("检测到: \t%v\n更新时间: \t%v\n", version.Tag, version.Created)
							utils.GetByBrowser(version.Assets["edever-win.zip"])
						}
					}
				}
			case "linux":
				{
					if tag == "latest" {
						utils.GetByBrowser(latestVersion.Assets["edever-linux.zip"])
					}
					if tag != "latest" {
						version := utils.ListTagVersion(repo, tag, "edever")
						if version.Assets == nil {
							fmt.Printf("没有查询到您请求的指定tag版本:\t%v\n", version.Tag)
							fmt.Printf("可以指定获取的版本为:\t%v\n", utils.ListTags(repo, "edever"))
						} else {
							version := utils.ListTagVersion(repo, tag, "edever")
							fmt.Printf("检测到: \t%v\n更新时间: \t%v\n", version.Tag, version.Created)
							utils.GetByBrowser(version.Assets["edever-linux.zip"])
						}
					}
				}
			case "darwin":
				{
					if tag == "latest" {
						utils.GetByBrowser(latestVersion.Assets["edever-darwin.zip"])
					}
					if tag != "latest" {
						version := utils.ListTagVersion(repo, tag, "edever")
						if version.Assets == nil {
							fmt.Printf("没有查询到您请求的指定tag版本:\t%v\n", version.Tag)
							fmt.Printf("可以指定获取的版本为:\t%v\n", utils.ListTags(repo, "edever"))
						} else {
							version := utils.ListTagVersion(repo, tag, "edever")
							fmt.Printf("检测到: \t%v\n更新时间: \t%v\n", version.Tag, version.Created)
							utils.GetByBrowser(version.Assets["edever-mac.zip"])
						}
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&repo, "repo", "r", "gitee", "指定目标仓库")
	updateCmd.Flags().BoolVarP(&list, "list", "l", false, "列出所有可更新")
	updateCmd.Flags().StringVarP(&tag, "tag", "t", "latest", "指定tag")
	updateCmd.Flags().BoolVarP(&down, "down", "d", false, "是否下载")
}
