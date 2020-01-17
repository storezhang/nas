package ds

import (
    `github.com/storezhang/nas/synology`
)

func (ds *DownloadStation) List(req *listDownloadRequest) (rsp *ListDownloadResponse, err error) {
    var callResponse ListDownloadResponse

    err = synology.CallApi(
        &callResponse,
        ds.synology,
        Session,
        synology.MethodPost,
        req,
    )
    rsp = &callResponse

    return
}
