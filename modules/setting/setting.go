package setting

import (
	"github.com/Unknwon/log"
)

var (
	HTTPPort = 3000
)

func init() {
	log.Prefix = "[FWS]"
}
