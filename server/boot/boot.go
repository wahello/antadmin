package boot

import (
	"github.com/antbiz/antadmin/db"
	"github.com/gogf/gf/frame/g"
)

func init() {
	if err := db.InitClients(); err != nil {
		g.Log().Fatalf("[boot] init db clients failed: %v", err)
	}
}
