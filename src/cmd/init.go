package cmd

import (
	"cobra.new/utils"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	remote string
	projName string
	temp string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "用于获取最新模板初始化项目",
	Long: `
从远程获取项目模板
edever init -r [github/gitee] 从Gitee获取(默认从Gitee)
edever init -s 选择含有sqlite3的模板(默认不含)
edever init -d 自定义clone的项目名称(默认为仓库名称)
edever init -r gitee -d test -s 生成目录为test的不含sqlite3的工程
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("正在从远端拉取模板...")
		if remote == "gitee" && temp == "sample" {
			fmt.Println("正在从Gitee拉取sample模板...")
			utils.GitTemp("https://gitee.com/HerbertHe/electron-react-antd-antv-cli.git", projName)
		}
		if remote == "gitee" && temp == "sqlite3" {
			fmt.Println("正在从Gitee拉取含sqlite3的sample模板...")
			utils.GitTemp("https://gitee.com/HerbertHe/electron-react-antd-antv-sqlite3-cli.git", projName)
		}
		if remote == "github" && temp == "sample" {
			fmt.Println("正在从GitHub拉取模板...")
			utils.GitTemp("https://github.com/HerbertHe/electron-react-antd-antv-cli.git", projName)
		}
		if remote == "github" && temp == "sqlite3" {
			fmt.Println("正在从GitHub拉取含sqlite3的模板...")
			utils.GitTemp("https://github.com/HerbertHe/electron-react-antd-antv-sqlite3-cli.git", projName)
		}
		if remote == "gitee" && temp == "ts" {
			fmt.Println("正在从Gitee拉取ts版模板...")
			utils.GitTemp("https://gitee.com/HerbertHe/electron-react-antd-antv-ts-cli.git", projName)
		}
		if remote == "github" && temp == "ts" {
			fmt.Println("正在从GitHub拉取ts版模板...")
			utils.GitTemp("https://github.com/HerbertHe/electron-react-antd-antv-ts-cli.git", projName)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&remote, "repo", "r", "gitee", "edever init -r [github/gitee]")
	//initCmd.Flags().BoolVarP(&sqlite, "sqlite3", "s", false, "edever init -s [true/false]")
	initCmd.Flags().StringVar(&temp, "temp", "sample", "edever init --temp[sample/sqlite3/ts]")
	initCmd.Flags().StringVarP(&projName, "dirname", "d", "ElectronProjectTemplate", "edever init -d (dirname)")
}
