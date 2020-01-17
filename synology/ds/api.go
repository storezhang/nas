package ds

import (
    `github.com/storezhang/nas/synology`
)

// DownloadApi 下载接口
type DownloadApi interface {
    // List 列出所有的下载任务
    List(req *listDownloadRequest) (rsp *ListDownloadResponse, err error)

    // SetTrackers 给任务设置Tracker
    SetTrackers(req *trackersRequest) (rsp *synology.BaseResponse, err error)

    // Set 设置任务
    Set(req *setRequest) (rsp *synology.BaseResponse, err error)

    // Delete 删除
    Delete(req *deleteRequest) (rsp *synology.BaseResponse, err error)
}
