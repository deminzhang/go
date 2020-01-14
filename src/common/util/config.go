package Util

var confs = make(map[string]interface{})

func GetCfg(name string) {

	cfg := confs[name]
	if cfg == nil {
		//load
	}

}
