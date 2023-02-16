package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	redisUtil "github.com/ClubWeGo/relationmicro/util"
)

type Message struct {
	Id        int64
	Create_at time.Time // 从redis默认设置的消息ID中的时间戳生成
	UserId    int64
	ToUserId  int64
	Content   string
}

// 根据用户id生成点对点的房间号：房间号：较小id_较大id
func GenerateP2PRoomID(user_id, to_user_id int64) string {
	// 用户查接时仅传入自己的id和目标id，需要规则来唯一确定房间号
	var max, min int64
	if user_id > to_user_id {
		max = user_id
		min = to_user_id
	} else {
		max = to_user_id
		min = user_id
	}
	room := fmt.Sprintf("%d_%d", min, max)
	return room
}

// TODO : 群聊房间号最好使用uuid生成，用户端存储房间号
func GenerateGroupRoomID() string {
	return "TODO : group room name"
}

// 发送消息
func SendP2PMsg(userId, toUserId int64, msg string) (string, error) {
	roomId := GenerateP2PRoomID(userId, toUserId)
	userIdString := strconv.FormatInt(userId, 10)
	toUserIdString := strconv.FormatInt(toUserId, 10)
	r, err := redisUtil.XADD(roomId, "*", userIdString, toUserIdString, msg, -1)
	println(r)
	if err != nil {
		return r, err
	}
	return r, nil
}

// 获取聊天室消息
func GetAllP2PMsg(userId, toUserId int64) ([]Message, error) {
	roomId := GenerateP2PRoomID(userId, toUserId)
	r, err := redisUtil.XREVRANGE(roomId, "+", "-")
	if err != nil {
		return nil, err
	}
	msg, err := ConvertReplyToMsg(r)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func convertMsgIdToTime(msgId string) time.Time {
	timeString := strings.Split(msgId, "-")[0]
	timeInt64, _ := strconv.ParseInt(timeString, 10, 64) // 13位的时间戳
	time := time.Unix(timeInt64/1e3, 0)                  // 使用10位的时间来计算
	return time
}

func ConvertReplyToMsg(reply []interface{}) (msgs []Message, err error) {
	msgs = make([]Message, len(reply))
	for index, replyItem := range reply {
		/*  消息形式
		1)  1) "1676537230806-0"
			2)  1) "from_to"
				2) "1_2"
				3) "content"
				4) "testmsg 1"
		*/
		var keyInfo = replyItem.([]interface{})
		var msgId = string(keyInfo[0].([]byte))
		msgs[index].Id, _ = strconv.ParseInt(msgId, 10, 64) // 正常不会出错，这里忽略这个错误
		msgs[index].Create_at = convertMsgIdToTime(msgId)

		var MsgList = keyInfo[1].([]interface{})
		// if len(MsgList) != 4 {
		// 	// 消息格式错误
		// } // 这里如何处理格式错误的消息？

		// 解析用户id
		userIdString := string(MsgList[1].([]byte))
		userIdSlice := strings.Split(userIdString, "_")
		fromUserIdString, toUserIdString := userIdSlice[0], userIdSlice[1]
		userId, _ := strconv.ParseInt(fromUserIdString, 10, 64) // 正常不会出错，这里忽略这个错误
		msgs[index].UserId = userId
		toUserId, _ := strconv.ParseInt(toUserIdString, 10, 64) // 正常不会出错，这里忽略这个错误
		msgs[index].ToUserId = toUserId

		// 解析消息
		msgString := string(MsgList[3].([]byte))
		msgs[index].Content = msgString

	}
	return msgs, nil
}
