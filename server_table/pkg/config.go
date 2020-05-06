package pkg

import (
	"github.com/starqingzhu/chess/common"
	"github.com/starqingzhu/chess/util/log"
)

func Init(confPath string) bool {
	if err := common.InitConfig(confPath + "/table.conf"); err != nil {
		log.Error("init common config fail")
		return false
	}

	if !InitRoomConfig(confPath + "/room.csv") {
		return false
	}

	return true
}
