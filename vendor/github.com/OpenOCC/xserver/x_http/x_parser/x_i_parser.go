package x_parser

import (
	"github.com/OpenOCC/xserver/x_http/x_req"
)

type IParser interface {
	Parse() (xReq *x_req.XReq)
}
