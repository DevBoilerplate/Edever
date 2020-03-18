package cmd

import (
	"cobra.new/utils"
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "用于获取任意版本模板",
	Long: `
edever package -d 即可自动获取最新版(默认为Gitee)
edever package -d -r [github/gitee](默认Gitee) 选择目标仓库
edever package -l 列出仓库所有的发行版(默认为Gitee)
edever package -d -t (tag) 更新到指定的版本，默认为最新(默认为Gitee)
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
							fmt.Printf("没有查询到您请求的指定tag版本:\t%v\n", version.Tag)
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
	rootCmd.AddCommand(packageCmd)
	packageCmd.Flags().StringVarP(&repo, "repo", "r", "gitee", "edever package -d -r [github/gitee]")
	packageCmd.Flags().BoolVarP(&list, "list", "l", false, "edever package -l")
	packageCmd.Flags().StringVar(&temp, "temp", "sample", "edever init --temp[sample/sqlite3/ts]")
	packageCmd.Flags().StringVarP(&tag, "tag", "t", "latest", "edever package -t (tag)")
	packageCmd.Flags().BoolVarP(&down, "down", "d", false, "edever package -d [-r/-t]")
}
