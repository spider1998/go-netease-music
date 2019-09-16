package api

import (
	"encoding/json"
	"music/api"
	"music/log"
	"reflect"
	"sort"
	"strconv"
)

var (
	redisDSN string = "192.168.35.233:6379"
	redisLog log.Logger
)

func init() {
	var err error
	redisLog, err = log.New(true, "cloud")
	if err != nil {
		panic(err)
	}

}

func writeApiCache(url string, req map[string]interface{}, response interface{}) {
	if redis, err := api.GetRedis(redisDSN, redisLog); err != nil {
		panic(err)
	} else {
		//两分钟缓存
		err = redis.Cmd("SET", generateKey(req, url), response, "EX", "120").Err
		if err != nil {
			panic(err)
		}
	}
}

func readApiCache(req map[string]interface{}, url string) []byte {
	if redis, err := api.GetRedis(redisDSN, redisLog); err != nil {
		panic(err)
	} else {
		res, _ := redis.Cmd("GET", generateKey(req, url)).Bytes()
		return res
	}
}

func generateKey(req map[string]interface{}, url string) []byte {
	req["url"] = url
	var slic []string
	for k, re := range req {
		if reflect.TypeOf(re).String() == "int" {
			slic = append(slic, k+":"+strconv.Itoa(re.(int))+"&")
		} else if reflect.TypeOf(re).String() == "string" {
			slic = append(slic, k+":"+re.(string)+"&")
		}
	}
	sort.Strings(slic)
	key, err := json.Marshal(slic)
	if err != nil {
		panic(err)
	}
	return key
}
