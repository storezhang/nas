package ds

import (
    `nas/synology`
)

// DownloadInfoRequest 请求下载信息
type DownloadInfoRequest struct {
    synology.BaseRequest

    Ids        []string
    Additional []string
}

// NewDownloadInfoRequest 创建下载请求
func NewDownloadInfoRequest(ids []string, additional []string) *DownloadInfoRequest {
    return &DownloadInfoRequest{
        BaseRequest: synology.NewBaseRequest("SYNO.DownloadStation2.Task", "get", 2),
        Ids:         ids,
        Additional:  additional,
    }
}

// DownloadInfo 下载信息
type DownloadInfo struct {
    Id         string
    Status     int
    Title      string
    Type       string
    Additional Additional
}

// Additional 附加信息
type Additional struct {
    Trackers []Tracker `json:"tracker"`
}

// Tracker Tracker信息
type Tracker struct {
    Peers       int
    Seeds       int
    status      string
    UpdateTimer int `json:"update_timer"`
    Url         string
}

// AddTrackersRequest 添加BT Tracker请求
type AddTrackersRequest struct {
    synology.BaseRequest

    TaskId   string   `url:"task_id"`
    Trackers []string `url:"tracker"`
}

func NewAddTrackersRequest(taskId string, trackers []string) *AddTrackersRequest {
    return &AddTrackersRequest{
        BaseRequest: synology.NewBaseRequest("SYNO.DownloadStation2.Task.BT.Tracker", "add", 2),
        TaskId:      taskId,
        Trackers:    trackers,
    }
}
