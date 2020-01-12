package ds

import (
    `songjiang/synology`
)

const Session string = "DownloadStation2"

type downloadStation struct {
    synology *synology.Synology
}

func NewDS(synology *synology.Synology) (ds *downloadStation) {
    return &downloadStation{
        synology: synology,
    }
}
