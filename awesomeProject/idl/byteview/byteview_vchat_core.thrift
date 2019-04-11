include "../base.thrift"
include "./gateway.thrift"
include "./callback.thrift"
include "./push.thrift"

namespace go byteview.vchat.core

const i32 STATUS_SUCCESS = 0 //命令执行成功
const i32 STATUS_CMD_INVALID = 100 //命令不支持
const i32 STATUS_PARAM_INVALID = 101 //参数错误
const i32 STATUS_INTERNAL_FAIL = 102 //内部执行错误

enum AdminCmd {
    QUERY = 1;
    CLEAN = 2;
    CHECK = 3;
}

struct AdminRequest {
    1: required AdminCmd Cmd;
    2: optional string Params;
    255: base.Base Base;
}

struct AdminResponse {
     1: optional string Result;
     255: base.BaseResp BaseResp;
}

struct Call {
    1: required i64 VideoChatID;
    2: required i64 HostUserID;
    3: required i64 HostDeviceID;
    4: required i64 CalleeUserID;
    5: required i64 CalleeDeviceID;
    6: required string VendorType;
    7: optional i64 StartTime;
    8: optional i64 EndTime;
}

struct GetVideoChatCallRequest {
    1: required i64 VideoChatID;
    255: optional base.Base Base;
}

struct GetVideoChatCallResponse {
    1: required Call Call;
    255: required base.BaseResp BaseResp;
}

struct Participant {
    1: required i64 UserID;
    2: required i64 DeviceID;
    3: optional i64 JoinTime;
    4: optional i64 LeaveTime;
}

struct GroupMeeting {
    1: required i64 MeetingID;
    2: list<Participant> Participants;
    3: optional i64 StartTime;
    4: optional i64 EndTime;
    5: required i64 HostID;
}

struct GetGroupMeetingRequest {
    1: required i64 GroupMeetingID;
    255: optional base.Base Base;
}

struct GetGroupMeetingResponse {
    1: required GroupMeeting GroupMeeting;
    255: required base.BaseResp BaseResp;
}

service VideoChatService {
    gateway.GatewayResponse Serve(1: gateway.GatewayRequest request);
    AdminResponse Manage(1: AdminRequest request);
    callback.HeartbeatStopResponse ByteviewMutexBeatStopCallback(1: callback.HeartbeatStopRequest request);
    push.SendPushResultsResponse SendPushResults(1: push.SendPushResultsRequest request);
    GetVideoChatCallResponse GetVideoChatCall(1: GetVideoChatCallRequest request);
    GetGroupMeetingResponse GetGroupMeeting(1: GetGroupMeetingRequest request);
}