package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/viper"
)

var (
	mapIds              []int          // 地图ID列表
	jsonRootPath        string         // json 文件所在的根目录
	jsonFileNamePrefix  string         // json 文件的前缀
	jsonFilePaths       []string       // json 文件路径集合
	allReplaceRules     []*replaceRule // 全部替换规则列表
	jsonRewriteFilePath string         // json 重写文件地址
	multiple_value      float64        // 倍数值
)

func init() {
	viper.SetDefault("map_ids", []int{10001, 10002, 10003, 10004, 10007, 10008, 10009, 20001, 20002, 20003, 20004, 20005, 20006, 20007, 20008, 21004, 22004, 30001, 30002, 30003, 30004, 30005, 30006, 30008, 30009, 30010, 30011, 30012, 30013, 30014, 30015, 30016, 40000, 41001, 41002, 41003, 42000, 42001, 42002, 42003, 42004, 42005, 42006, 42007, 42008, 42009, 42010, 42100, 50001, 50002, 50003, 50011, 50012, 50013, 50021, 50022, 50023, 50031, 50032, 50033, 51000, 51001, 51010, 51020, 51021, 52000, 52001, 52002, 53001, 53002, 53003, 53004, 53005, 53006, 53007, 53008, 53009, 53010, 53011, 53012, 53013, 53014, 53015, 60001, 70001, 70002, 70003, 70004, 70005, 70006, 70007, 70008, 70009, 70010, 70011, 70012, 70013, 70014, 71000, 71001, 71002, 71003, 71004, 71005})
	viper.SetDefault("json_root_path", "./deal-json-file/json/")
	viper.SetDefault("json_file_name_prefix", "map_")
	viper.SetDefault("config_path", "./deal-json-file/conf/config.toml")
	viper.SetDefault("json_rewrite_file_path", "./deal-json-file/rejson")
	viper.SetDefault("multiple_value", 100000) // 倍数值
	viper.SetConfigFile("./deal-json-file/conf/config.toml")
	// 用于写入默认配置
	viper.WriteConfigAs("./deal-json-file/conf/config_defalut.toml")

	// 读取配置
	viper.ReadInConfig()
	mapIds = viper.GetIntSlice("map_ids")
	jsonRootPath = viper.GetString("json_root_path")
	jsonFileNamePrefix = viper.GetString("json_file_name_prefix")
	jsonRewriteFilePath = viper.GetString("json_rewrite_file_path")
	multiple_value = viper.GetFloat64("multiple_value")
	remkdir()
}

func main() {
	// 获取所有json文件的路径集合
	getAllJsonFilePath()

	// 读取所有json文件的内容
	readAllJsonFile()

}

// 替换规则
type replaceRule struct {
	Old   *old    `json:"erlangUnitInfo"` // 要被替换的内容的描述信息
	X     float64 `json:"x"`              // 转换到的 X
	Y     float64 `json:"y"`              // 转换到的 Y
	Z     float64 `json:"z"`              // 转换到的 Z
	Sacle float64 `json:"sacle"`          // 转换到的 Sacle
	Rx    float64 `json:"r_x"`            // 转换到的 Rx
	Ry    float64 `json:"r_y"`            // 转换到的 Ry
	Rz    float64 `json:"r_z"`            // 转换到的 Rz
	ToX   float64 `json:"to_x"`           // 转换到的 ToX
	ToY   float64 `json:"to_y"`           // 转换到的 ToY
	ToZ   float64 `json:"to_z"`           // 转换到的 ToZ
	Xp    int32   `json:"x_p"`            // 处理过精度后的 X
	Yp    int32   `json:"y_p"`            // 处理过精度后的 Y
	Zp    int32   `json:"z_p"`            // 处理过精度后的 Z
	ToXp  int32   `json:"to_x_p"`         // 处理过精度后的 ToX
	ToYp  int32   `json:"to_y_p"`         // 处理过精度后的 ToY
	ToZp  int32   `json:"to_z_p"`         // 处理过精度后的 ToZ
}

// npc 结构
type old struct {
	BaseID     int    `json:"base_id"`      // npc 配置ID
	X          int    `json:"x"`            // x坐标
	Y          int    `json:"y"`            // y坐标
	MapID      int    `json:"map_id"`       // 地图配置ID
	ToX        int    `json:"to_x"`         // 去往X坐标
	ToY        int    `json:"to_y"`         // 去往Y坐标
	ToMapID    int    `json:"to_map_id"`    // 去往地图ID
	File       string `json:"file"`         // 源码所在文件
	Line       int    `json:"line"`         // 所在行
	Raw        string `json:"raw"`          // 源码中的原始字符串
	IsAutoBase bool   `json:"is_auto_base"` // 是否为程序自动查找到的模型ID
}

func readAllJsonFile() {
	for i, v := range jsonFilePaths {

		if data, err := readAll(v); err != nil {
			if os.IsNotExist(err) {

			} else {
				panic(err)
			}
		} else {
			replaceRules := make([]*replaceRule, 0)
			if err := json.Unmarshal(data, &replaceRules); err != nil {
				panic(err)
			} else {
				dealPrecision(replaceRules)
				allReplaceRules = append(allReplaceRules, replaceRules...)
				mapId := mapIds[i]
				rewriteJson(&replaceRules, mapId)
			}
		}
	}
}

func getAllJsonFilePath() {
	for _, v := range mapIds {
		filepath := fmt.Sprintf(jsonRootPath+jsonFileNamePrefix+"%d.json", v)
		jsonFilePaths = append(jsonFilePaths, filepath)
	}
}

// 读取文件所有内容并返回文件内容
func readAll(filepath string) ([]byte, error) {
	if f, err := os.Open(filepath); err != nil {
		return nil, err
	} else {
		defer f.Close()
		return ioutil.ReadAll(f)
	}
}

func rewriteJson(replaceRule *[]*replaceRule, mapId int) {
	writeFile(replaceRule, jsonRewriteFilePath, mapId)
}

func writeFile(replaceRule *[]*replaceRule, filepath string, id int) {
	f, err := os.OpenFile(fmt.Sprintf("%s/map_%d.json", filepath, id), os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	j, err := json.MarshalIndent(replaceRule, "", "\t")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(j))
	if err != nil {
		panic(err)
	}

}

func remkdir() {
	if _, err := os.Stat(jsonRewriteFilePath); err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(jsonRewriteFilePath, 0777)
		} else {
			panic(err)
		}
	} else {
		os.RemoveAll(jsonRewriteFilePath)
		os.Mkdir(jsonRewriteFilePath, 0777)
	}
}

// 处理精度值 将 float64 转换为 int32
func dealPrecision(replaceRules []*replaceRule) {
	for _, v := range replaceRules {
		v.ToXp = about(v.ToX)
		v.ToXp = about(v.ToY)
		v.ToXp = about(v.ToZ)
		v.Xp = about(v.X)
		v.Yp = about(v.Y)
		v.Zp = about(v.Z)
	}
}

func about(f float64) int32 {
	return int32(f * multiple_value)
}
