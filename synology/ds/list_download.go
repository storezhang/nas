package ds

import (
    `github.com/storezhang/nas/synology`
)

// listDownloadRequest 获得下载列表
type listDownloadRequest struct {
    synology.BaseRequest

    SortBy     string `url:"sort_by"`
    Order      string
    Action     string
    Limit      int
    Type       []string
    Additional []string
    Status     []int
}

func NewListAllDownloadRequest(
    sortBy string,
    order string,
    action string,
    limit int,
    typ []string,
    additional []string,
    status []int,
) *listDownloadRequest {
    return &listDownloadRequest{
        BaseRequest: synology.NewBaseRequest("SYNO.DownloadStation2.Task", "list", 2),
        SortBy:      sortBy,
        Order:       order,
        Action:      "getall",
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
