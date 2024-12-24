package redisServer

import (
	"gvb_server/global"
	"strconv"
)

/* 文章浏览量 */
const lookPrefix = "look"

// Look 浏览某一篇文章
func Look(id string) error {
	num, _ := global.Redis.HGet(lookPrefix, id).Int()
	num++
	err := global.Redis.HSet(lookPrefix, id, num).Err()
	return err
}

// GetLook 获取某一篇文章下的浏览数
func GetLook(id string) int {
	num, _ := global.Redis.HGet(lookPrefix, id).Int()
	return num
}

// GetLookInfo 取出浏览量数据
func GetLookInfo() map[string]int {
	var lookInfo = make(map[string]int)
	maps := global.Redis.HGetAll(lookPrefix).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		lookInfo[id] = num
	}
	return lookInfo
}

func LookClear() {
	global.Redis.Del(lookPrefix)
}
