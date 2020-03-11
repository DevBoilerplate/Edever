package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	remote string
	sqlite bool
	projName string
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "用于获取模板初始化项目",
	Long: `
从远程获取项目模板
edever init -r gitee 从Gitee获取(默认从GitHub)
edever init -s 选择含有sqlite3的模板(默认不含)
edever init -d 自定义clone的项目名称(默认为仓库名称)
edever init -r gitee -d test -s 生成目录为test的不含sqlite3的工程
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("正在从远端拉取模板...")
		if remote == "gitee" && !sqlite {
			fmt.Println("正在从Gitee拉取模板...")
			gitShell := exec.Command("git", "clone", "https://gitee.com/HerbertHe/electron-react-antd-antv-cli.git", projName)
			_, err := gitShell.Output()
			if err != nil {
				fmt.Printf("模板拉取报错:\n%v", err)
			} else {
				fmt.Println("模板拉取成功！使用请参考工程目录下的README.md")
				fmt.Printf("工程位于: %v", projName)
			}
		}
		if remote == "gitee" && sqlite {
			fmt.Println("正在从Gitee拉取含sqlite3的模板...")
			gitShell := exec.Command("git","clone", "https://gitee.com/HerbertHe/electron-react-antd-antv-sqlite3-cli.git", projName)
			_, err := gitShell.Output()
			if err != nil {
				fmt.Printf("模板拉取报错:\n%v", err)
			} else {
				fmt.Println("模板拉取成功！使用请参考工程目录下的README.md")
				fmt.Printf("工程位于: %v", projName)
			}
		}
		if remote == "github" && !sqlite {
			fmt.Println("正在从GitHub拉取模板...")
			gitShell := exec.Command("git", "clone", "https://github.com/HerbertHe/electron-react-antd-antv-cli.git", projName)
			_, err := gitShell.Output()
			if err != nil {
				fmt.Printf("模板拉取报错:\n%v", err)
			} else {
				fmt.Println("模板拉取成功！使用请参考工程目录下的README.md")
				fmt.Printf("工程位于: %v", projName)
			}
		}
		if remote == "github" && sqlite {
			fmt.Println("正在从GitHub拉取含sqlite3的模板...")
			gitShell := exec.Command("git", "clone", "https://github.com/HerbertHe/electron-react-antd-antv-sqlite3-cli.git", projName)
			_, err := gitShell.Output()
			if err != nil {
				fmt.Printf("模板拉取报错:\n%v", err)
			} else {
				fmt.Println("模板拉取成功！使用请参考工程目录下的README.md")
				fmt.Printf("工程位于: %v", projName)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringVarP(&remote, "repo", "r", "github", "edever init -r [github/gitee]")
	initCmd.Flags().BoolVarP(&sqlite, "sqlite3", "s", false, "edever init -s [true/false]")
	initCmd.Flags().StringVarP(&projName, "dirname", "d", "", "edever init -d (dirname)")
}
