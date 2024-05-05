package jwt

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gvb_server/config"
	"gvb_server/global"
	"log"
	"os"
	"testing"
)

func TestParseToken(t *testing.T) {
	func() {
		c := &config.Config{}
		yamlConf, err := os.ReadFile("../../settings.yaml")
		if err != nil {
			panic(fmt.Errorf("get yamlConf error : %s", err))
		}
		err = yaml.Unmarshal(yamlConf, c)
		if err != nil {
			log.Fatalf("conf init Unmarshal: %v", err)
		}
		log.Println("config yamlFile load init success")
		global.Config = c
	}()
	type args struct {
		tokenStr string
	}
	test := struct {
		name string
		args args
	}{
		name: "ParseToken_CASE",
		args: args{
			tokenStr: TOKEN,
		},
	}

	t.Run(test.name, func(t *testing.T) {

		res, err := ParseToken(test.args.tokenStr)

		if err != nil {
			t.Error("invalid token")
			return
		}
		t.Log("res: ", res)
		t.Log("ParseToken_CASE SUCCESS")
	})

}
