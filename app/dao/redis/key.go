package redis

// redis key尽量使用命名空间方式，方便查询和拆分

const (
	KeyPrefix               = "teamGf:"
	KeyCreateTeamListPrefix = "create:team:" //不同比赛创建队伍 参数是队伍编号id
	//KeyGameSetPrefix        = "game:"        //保存不同比赛创建的队伍
)

//为key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}

//func getRedisKeyAndData(key string, data interface{}, flag bool) string {
//	var Str string
//	if value, ok := data.(int); ok {
//		Str = strconv.Itoa(value)
//	}
//	if value, ok := data.(int64); ok {
//		Str = strconv.FormatInt(value, 10)
//	}
//	if flag {
//		return key + Str + ":"
//	}
//	return key + Str
//}
