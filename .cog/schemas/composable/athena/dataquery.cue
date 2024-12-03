package athena

import (
	"github.com/grafana/grafana/packages/grafana-schema/src/common"
)

// Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts
Dataquery: {
  common.DataQuery

	format: #FormatOptions
	connectionArgs: #ConnectionArgs
	table?: string
	column?: string

	queryID?: string

	rawSQL: string | *""
}

defaultKey: "__default"

#ConnectionArgs: {
	region?: string | *defaultKey
	catalog?: string | *defaultKey
	database?: string | *defaultKey
	resultReuseEnabled?: bool | *false
	resultReuseMaxAgeInMinutes?: number | *60
}

#FormatOptions: 0 | 1 | 2 @cog(kind="enum", memberNames="TimeSeries|Table|Logs")
