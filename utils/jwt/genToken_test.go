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

var TOKEN string

func TestGenToken(t *testing.T) {
	// 读取配置文件
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
		user JwtPayLoad
	}
	test := struct {
		name string
		args args
	}{

		name: "GenToken_CASE",
		args: args{
			user: JwtPayLoad{
				Username: "test_user",
				NickName: "test_nick",
				Role:     1,
				UserID:   1,
			},
		},
	}

	t.Run(test.name, func(t *testing.T) {
		token, err := GenToken(test.args.user)
		TOKEN = token
		t.Log("token: ", token)
		t.Log("GenToken_CASE SUCCESS")
		if err != nil {
			t.Errorf("GenToken() error = %v", err)
			return
		}
	})

}
