package ds

import (
    `github.com/storezhang/nas/synology`
)

func (ds *DownloadStation) SetTrackers(req *trackersRequest) (rsp *synology.BaseResponse, err error) {
    if nil == req.Trackers || 0 == len(req.Trackers) {
        rsp = synology.NewSuccessResponse()
        return
    }

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

func (ds *DownloadStation) Set(req *setRequest) (rsp *synology.BaseResponse, err error) {
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
