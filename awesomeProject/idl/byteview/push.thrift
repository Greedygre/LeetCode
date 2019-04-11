include "../base.thrift"

namespace go push

enum PushErrorCode {
    ACKED               = 0 // 收到了客户端的 ack
    NOT_ONLINE          = 1 // 客户端 not online
    SUCCESS_BUT_NOT_ACK = 2 // 调用 frontier 返回 success, 但是多次重试没收到客户端 ack
}

struct PushRes {
    1: i64 UserID;
    2: i64 DeviceID;
    3: PushErrorCode Code;
}

struct SendPushResultsRequest {
    1: optional i64 SequenceID;
    2: list<PushRes> Results;
    255: base.Base Base;

}

struct SendPushResultsResponse {
    255: base.BaseResp BaseResp;
}

service PushService {
    SendPushResultsResponse SendPushResults(1: SendPushResultsRequest req);
}
