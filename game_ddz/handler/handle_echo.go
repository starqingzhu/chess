package handler

import (
	"github.com/starqingzhu/chess/common"
	"github.com/starqingzhu/chess/game/server"
)

func HandleEcho(userid uint32, connid uint32, msgBody []byte) {
	server.SendResp(userid, MsgidEchoResp, common.ResultSuccess, msgBody)
}
