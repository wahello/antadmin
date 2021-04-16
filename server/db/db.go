package db

import (
	"github.com/antbiz/antadmin/db/system"
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
	slaveDrv, err := GetSlaveDriver()
	if err != nil {
		g.Log().Errorf("init slave db failed: %v", err)
		slaveDrv = masterDrv
	}

	SystemClientMaster = system.NewClient(masterDrv)
	SystemClientSalve = system.NewClient(slaveDrv)

	return nil
}
