package errno

var (
    // 由于各种原因导致的Bad Request
    PayloadBadRequest = &Error{ErrorCode: "00001", Message: "Bad Request: lack of payload."}
    ParamBadRequest   = &Error{ErrorCode: "00002", Message: "Bad Request: lack of parameters."}
    
    // User类型相关的错误
    UserNotFound      = &Error{ErrorCode: "10001", Message: "DB: User Not Found!"}
    UserAleadyExisted = &Error{ErrorCode: "10002", Message: "Register: User already existed!"}
    PasswordInvalid   = &Error{ErrorCode: "10003", Message: "Login: Password Invalid!"}
)
