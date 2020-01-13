package ds

import (
    `github.com/storezhang/nas/synology`
)

// ListDownloadRequest 获得下载列表
type ListDownloadRequest struct {
    synology.BaseRequest

    SortBy     string `url:"sort_by"`
    Order      string
    Action     string
    Limit      int
    Type       []string
    Additional []string
    Status     []int
}

func NewListDownloadRequest(
    sortBy string,
    order string,
    action string,
    limit int,
    typ []string,
    additional []string,
    status []int,
) *ListDownloadRequest {
    return &ListDownloadRequest{
        BaseRequest: synology.NewBaseRequest("SYNO.DownloadStation2.Task", "list", 2),
        SortBy:      sortBy,
        Order:       order,
        Action:      action,
        Limit:       limit,
        Type:        typ,
        Additional:  additional,
        Status:      status,
    }
}

// ListDownloadResponse 下载列表响应
type ListDownloadResponse struct {
    synology.BaseResponse

    Data struct {
        Offset int
        Tasks  []DownloadInfo `json:"task"`
        Total  int
    }
}
