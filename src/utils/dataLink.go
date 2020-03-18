package utils

import "encoding/json"

// 数据类型
type Source struct {
	Tag     string
	Created string
	Assets  map[string]string
}

// Gitee仓库API地址
var giteeRemote = map[string]string{
	"eraac": "https://gitee.com/api/v5/repos/HerbertHe/electron-react-antd-antv-cli",
	"eraasc": "https://gitee.com/api/v5/repos/HerbertHe/electron-react-antd-antv-sqlite3-cli",
	"eraatc": "https://gitee.com/api/v5/repos/HerbertHe/electron-react-antd-antv-ts-cli",
	"edever": "https://gitee.com/api/v5/repos/HerbertHe/Edever",
}

// 构造Release
func MakeGiteeReleases (remote string) string {
	return giteeRemote[remote] + "/releases?page=1&per_page=20"
}

// 构造最新
func MakeGiteeReleaseLatest(remote string) string {
	return giteeRemote[remote] + "/releases/latest"
}

// 构造tags

func MakeGiteeTags(remote string) string {
	return giteeRemote[remote] + "/tags"
}

// 获取最新更新
func GetLatest(repo, remote string) Source {
	if repo == "gitee" {
		result := SendGet(MakeGiteeReleaseLatest(remote))
		var dataMap map[string]interface{}
		if err := json.Unmarshal([]byte(result), &dataMap); err == nil {
			assets := make(map[string]string)
			var source Source
			for _, value := range dataMap["assets"].([]interface{}) {
				item := value.(map[string]interface{})
				if item["name"] != nil {
					assets[item["name"].(string)] = item["browser_download_url"].(string)
				}
			}
			source.Tag = dataMap["tag_name"].(string)
			source.Created = dataMap["created_at"].(string)
			source.Assets = assets
			return source
		}
		return Source{
			Tag:     "",
			Created: "",
			Assets:  nil,
		}
	}
	return Source{
		Tag:     "",
		Created: "",
		Assets:  nil,
	}
}

// 获取所有的tag
func ListTags(repo, remote string) []string {
	// Gitee
	if repo == "gitee" {
		result := SendGet(MakeGiteeTags(remote))
		var res []map[string]interface{}
		if err := json.Unmarshal(result, &res); err == nil {
			var backTags []string
			for _, value := range res {
				backTags = append(backTags, value["name"].(string))
			}
			return backTags
		}
		return nil
	}
	return nil
}

// 获取指定tag版本
func ListTagVersion(repo, tag, remote string) Source {
	//	Gitee
	if repo == "gitee" {
		result := SendGet(MakeGiteeTags(remote) + "/" + tag)
		var dataMap map[string]interface{}
		if err := json.Unmarshal([]byte(result), &dataMap); err == nil {
			assets := make(map[string]string)
			var source Source
			if dataMap["assets"] != nil {
				for _, value := range dataMap["assets"].([]interface{}) {
					item := value.(map[string]interface{})
					if item["name"] != nil {
						assets[item["name"].(string)] = item["browser_download_url"].(string)
					}
				}
				source.Tag = dataMap["tag_name"].(string)
				source.Created = dataMap["created_at"].(string)
				source.Assets = assets
				return source
			} else {
				return Source{
					Tag:     tag,
					Created: "",
					Assets:  nil,
				}
			}
		}
		return Source{
			Tag:     "",
			Created: "",
			Assets:  nil,
		}
	}
	return Source{
		Tag:     "",
		Created: "",
		Assets:  nil,
	}
}

// 获取所有可下载版本
func ListAll(repo, remote string) []Source {
	// Gitee
	if repo == "gitee" {
		result := SendGet(MakeGiteeReleases(remote))
		var dataS []map[string]interface{}
		if err := json.Unmarshal(result, &dataS); err == nil {
			var backArray []Source
			for _, value := range dataS {
				var backMap Source
				cache := make(map[string]string)
				for _, v := range value["assets"].([]interface{}) {
					item := v.(map[string]interface{})
					if item["name"] != nil {
						cache[item["name"].(string)] = item["browser_download_url"].(string)
					}
				}
				backMap.Tag = value["tag_name"].(string)
				backMap.Created = value["created_at"].(string)
				backMap.Assets = cache
				backArray = append(backArray, backMap)
			}
			return backArray
		}
		return nil
	}
	return nil
}