package db

import (
	"github.com/antbiz/antadmin/db/umc"
	"github.com/gogf/gf/frame/g"
)

var (
	UmcClientMaster umc.Client
	UmcClientSalve  umc.Client
)

// InitClients 初始化全部 ent client
func InitClients() error {
	masterDrv, err := GetMasterDriver()
	if err != nil {
		return err
	}
	UmcClientMaster = umc.NewClient(masterDrv)

	slaveDrv, err := GetSlaveDriver()
	if err != nil {
		g.Log().Errorf("init slave db failed: %v", err)
		UmcClientSalve = UmcClientMaster
	} else {
		UmcClientSalve = umc.NewClient(slaveDrv)
	}

	return nil
}
