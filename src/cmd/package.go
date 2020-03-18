package cmd

import (
	"cobra.new/utils"
	"fmt"
	"github.com/spf13/cobra"
)

// packageCmd represents the package command
var packageCmd = &cobra.Command{
	Use:   "package",
	Short: "用于获取任意版本模板",
	Long: `
默认sample，tag默认latest
edever package -d [-r (github/gitee)] [--temp (sample/ts/sqlite3)] [-t (tag名)]
edever package -l [--temp (sample/ts/sqlite3)]
`,
	Run: func(cmd *cobra.Command, args []string) {
		if down {
			switch temp {
			case "sample":
				{
				if tag == "latest" {
					latestVersion := utils.GetLatest(repo, "eraac")
					fmt.Printf("最新版本: \t%v\n创建时间: \t%v\n\n", latestVersion.Tag, latestVersion.Created)
					utils.GetByBrowser(latestVersion.Assets["url"])
				}
				if tag != "latest" {
					version := utils.ListTagVersion(repo, tag, "eraac")
					if version.Assets == nil {
						fmt.Printf("没有查询到您请求的指定tag版本:\t%v\n", tag)
						fmt.Printf("可以指定获取的版本为:\t%v\n", utils.ListTags(repo, "eraac"))
					} else {
						version := utils.ListTagVersion(repo, tag, "eraac")
						fmt.Printf("检测到: \t%v\n更新时间: \t%v\n", version.Tag, version.Created)
						utils.GetByBrowser(version.Assets["url"])
					}
				}
				}
			case "sqlite3":
				{
					if tag == "latest" {
						latestVersion := utils.GetLatest(repo, "eraasc")
						fmt.Printf("最新版本: \t%v\n创建时间: \t%v\n\n", latestVersion.Tag, latestVersion.Created)
						utils.GetByBrowser(latestVersion.Assets["url"])
					}
					if tag != "latest" {
						version := utils.ListTagVersion(repo, tag, "eraasc")
						if version.Assets == nil {
							fmt.Printf("没有查询到您请求的指定tag版本:\t%v\n", tag)
							fmt.Printf("可以指定获取的版本为:\t%v\n", utils.ListTags(repo, "eraasc"))
						} else {
							version := utils.ListTagVersion(repo, tag, "eraasc")
							fmt.Printf("检测到: \t%v\n更新时间: \t%v\n", version.Tag, version.Created)
							utils.GetByBrowser(version.Assets["url"])
						}
					}
				}
			case "ts":
				{
					if tag == "latest" {
						latestVersion := utils.GetLatest(repo, "eraatc")
						fmt.Printf("最新版本: \t%v\n创建时间: \t%v\n\n", latestVersion.Tag, latestVersion.Created)
						utils.GetByBrowser(latestVersion.Assets["url"])
					}
					if tag != "latest" {
						version := utils.ListTagVersion(repo, tag, "eraatc")
						if version.Assets == nil {
							fmt.Printf("没有查询到您请求的指定tag版本:\t%v\n", tag)
							fmt.Printf("可以指定获取的版本为:\t%v\n", utils.ListTags(repo, "eraatc"))
						} else {
							version := utils.ListTagVersion(repo, tag, "eraatc")
							fmt.Printf("检测到: \t%v\n更新时间: \t%v\n", version.Tag, version.Created)
							utils.GetByBrowser(version.Assets["url"])
						}
					}
				}
			default:
				{
					fmt.Println("请指定正确的模板名称[sample/sqlite3/ts]")
				}
			}
		}
		if list {
			switch temp {
			case "sample":
				{
					sources := utils.ListAll(repo, "eraac")
					fmt.Println("\n可更新的发行版如下")
					for _, value := range sources {
						fmt.Printf("%v\t\t\t%v\n", value.Tag, value.Created)
						fmt.Println(value.Assets)
					}
				}
			case "sqlite3":
				{
					sources := utils.ListAll(repo, "eraasc")
					fmt.Println("\n可更新的发行版如下")
					for _, value := range sources {
						fmt.Printf("%v\t\t\t%v\n", value.Tag, value.Created)
						fmt.Println(value.Assets)
					}

				}
			case "ts":
				{
					sources := utils.ListAll(repo, "eraatc")
					fmt.Println("\n可更新的发行版如下")
					for _, value := range sources {
						fmt.Printf("%v\t\t\t%v\n", value.Tag, value.Created)
						fmt.Println(value.Assets)
					}
				}
			default:
				{
					fmt.Println("请指定正确的模板名称[sample/sqlite3/ts]")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(packageCmd)
	packageCmd.Flags().StringVarP(&repo, "repo", "r", "gitee", "指定仓库")
	packageCmd.Flags().BoolVarP(&list, "list", "l", false, "列出所有")
	packageCmd.Flags().StringVar(&temp, "temp", "sample", "指定模板")
	packageCmd.Flags().StringVarP(&tag, "tag", "t", "latest", "指定tag")
	packageCmd.Flags().BoolVarP(&down, "down", "d", false, "是否获取")
}
