package api

import (
	"encoding/json"
	"fmt"
	"github.com/OpenOCC/OCC/conf"
	"github.com/OpenOCC/OCC/core/types"
	"github.com/OpenOCC/OCC/node"
	"github.com/OpenOCC/OCC/param"
	"github.com/OpenOCC/OCC/util"

	"github.com/OpenOCC/xserver/x_err"
	"github.com/OpenOCC/xserver/x_http/x_req"
	"github.com/OpenOCC/xserver/x_http/x_resp"
	"github.com/OpenOCC/xserver/x_http/x_router"
)

func init() {
	x_router.All("/peer/api/ping", ping)
	x_router.Post("/peer/api/peers", delegatePeers)
	x_router.Post("/peer/api/heartbeat", heartbeat)
}

func delegatePeers(req *x_req.XReq) (*x_resp.XRespContainer, *x_err.XErr) {
	peers := param.MainChainDelegateNode
	return x_resp.Success(peers), x_err.NewXErr(nil)
}

func heartbeat(req *x_req.XReq) (*x_resp.XRespContainer, *x_err.XErr) {
	var heartbeat types.Heartbeat
	err := json.Unmarshal(req.Body, &heartbeat)
	if err != nil {
		return x_resp.Return(nil, err)
	}
	node.GetInst().Heartbeat(heartbeat)
	return x_resp.Return(nil, nil)
}

func ping(req *x_req.XReq) (*x_resp.XRespContainer, *x_err.XErr) {
	resp := &x_resp.XRespContainer{
		HttpCode: 200,
		Body:     []byte("pong"),
	}
	return resp, nil
}

func broadcast(req *x_req.XReq) (*x_resp.XRespContainer, *x_err.XErr) {
	if len(req.Query) == 0 {
		for _, peer := range param.MainChainDelegateNode {
			if !peer.Equal(conf.EKTConfig.Node) {
				url := fmt.Sprintf(`http://%s:%d%s?broadcast=true`, peer.Address, peer.Port, req.Path)
				util.HttpPost(url, req.Body)
			}
		}
	}
	return nil, nil
}
