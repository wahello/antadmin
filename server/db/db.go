package db

import (
	"github.com/antbiz/antadmin/db/smc"
	"github.com/antbiz/antadmin/db/umc"
)

var (
	UmcClientMaster umc.Client
	UmcClientSalve  umc.Client

	SmcClientMaster smc.Client
	SmcClientSalve  smc.Client
)

// InitClients 初始化全部 ent client
func InitClients() (err error) {
	if UmcClientMaster, UmcClientSalve, err = umc.InitClient(); err != nil {
		return
	}

	if SmcClientMaster, SmcClientSalve, err = smc.InitClient(); err != nil {
		return
	}

	return
}
