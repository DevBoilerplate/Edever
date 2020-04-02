package cmd

import (
	"fmt"

	"cobra.new/utils"
	"github.com/spf13/cobra"
)

var (
	repo string        // 目标仓库
	list bool          // 确认列出
	down bool          // 确认下载
	tag  string        // tag指定
	temp string        // package的模板选择
	projectName string    // 指定Clone的仓库
	platform string       // 指定编译平台
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "用于获取最新模板初始化项目",
	Long: `
从远程获取项目模板
edever init [-d (自定义项目名称)] [--temp (sample/sqlite3/ts)] [-r (github/gitee)] 
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("正在从远端拉取模板...")
		if repo == "gitee" && temp == "sample" {
			fmt.Println("正在从Gitee拉取sample模板...")
			utils.GitTemp("https://gitee.com/HerbertHe/electron-react-antd-antv-cli.git", projectName)
		}
		if repo == "gitee" && temp == "sqlite3" {
			fmt.Println("正在从Gitee拉取含sqlite3的sample模板...")
			utils.GitTemp("https://gitee.com/HerbertHe/electron-react-antd-antv-sqlite3-cli.git", projectName)
		}
		if repo == "github" && temp == "sample" {
			fmt.Println("正在从GitHub拉取模板...")
			utils.GitTemp("https://github.com/HerbertHe/electron-react-antd-antv-cli.git", projectName)
		}
		if repo == "github" && temp == "sqlite3" {
			fmt.Println("正在从GitHub拉取含sqlite3的模板...")
			utils.GitTemp("https://github.com/HerbertHe/electron-react-antd-antv-sqlite3-cli.git", projectName)
		}
		if repo == "gitee" && temp == "ts" {
			fmt.Println("正在从Gitee拉取ts版模板...")
			utils.GitTemp("https://gitee.com/HerbertHe/electron-react-antd-antv-ts-cli.git", projectName)
		}
		if repo == "github" && temp == "ts" {
			fmt.Println("正在从GitHub拉取ts版模板...")
			utils.GitTemp("https://github.com/HerbertHe/electron-react-antd-antv-ts-cli.git", projectName)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&repo, "repo", "r", "gitee", "edever init -r [github/gitee]")
	initCmd.Flags().StringVar(&temp, "temp", "sample", "edever init --temp[sample/sqlite3/ts]")
	initCmd.Flags().StringVarP(&projectName, "dirname", "d", "ElectronProjectTemplate", "edever init -d (dirname)")
}
