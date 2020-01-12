package synology

import (
    `fmt`
    `net/http`

    `github.com/parnurzeal/gorequest`
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

func (synology *Synology) UpdateToBest(trackers []string) (err error) {
    return
}

type Response interface {
    IsSuccess() bool
    Code() Code
}

// BaseResponse 返回基类
type BaseResponse struct {
    Success bool
    Error   struct {
        Code Code
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
func Call(
    synology *Synology,
    session string,
    method func(httpClient *gorequest.SuperAgent) (callRsp Response, callErr error),
) (rsp Response, err error) {
    if callRsp, callErr := method(httpClient.Clone()); nil != callErr {
        err = callErr
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
            rsp, err = Call(synology, session, method)
        }
    } else { // 调用成功，返回
        rsp = callRsp
    }

    return
}
