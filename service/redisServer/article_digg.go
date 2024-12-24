package redisServer

import (
	"errors"
	"gvb_server/global"
	"strconv"
)

/*点赞*/
// diggPrefix 点赞前缀
const diggPrefix = "digg"

// Digg 文章点赞处理
func Digg(id string) error {
	if id == "" {
		return errors.New("文章id为空")
	}
	num, _ := global.Redis.HGet(diggPrefix, id).Int()
	num++
	err := global.Redis.HSet(diggPrefix, id, num).Err()
	return err
}

// GetDigg 获取文章的点赞数
func GetDigg(id string) int {
	num, _ := global.Redis.HGet(diggPrefix, id).Int()
	return num
}

// GetDiggInfo 取出所有文章的点赞数据
func GetDiggInfo() map[string]int {
	diggInfo := make(map[string]int)
	maps := global.Redis.HGetAll(diggPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		diggInfo[id] = num
	}
	return diggInfo
}

// DiggClear 删除索引
func DiggClear() {
	global.Redis.Del(diggPrefix)
}
