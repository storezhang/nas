package synology

import (
    `fmt`
    `net/http`

    `github.com/storezhang/gos/urls`
)

type Code int

const (
    CodeNeedLogin Code = 119
)

// Synology 群晖NAS
type Synology struct {
    Url      string
    Username string
    Password string
}

// Response 响应接口
type Response interface {
    // IsSuccess 响应是否成功
    IsSuccess() bool
    // Code 响应码
    Code() Code
}

// BaseResponse 返回基类
type BaseResponse struct {
    Success bool
    Error   struct {
        Code Code
    }
}

// NewSuccessResponse 创建新的成功的响应
func NewSuccessResponse() *BaseResponse {
    return &BaseResponse{
        Success: true,
        Error:   struct{ Code Code }{Code: 0},
    }
}

func (rsp *BaseResponse) IsSuccess() bool {
    return rsp.Success
}

func (rsp *BaseResponse) Code() Code {
    return rsp.Error.Code
}

// LoginResponse 登录响应
type LoginResponse struct {
    BaseResponse

    Data struct {
        Sid string
    }
}

// BaseRequest 请求基类
type BaseRequest struct {
    Api     string
    Version int
    Method  string
}

// NewBaseRequest 创建基础请求
func NewBaseRequest(api string, method string, version int) BaseRequest {
    return BaseRequest{
        Api:     api,
        Version: version,
        Method:  method,
    }
}

// LoginRequest 登录请求
type LoginRequest struct {
    BaseRequest

    Account string
    Passwd  string
    Session string
    Format  string
}

// NewLoginRequest 创建登录请求
func NewLoginRequest(username string, password string, session string) LoginRequest {
    return LoginRequest{
        BaseRequest: NewBaseRequest("SYNO.API.Auth", "login", 2),
        Account:     username,
        Passwd:      password,
        Session:     session,
        Format:      "sid",
    }
}

// NewDSLoginRequest 创建新的DownloadStation登录请求
func NewDSLoginRequest(username string, password string) LoginRequest {
    return NewLoginRequest(username, password, "DownloadStation2")
}

// Call 统一请求，增加重试机制等
type MethodCall string

const (
    MethodGet  MethodCall = "GET"
    MethodPost MethodCall = "POST"
)

func Call(
    synology *Synology,
    session string,
    url string,
    method MethodCall,
    body interface{},
) (rsp Response, err error) {
    var callRsp Response

    _, _, callErr := httpClient.Clone().CustomMethod(string(method), fmt.Sprintf("%s/%s", synology.Url, url)).
        Send(urls.QueryString(body)).
        EndStruct(callRsp)
    if nil != callErr {
        err = callErr[0]
    } else if !callRsp.IsSuccess() && CodeNeedLogin == callRsp.Code() { // 需要登录
        var loginRsp LoginResponse

        _, _, loginErr := httpClient.Get(fmt.Sprintf("%s/webapi/auth.cgi", synology.Url)).
            Query(NewLoginRequest(synology.Username, synology.Password, session)).
            EndStruct(&loginRsp)
        if nil != loginErr {
            err = loginErr[0]
            return
        } else {
            httpClient.AddCookie(&http.Cookie{
                Name:  "id",
                Value: loginRsp.Data.Sid,
            })
            rsp, err = Call(synology, session, url, method, body)
        }
    } else { // 调用成功，返回
        rsp = callRsp
    }

    return
}

// CallApi Api调用
func CallApi(
    synology *Synology,
    session string,
    method MethodCall,
    body interface{},
) (rsp Response, err error) {
    return Call(synology, session, "webapi/entry.cgi", method, body)
}
