package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func init() {
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
}

func main() {
	viper.SetConfigFile("./viper-demo/set-default-demo/config.yml")
	viper.WriteConfig()
	for k, v := range viper.GetStringMap("Taxonomies") {
		fmt.Println(k, " => ", v)
	}

}
