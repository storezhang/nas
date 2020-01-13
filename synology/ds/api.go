package ds

import (
    `github.com/storezhang/nas/synology`
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

    // DeleteTrackers 删除Tracker
    DeleteTrackers(taskId string, trackers []string) (rsp *synology.BaseResponse, err error)
}

func (ds *DownloadStation) List(
    sortBy string,
    order string,
    limit int,
    typ []string,
    additional []string,
    status []int,
) (data *ListDownloadResponse, err error) {
    var callRsp synology.Response
    callRsp, err = synology.CallApi(
        ds.synology,
        Session,
        synology.MethodPost,
        NewListDownloadRequest(sortBy, order, "getall", limit, typ, additional, status),
    )
    data = callRsp.(*ListDownloadResponse)

    return
}

func (ds *DownloadStation) AddTrackers(taskId string, trackers []string) (data *synology.BaseResponse, err error) {
    if nil == trackers || 0 == len(trackers) {
        data = synology.NewSuccessResponse()
        return
    }

    var callRsp synology.Response
    callRsp, err = synology.CallApi(
        ds.synology,
        Session,
        synology.MethodPost,
        NewTrackersAddRequest(taskId, trackers),
    )
    data = callRsp.(*synology.BaseResponse)

    return
}

func (ds *DownloadStation) RemoveTrackers(taskId string, trackers []string) (data *synology.BaseResponse, err error) {
    if nil == trackers || 0 == len(trackers) {
        data = synology.NewSuccessResponse()
        return
    }

    var callRsp synology.Response
    callRsp, err = synology.CallApi(
        ds.synology,
        Session,
        synology.MethodPost,
        NewTrackersDeleteRequest(taskId, trackers),
    )
    data = callRsp.(*synology.BaseResponse)

    return
}
