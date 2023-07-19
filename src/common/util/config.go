package util

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/BurntSushi/toml"
)

func NewConfig(defaultPath string, v interface{}) bool {
	parser := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	confName := parser.String("config", defaultPath, "Input config file")
	etcdEnv := parser.String("eenv", "", "Input etcd env")
	optionVal := parser.String("option", "", "Input option")
	shardVal := parser.String("shard", "", "Input server shard")
	assetVal := parser.String("asset", "", "Input asset dir")
	hostPortVal := parser.String("host", "", "Input host id and port")
	parser.Parse(os.Args[1:])

	bRet := false
	if *confName != "" {
		bRet = true
		fmt.Printf("Load config file: %s\n", *confName)
		_, err := toml.DecodeFile(*confName, v)
		if err != nil {
			fmt.Println("config file load failed:", *confName, err)
			return false
		}
	}

	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		if val.Kind() == reflect.Struct {
			vEEnv := val.FieldByName("EtcdEnv")
			if vEEnv.Kind() == reflect.String {
				if *etcdEnv != "" {
					vEEnv.SetString(*etcdEnv)
				} else if vEEnv.String() == "" {
					vEnv := val.FieldByName("Env")
					if vEnv.Kind() == reflect.String {
						vEEnv.SetString(vEnv.String())
					}
				}
			}
			if *optionVal != "" {
				vOption := val.FieldByName("Option")
				if vOption.Kind() == reflect.String {
					vOption.SetString(*optionVal)
				}
			}
			if *shardVal != "" {
				vShard := val.FieldByName("ServerShard")
				if vShard.Kind() == reflect.String {
					vShard.SetString(*shardVal)
				}
			}

			vAsset := val.FieldByName("AssetsDir")
			if vAsset.Kind() == reflect.String {
				if *assetVal != "" {
					vAsset.SetString(*assetVal)
				} else if vAsset.String() == "" {
					rootPath, _ := filepath.Abs(filepath.Dir(os.Args[0]))
					path := filepath.Join(rootPath, "pff_assets")
					if _, err := os.Stat(path); err != nil {
						path = filepath.Join(rootPath, "../pff_assets")
					}
					vAsset.SetString(path)
				}
			}

			vHost := val.FieldByName("Host")
			if vHost.Kind() == reflect.String {
				if *hostPortVal != "" {
					vHost.SetString(*hostPortVal)
				}
			}

			//if f, ok := val.Type().FieldByName("EtcdEnv"); ok {
			//	if f.Type.Kind() == reflect.String {
			//		eEnv := *etcdEnv
			//		if eEnv == "" {
			//			eEnv = val.FieldByName("Env").String()
			//		}
			//		val.FieldByIndex(f.Index).SetString(eEnv)
			//	}
			//}
		}
	}

	return bRet
}
