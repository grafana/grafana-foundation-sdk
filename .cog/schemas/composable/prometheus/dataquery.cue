// This schema was deleted from grafana/grafana in https://github.com/grafana/grafana/pull/83808
// Until a proper alternative is published, we cheat slightly by copy-pasting it here.
package grafanaplugin

import (
	"github.com/grafana/kindsys"
	"github.com/grafana/grafana/packages/grafana-schema/src/common"
)

kindsys.Composable & {
	maturity:        "experimental"
	name:            "Prometheus" + "DataQuery"
	schemaInterface: "DataQuery"
	lineage: {
		schemas: [{
			version: [0, 0]
			schema: {
				common.DataQuery

				// The actual expression/query that will be evaluated by Prometheus
				expr: string
				// Returns only the latest value that Prometheus has scraped for the requested time series
				instant?: bool
				// Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
				range?: bool
				// Execute an additional query to identify interesting raw samples relevant for the given expr
				exemplar?: bool
				// Specifies which editor is being used to prepare the query. It can be "code" or "builder"
				editorMode?: #QueryEditorMode
				// Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
				format?: #PromQueryFormat
				// Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
				legendFormat?: string
				// @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
				// See https://github.com/grafana/grafana/issues/48081
				intervalFactor?:  number
				#QueryEditorMode: "code" | "builder"                  @cuetsy(kind="enum")
				#PromQueryFormat: "time_series" | "table" | "heatmap" @cuetsy(kind="type")
			}
		}]
		name: "Prometheus" + "DataQuery"
		lenses: []
	}
}
