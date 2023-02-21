namespace go relation
/**
    "id": 0,
    "name": "string",
    "follow_count": 0,
    "follower_count": 0,
    "is_follow": true
**/
struct User {
    1:  required i64        id;
    2:  required string     name;               // 用户昵称
    3:  optional i64        follow_count;       // 关注数
    4:  optional i64        follower_count;     // 粉丝数
    5:  required bool       is_follow;          // 是否关注 true-已关注 false-未关注
    6:  optional string     avatar;             // 头像
    7:  optional string     background_image;   // 个人顶部大图
    8:  optional string     signature;          // 个人简介
    9:  optional i64        total_favorited;    // 获赞数
    10: optional i64        work_count;         // 作品数
    11: optional i64        favorite_count;     // 喜欢数
}

struct FollowInfo {
    1: required i64 follow_count;    // 关注数
    2: required i64 follower_count;  // 粉丝数
    3: required bool is_follow;      // 是否关注
}

// 好友
struct FriendInfo {
    1:  required i64        id;
    2:  required string     name;               // 昵称
    3:  optional i64        follow_count;       // 关注数
    4:  optional i64        follower_count;     // 粉丝数
    5:  required bool       is_follow;          // 是否关注
    6:  optional string     avatar;             // 用户头像url
    7:  optional string     background_image;   // 用户个人页顶部大图url
    8:  optional string     signature;          // 个人简介
    9:  optional i64        total_favorited;    // 获赞数量
    10: optional i64        work_count;         // 作品数
    11: optional i64        favorite_count;     // 喜欢数
}

// 关注
struct FollowReq {
    // 发请求的userId
    1: required i64 my_uid;
    // 关注目标userId
    2: required i64 target_uid;
    // 1-关注，2-取消关注
    3: required i32 action_type;
}

struct FollowResp {
    1: required i32 status_code;
    2: optional string msg;
}

struct GetFollowInfoReq {
    1: optional i64 my_uid;     // 发请求的userId
    2: required i64 target_uid; // 目标userId
}

struct GetFollowInfoResp {
    1: required i32 status_code;
    2: optional FollowInfo follow_info; // 用户的关注信息
    3: optional string msg;
}



// 获取关注列表
struct GetFollowListReq {
    1: optional i64 myId;     // 发出请求的userId
    2: required i64 targetId; // 查询目标userId
}


struct GetFollowListResp {
    1: required i32 statusCode;
    // 关注的用户列表
    2: required list<User> userList;
    3: optional string msg;
}

// 获取粉丝列表
struct GetFollowerListReq {
    1: optional i64 myId;       // 发出请求的userId
    2: required i64 targetId;   // 查询目标userId
}

struct GetFollowerListResp {
    1: required i32 statusCode;
    2: required list<User> user_list; // 粉丝用户列表
    3: optional string msg;
}

struct GetFriendListReq {
    1: optional i64 myUid;      // 发起请求的userId
    2: required i64 targetUid;  // 目标userId
}

struct GetFriendListResp {
    1: required i32 status_code;
    2: optional list<User> friend_list; // 好友列表
    3: optional string msg;
}

struct GetIsFollowsReq{
    1: required i64         myUid;      // 发出请求的userId
    2: required list<i64>   userIds;    // 查询目标用户集合
}

struct GetIsFollowsResp{
    1: required i32                 status_code;
    2: optional map<i64, bool>      is_follow_map;    // userId - isFollow
    3: optional string              msg;
}

service RelationService {
    // 关注
    FollowResp FollowMethod(1: FollowReq request)
    // 用户关注信息 关注数 粉丝数 是否关注
    GetFollowInfoResp GetFollowInfoMethod(1: GetFollowInfoReq request)
    // 获取关注列表
    GetFollowListResp GetFollowListMethod(1: GetFollowListReq request)
    // 获取粉丝列表
    GetFollowerListResp GetFollowerListMethod(1: GetFollowerListReq request)
    // 获取好友列表
    GetFriendListResp GetFriendListMethod(1: GetFriendListReq request)
    // 根据userIds 获取 各用户关注状态
    GetIsFollowsResp GetIsFollowsMethod(1: GetIsFollowsReq request)
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
