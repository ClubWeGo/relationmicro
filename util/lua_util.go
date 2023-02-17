package util

import (
	"os"
)

const (
	FOLLOW_ADDRESS   = "/util/lua/follow"
	UNFOLLOWER_ADDRESS = "/util/lua/unfollow"
)

type LuaScripts struct {
	Follow   string
	UnFollow string
}

var luaScripts LuaScripts

/*
*
初始化加载各脚本
*/
func InitLoadLua() {
	// 关注
	followScript := `
redis.call('zadd', KEYS[1], ARGV[1], ARGV[2]);
redis.call('zadd', KEYS[2], ARGV[1], ARGV[3]);
return 1;
` // 取关
	unFollowScript := `
redis.call('zrem', KEYS[1], ARGV[1]);
redis.call('zrem', KEYS[2], ARGV[2]);
return 1;
`


	// todo golang test run build相对路径不一样 目前没找到通用方法
	//rootPath, err := GetRootPath()
	//dirs, err := os.ReadDir(rootPath)
	//for _, dir := range dirs {
	//	fmt.Println(dir.Name())
	//}
	//followScript, err := ReadAll(rootPath + FOLLOW_ADDRESS)
	//if err != nil {
	//	log.Panicf("follow lua script file loading err:%s", err)
	//}
	//followerScript, err := ReadAll(rootPath + FOLLOWER_ADDRESS)
	//if err != nil {
	//	log.Panicf("follower lua script file loading err:%s", err)
	//}
	luaScripts.Follow = followScript
	luaScripts.UnFollow = unFollowScript
}

/*
*
follow lua script
*/
func GetFollowScript() string {
	return luaScripts.Follow
}

/*
*
unFollow lua script
*/
func GetUnFollowScript() string {
	return luaScripts.UnFollow
}

func ReadAll(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func GetRootPath() (string, error) {
	//file, _ := exec.LookPath(os.Args[0])
	//path, _ := filepath.Abs(file)
	//index := strings.LastIndex(path, string(os.PathSeparator))
	//return path[:index]
	return os.Getwd()
}
