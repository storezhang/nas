package ds

import (
    `github.com/storezhang/nas/synology`
)

const Session string = "DownloadStation2"

// DownloadStation 封装
type DownloadStation struct {
    synology *synology.Synology
}

// NewDS 创建新的DS
func NewDS(synology *synology.Synology) (ds *DownloadStation) {
    return &DownloadStation{
        synology: synology,
    }
}

type deleteRequest struct {
    synology.BaseRequest

    Status      []int
    Type        []string
    TypeInverse bool `url:"type_inverse"`
}

// NewDeleteCompletedRequest 清除已完成的任务
func NewDeleteCompletedRequest() *deleteRequest {
    return &deleteRequest{
        BaseRequest: synology.NewBaseRequest("SYNO.DownloadStation2.Task", "delete_condition", 2),
        Status:      []int{5},
        Type:        []string{"emule"},
        TypeInverse: true,
    }
}
