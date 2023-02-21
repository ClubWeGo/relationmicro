package util

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
)

const (
	FOLLOW_ADDRESS     = "/util/lua/follow"
	UNFOLLOWER_ADDRESS = "/util/lua/unfollow"
)

type LuaScripts struct {
	Follow    string
	UnFollow  string
	IsFollows string // 多个用户的关注状态
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
	isFollows := `
local isFollows = {}
for i=1, #ARGV
	do
	local isFollow = redis.call('zrank', KEYS[1], ARGV[i])
	if type(isFollow) == 'number'	then isFollows[i] = 1
	else isFollows[i] = 0
	end
end
return isFollows
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
	luaScripts.IsFollows = isFollows
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

func GetIsFollowsScript() string {
	return luaScripts.IsFollows
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

/*
*
获取脚本的Sha1
*/
func GetLuaSha1(scriptStr string) (string, error) {
	o := sha1.New()
	_, err := o.Write([]byte(scriptStr))
	if err != nil {
		//log.Errorf("GetLuaSha1 encrypt error, scriptStr:%s", scriptStr)
		return "", fmt.Errorf("get lua sha1 exception:%s", err)
	}
	return hex.EncodeToString(o.Sum(nil)), nil
}
