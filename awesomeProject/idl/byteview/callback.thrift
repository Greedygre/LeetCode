include "../base.thrift"

namespace go callback

//返回的Base Code，标识Callback是否正常
const i32 BaseStatusCode_SUCCESS = 0; //成功
const i32 BaseStatusCode_FAIL = 1; //callback 失败

struct HeartbeatStopRequest {
    1: required i64 UserID;
    2: required string Token;
    3: required string Extra;
    255: optional base.Base Base;
}

struct HeartbeatStopResponse {
    255: required base.BaseResp BaseResp;
}

service ByteviewMutexCallbackService {
    HeartbeatStopResponse ByteviewMutexBeatStopCallback(1: HeartbeatStopRequest request);
}
