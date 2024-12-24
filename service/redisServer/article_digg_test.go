package redisServer

import (
	"github.com/sirupsen/logrus"
	"gvb_server/core"
	"gvb_server/global"
	"testing"
)

func init() {
	core.InitCore("../../settings.yaml")
	global.Redis = core.InitRedis()
	global.ESClient = core.InitElasticSearch()
}

func TestDigg(t *testing.T) {
	type args struct {
		id string
	}
	// 测试用例1：正常点赞操作，期望无错误
	testCase1 := struct {
		name    string
		args    args
		wantErr bool
	}{
		name: "正常点赞，无错误",
		args: args{
			id: "test",
		},
		wantErr: false,
	}
	// 测试用例2：传入空的id进行点赞，期望有错误（根据实际业务逻辑，空id可能不合理）
	testCase2 := struct {
		name    string
		args    args
		wantErr bool
	}{
		name: "传入空id点赞，应返回错误",
		args: args{
			id: "",
		},
		wantErr: true,
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		testCase1,
		testCase2,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Digg(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Digg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetDigg(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want int
	}{
		{
			name: "test_get_digg",
			id:   "test_article_id",
			want: 10, // 假设期望从Redis中获取到的点赞数为10，可根据实际情况调整
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 向Redis中设置测试数据（模拟已有点赞数存在的情况）
			err := global.Redis.HSet(diggPrefix, tt.id, tt.want).Err()
			if err != nil {
				t.Fatalf("向Redis中设置测试数据失败: %v", err)
			}
			// 调用GetDigg函数获取点赞数
			actualDiggNum := GetDigg(tt.id)
			if actualDiggNum != tt.want {
				logrus.Info("与预期一致")
				logrus.Info(tt.want, actualDiggNum)
				t.Errorf("GetDigg(%s) = %d, want %d", tt.id, actualDiggNum, tt.want)
			} else {
				logrus.Info("与预期一致")
				logrus.Info("want:", tt.want, " actual:", actualDiggNum)
			}

		})
	}
}

func TestGetDiggInfo(t *testing.T) {
	diggInfo := GetDiggInfo()
	for key, info := range diggInfo {
		t.Log(key, info)
	}
}
