package ds

import (
    `strings`

    `github.com/storezhang/nas/synology`
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

func (info *DownloadInfo) GetTrackers() (trackers []string) {
    for _, t := range info.Additional.Trackers {
        trackers = append(trackers, t.Url)
    }

    return
}

func (info *DownloadInfo) GetFailedTrackers() (trackers []string) {
    for _, t := range info.Additional.Trackers {
        if "success" != strings.ToLower(t.Status) {
            trackers = append(trackers, t.Url)
        }
    }

    return
}

// Additional 附加信息
type Additional struct {
    Trackers []Tracker `json:"tracker"`
}

// Tracker Tracker信息
type Tracker struct {
    Peers       int
    Seeds       int
    Status      string
    UpdateTimer int `json:"update_timer"`
    Url         string
}

type trackersRequest struct {
    synology.BaseRequest

    TaskId   string   `url:"task_id"`
    Trackers []string `url:"tracker"`
}

// NewTrackersAddRequest 创建增加Tracker列表请求
func NewTrackersAddRequest(taskId string, trackers []string) *trackersRequest {
    return &trackersRequest{
        BaseRequest: synology.NewBaseRequest("SYNO.DownloadStation2.Task.BT.Tracker", "add", 2),
        TaskId:      taskId,
        Trackers:    trackers,
    }
}

// NewTrackersDeleteRequest 创建删除Tracker列表请求
func NewTrackersDeleteRequest(taskId string, trackers []string) *trackersRequest {
    return &trackersRequest{
        BaseRequest: synology.NewBaseRequest("SYNO.DownloadStation2.Task.BT.Tracker", "delete", 2),
        TaskId:      taskId,
        Trackers:    trackers,
    }
}
