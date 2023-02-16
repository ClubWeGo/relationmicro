namespace go relation
/**
    "id": 0,
    "name": "string",
    "follow_count": 0,
    "follower_count": 0,
    "is_follow": true
**/
struct User {
    // 用户ID
    1: required i64 id;
    // 用户昵称
    2: required string name;
    // 用户关注数
    3: required i64 follow_count;
    // 用户粉丝数
    4: required i64 follower_count;
    // 是否关注 true-已关注 false-未关注
    5: required bool is_follow;
}

// 关注
struct ActionReq {
    // 关注对方用户id
    1: required string to_user_id;
    // 1-关注，2-取消关注
    2: required string action_type;
}

struct ActionResp {
    // 状态码：0-成功、其他-失败
    1: required i64 status_code;
    // 状态信息
    2: required string status_msg;

}

// 查询用户 关注数 和 粉丝数
struct GetFollowAndFollowerReq {
    // 查询的用户id
    1: required i64 user_id;
    2: optional i64 me_id;
}

struct GetFollowAndFollowerResp {
    // 当前用户是否关注
    1: required bool is_follow;
    // 关注数
    2: required i64 follow_count;
    // 粉丝数
    3: required i64 follower_count;
}

// 获取关注列表
struct GetFollowListReq {
    // 用户id
    1: required i64 user_id;
}


struct GetFollowListResp {
    // 关注的用户列表
    1: required list<User> user_list;
}

// 获取粉丝列表
struct GetFollowerListReq {
    1: required i64 user_id; // 用户id
}

struct GetFollowerListResp {
    1: required list<User> user_list; // 用户列表
}

service RelationService {
    // 关注
    ActionResp ActionMethod(1: ActionReq request)
    // 获取关注数、粉丝数、是否关注
    GetFollowAndFollowerResp GetFollowAndFollowerMethod(1: GetFollowAndFollowerReq request)
    // 获取关注列表
    GetFollowListResp GetFollowListReqMethod(1: GetFollowListReq request)
    // 获取粉丝列表
    GetFollowerListResp GetFollowerListMethod(1: GetFollowerListReq request)
}

// message
struct Message {
    // 消息id
    1: required i64 id;
    // 该消息接收者的id
    2: required i64 to_user_id;
     // 该消息发送者的id
    3: required i64 from_user_id;
     // 消息内容
    4: required string content;
    // 消息创建时间
    5: optional string create_time;
}

struct GetAllMessageReq {
    // from user_id
    1: required i64 user_id;
    2: required i64 to_user_id;
}

struct GetAllMessageResp {
    1: required bool status;
    2: required list<Message> msg;
}

struct SendMessageReq {
    // from user id
    1: required i64 user_id;
    2: required i64 to_user_id;
    3: required string content;
}

struct SendMessageResp {
    1: required bool status;
}

service MessageService {
    GetAllMessageResp GetAllMessageMethod (1: GetAllMessageReq request);
    SendMessageResp SendMessageMethod (1: SendMessageReq request);
}
