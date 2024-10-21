package web_model

import (
	"encoding/json"
	"os"
	"uso.link/internal/utils/log2"
)

type WebModel struct {
	Title       string
	Canonical   string
	Keywords    string
	Description string
}

var WebConfig map[string]WebModel

func InitWebModel(conf_path string) error {
	if buf, err := os.ReadFile(conf_path); !log2.IfErrPrt(err) {

		WebConfig = make(map[string]WebModel)
		if err := json.Unmarshal(buf, &WebConfig); log2.IfErrPrt(err) {
			log2.Logf("%v", string(buf))
			return err
		}
		return nil

	} else {
		return err
	}

}
