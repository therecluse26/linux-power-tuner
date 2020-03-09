package builtin

import "github.com/therecluse26/uranium/pkg/monitors/net_monitors"

var ExportedFuncs = map[string]interface{}{
	"Hello": Hello,
	"Notify": Notify,
	"Ping": net_monitors.Ping,
}