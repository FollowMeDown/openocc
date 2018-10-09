package x_http

import (
	"github.com/OpenOCC/xserver/x_http/x_processer"
	"net/http"
)

func Service(w http.ResponseWriter, r *http.Request) {
	x_processer.Process(w, r)
}
