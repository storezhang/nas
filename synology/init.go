package synology

import (
    `crypto/tls`
    `net/http`
    `time`

    `github.com/parnurzeal/gorequest`
)

var httpClient *gorequest.SuperAgent

func init() {
    // 初始化Http客户端
    httpClient = gorequest.New()
    httpClient.Timeout(60 * time.Second)
    // 忽略TLS证书
    httpClient.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
    httpClient.Retry(3, 5*time.Second, http.StatusBadRequest, http.StatusInternalServerError)
}
