package redisServer

import (
	"gvb_server/global"
	"strconv"
)

const (
	articleLookPrefix         = "article_look"
	articleCommentCountPrefix = "article_comment_count"
	articleDiggPrefix         = "article_digg"
	commentDiggPrefix         = "comment_digg"
)

type CountDB struct {
	Index string
}

// Set 设置某一个数据,add
func (c CountDB) Set(id string) error {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	num++
	err := global.Redis.HSet(c.Index, id, num).Err()
	return err
}
func (c CountDB) SetCount(id string, num int) error {
	oldNum, _ := global.Redis.HGet(c.Index, id).Int()
	newNum := oldNum + num
	err := global.Redis.HSet(c.Index, id, newNum).Err()
	return err
}

// Get 获取某个的数据
func (c CountDB) Get(id string) int {
	num, _ := global.Redis.HGet(c.Index, id).Int()
	return num
}

// GetInfo 取出数据
func (c CountDB) GetInfo() map[string]int {
	var DiggInfo = map[string]int{}
	maps := global.Redis.HGetAll(c.Index).Val()
	for id, val := range maps {
		num, _ := strconv.Atoi(val)
		DiggInfo[id] = num
	}
	return DiggInfo
}

func (c CountDB) Clear() {
	global.Redis.Del(c.Index)
}

func NewDigg() CountDB {
	return CountDB{
		Index: articleDiggPrefix,
	}
}
func NewArticleLook() CountDB {
	return CountDB{
		Index: articleLookPrefix,
	}
}
func NewCommentCount() CountDB {
	return CountDB{
		Index: articleCommentCountPrefix,
	}
}
func NewCommentDigg() CountDB {
	return CountDB{
		Index: commentDiggPrefix,
	}
}
