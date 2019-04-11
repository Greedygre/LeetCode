namespace go errors

enum ErrorCode {
    // HTTP CODE
    OK = 200;
    ACCEPTED = 201;
    BAD_REQUEST = 400;
    NOT_FOUND = 404;
    UNAUTHORIZED = 401;
    FORBIDDEN = 403;
    CONFLICT = 409;
    INTERNAL_SERVER_ERROR = 500;

    // INNER CODE
    LARK_ERROR = 900;
}

exception LarkException {
    1: ErrorCode Code = ErrorCode.OK;
    2: string Message = "";
}