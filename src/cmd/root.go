package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd代表当没有子命令调用的时候的基础的命令(根命令)
var rootCmd = &cobra.Command{
	Use:   "edever",
	Short: "Let you develop Electron App quickly!",  // 简短介绍下面是完整介绍
	Long: `
edever是用于快速生成electron开发模板的shell工具, 集成了antd, antv, sqlite3等服务。
author: Herbert He
GitHub仓库位于https://github.com/HerbertHe
Gitee仓库位于https://gitee.com/HerbertHe

edever init -h 获取模板初始化帮助
`,
}

// 将所有的子命令添加到根命令并且合适的设置flags，这由main.main()只会在rootCmd调用一次
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig从配置文件和环境变量读取
func initConfig() {
	if cfgFile != "" {
		// 从flag使用config文件
		viper.SetConfigFile(cfgFile)
	} else {
		// 寻找home目录
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// 在home目录下搜名称为".edever"的配置，没有扩展
		viper.AddConfigPath(home)
		viper.SetConfigName(".edever")
	}

	viper.AutomaticEnv() // 读取匹配的环境变量值

	// 如果配置文件存在，读取他
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
