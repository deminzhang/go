package utilX

//配置读取https://blog.csdn.net/wade3015/article/details/83351776

import (
	"log"

	"github.com/BurntSushi/toml"
)

//yaml
func ReadYaml(pathFile string, out interface{}) {
	//TODO yaml
}

//toml
//go get github.com/BurntSushi/toml
func ReadToml(pathFile string, out interface{}) {
	if _, err := toml.DecodeFile(pathFile, out); err != nil {
		log.Fatal(err)
	}
}

var confs = make(map[string]interface{})

func GetCfg(name string) {

	cfg := confs[name]
	if cfg == nil {
		//load
	}

}
