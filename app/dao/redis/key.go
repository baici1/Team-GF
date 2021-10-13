package redis

// redis key尽量使用命名空间方式，方便查询和拆分

const (
	KeyPrefix               = "teamGf:"
	KeyCreateTeamListPrefix = "create:team:" //不同比赛创建队伍 参数是队伍编号id
	KeyUserInTeamsPrefix    = "user:team:"   //用户的队伍表 参数是用户id
)

//为key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
