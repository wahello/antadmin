package db

import (
	"github.com/antbiz/antadmin/db/upms"
	"github.com/gogf/gf/frame/g"
)

var (
	SystemClientMaster system.Client
	SystemClientSalve  system.Client
)

// InitClients 初始化全部 ent client
func InitClients() error {
	masterDrv, err := GetMasterDriver()
	if err != nil {
		return err
	}
	SystemClientMaster = system.NewClient(masterDrv)

	slaveDrv, err := GetSlaveDriver()
	if err != nil {
		g.Log().Errorf("init slave db failed: %v", err)
		SystemClientSalve = SystemClientMaster
	} else {
		SystemClientSalve = system.NewClient(slaveDrv)
	}

	return nil
}
