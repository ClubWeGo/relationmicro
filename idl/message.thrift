namespace go message

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