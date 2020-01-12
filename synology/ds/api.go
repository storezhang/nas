package ds

import (
    `fmt`

    `github.com/parnurzeal/gorequest`
    `github.com/storezhang/gos/urls`

    `nas/synology`
)

// DownloadApi 下载接口
type DownloadApi interface {
    // List 列出所有的下载任务
    List(
        sortBy string,
        order string,
        action string,
        limit int,
        typ []string,
        additional []string,
        status []int,
    ) (rsp *ListDownloadResponse, err error)

    // AddTrackers 给任务添加Tracker
    AddTrackers(taskId string, trackers []string) (rsp *synology.BaseResponse, err error)
}

func (ds *DownloadStation) List(
    sortBy string,
    order string,
    limit int,
    typ []string,
    additional []string,
    status []int,
) (data *ListDownloadResponse, err error) {
    if callRsp, callErr := synology.Call(
        ds.synology,
        Session,
        func(httpClient *gorequest.SuperAgent) (callRsp synology.Response, callErr error) {
            var listDownloadRsp ListDownloadResponse

            _, _, listErr := httpClient.Post(fmt.Sprintf("%s/webapi/entry.cgi", ds.synology.Url)).
                Send(urls.QueryString(NewListDownloadRequest(sortBy, order, "getall", limit, typ, additional, status))).
                EndStruct(&listDownloadRsp)
            if nil != listErr {
                callErr = listErr[0]
            } else {
                callRsp = &listDownloadRsp
            }

            return
        },
    ); nil != callErr {
        err = callErr
    } else {
        data = callRsp.(*ListDownloadResponse)
    }

    return
}

func (ds *DownloadStation) AddTrackers(taskId string, trackers []string) (data *synology.BaseResponse, err error) {
    if callRsp, callErr := synology.Call(
        ds.synology,
        Session,
        func(httpClient *gorequest.SuperAgent) (callRsp synology.Response, callErr error) {
            var addTrackerRsp synology.BaseResponse

            _, _, listErr := httpClient.Post(fmt.Sprintf("%s/webapi/entry.cgi", ds.synology.Url)).
                Send(urls.QueryString(NewAddTrackersRequest(taskId, trackers))).
                EndStruct(&addTrackerRsp)
            if nil != listErr {
                callErr = listErr[0]
            } else {
                callRsp = &addTrackerRsp
            }

            return
        },
    ); nil != callErr {
        err = callErr
    } else {
        data = callRsp.(*synology.BaseResponse)
    }

    return
}
