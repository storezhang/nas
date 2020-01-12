package ds

import (
    `nas/synology`
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
