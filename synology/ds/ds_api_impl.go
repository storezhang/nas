package ds

import (
    `github.com/storezhang/nas/synology`
)

func (ds *DownloadStation) Delete(req *deleteRequest) (rsp *synology.BaseResponse, err error) {
    var callResponse synology.BaseResponse
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
