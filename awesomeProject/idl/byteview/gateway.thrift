include "../base.thrift"
include "./errors.thrift"

namespace go gateway

typedef errors.ErrorCode ErrorCode

struct GatewayRequest {
    1: required i64 UserId;
    2: required i64 DeviceId;
    3: required i32 Version;
    4: required string RequestId;
    5: required binary Payload;
    6: map<string,string> Context;
    255: base.Base Base;
}

struct GatewayResponse {
     1: required ErrorCode StatusCode = 200;
     2: required binary Payload;
     255: base.BaseResp BaseResp;
}