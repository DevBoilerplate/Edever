package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// 用于直接构建应用程序
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "用于编译应用程序",
	Long: `
用于打包构建不同平台的应用程序，默认编译平台为本地环境
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.Flags().StringVar(&platform, "build", runtime.GOOS, "edever build")
}
