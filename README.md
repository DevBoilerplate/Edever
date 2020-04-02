# Edever

Edever是electron、antd、antv集成开发项目的命令行工具，用于项目模板的快速获取，将持续关注针对框架新特性的开发

## 环境依赖

* Git

## 版本更新须知

* 2020/03/18: 重构代码结构和命令行结构，`v0.1.0`发布
* 2020/03/14: 本次更新在初始化模板命令`edever init`中废除了`-s`的flag，请在`v0.0.1`之后的版本使用`--temp`命令替代，默认下载仓库已改成`Gitee`

## 如何使用

* 请在release中下载最新对应的可执行文件，只提供64位可执行文件

| 文件名 | 平台 |
| :---- | ---- |
| edever-win | Windows |
| edever-mac | MacOS |
| edever-linux | Linux |

* 对于Windows平台

> 局部命令, 在工程文件夹下`PowerShell`打开

```shell
./edever -h
```

> 全局命令, 要求把工具所在目录添加到`Path`环境变量

```shell
edever -h
```

## 完整使用说明

> `v0.0.1`只支持带有`-s`参数的`init`, 之后的稳定版本为为大版本更新

| edever子命令 | 使用说明 |
| ---- | ---- |
| help, -h, --help | 获取帮助 |
| init | 初始化工程模板（`v0.0.1之后版本支持ts`） |
| package | 获取指定的工程模板 |
| update | 获取Edever版本更新 |
| version | 打印当前的Edever版本信息 |

### `edever init`子命令参数

| 参数 | 说明 | 合法值 |
| ---- | ---- | ---- |
| -d, --dirname | 指定保存工程模板的目录名 | 默认值为ElectronProjectTemplate |
| -h, --help | 获取edever init命令帮助 | 无 |
| -r, --repo | 指定获取的仓库 | gitee/github，默认gitee |
| --temp | 指定获取的模板类型 | sample/sqlite3/ts，默认sample |

> `--temp`取代了`v0.0.1`的`-s`命令

### `edever update`子命令参数

| 参数 | 说明 | 合法值 |
| ---- | ---- | ---- |
| -d, --down | 是否更新控制 | 加上才能启动下载 |
| -h, --help | 获取edever update命令帮助 | 无 |
| -l, --list | 列出所有可更新的版本 | 无 |
| -r, --repo | 指定获取仓库 | github/gitee，默认gitee |
| -t, --tag | 指定tag更新 | 默认为latest |

> 所有的下载命令都需要加上`-d`, `-r` `-t`可选

### `edever package`子命令参数

| 参数 | 说明 | 合法值 |
| ---- | ---- | ---- |
| -d, --down | 是否更新控制 | 加上才能启动下载 |
| -h, --help | 获取edever update命令帮助 | 无 |
| -l, --list | 列出所有可更新的版本 | 无 |
| -r, --repo | 指定获取仓库 | github/gitee，默认gitee |
| -t, --tag | 指定tag更新 | 默认为latest |
| --temp | 指定模板 | sample/ts/sqlite3, 默认为sample |

## TODOS

* [ ] 完善build子命令，支持`edever`接管编译`package.json`对于不同编译配置的需求

## Gitee&&GitHub

GitHub: [HerbertHe/Edever](https://github.com/HerbertHe/Edever)

Gitee: [HerbertHe/Edever](https://gitee.com/HerbertHe/Edever)
