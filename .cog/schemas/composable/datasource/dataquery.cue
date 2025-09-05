package datasource

import (
	"github.com/grafana/grafana/packages/grafana-schema/src/common"
)

Dataquery: {
	common.DataQuery

	// Panel ID from wich the queries will be reused.
	panelId: uint32
	
	withTransforms: bool
}
