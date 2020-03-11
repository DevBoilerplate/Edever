package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "显示Edever版本信息",
	Long: `
edever version 显示当前版本信息
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Edever当前版本为: %v\n", "v0.0.1")
		fmt.Printf("当前系统: %v\n", runtime.GOOS)
		fmt.Printf("系统架构: %v\n", runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
