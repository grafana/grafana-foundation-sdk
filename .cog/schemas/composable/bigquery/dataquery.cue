package bigquery

import (
	"github.com/grafana/grafana/packages/grafana-schema/src/common"
)

// Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75
Dataquery: {
    common.DataQuery

    dataset?: string
    table?: string
    project?: string

    format: #QueryFormat
    rawQuery?: bool
    rawSql: string
    location?: string

    partitioned?: bool
    partitionedField?: string
    convertToUTC?: bool
    sharded?: bool
    queryPriority?: #QueryPriority
    timeShift?: string
    editorMode?: #EditorMode
    sql?: #SQLExpression

    #QueryFormat: 0 | 1 @cog(kind="enum", memberNames="Timeseries|Table")
    #QueryPriority: "INTERACTIVE" | "BATCH" @cog(kind="enum", memberNames="Interactive|Batch")
    #EditorMode: "code" | "builder"

    #SQLExpression: {
        columns?: [...#QueryEditorFunctionExpression]
        from?: string
        // whereJsonTree?: _
        whereString?: string
        groupBy?: [...#QueryEditorGroupByExpression]
        orderBy?: #QueryEditorPropertyExpression
        orderByDirection?: #OrderByDirection
        limit?: int
        offset?: int
    }

    #QueryEditorExpressionType: "property" | "operator" | "or" | "and" | "groupBy" | "function" | "functionParameter"

    #QueryEditorFunctionExpression: {
        type: #QueryEditorExpressionType & "function"
        name?: string
        parameters?: [...#QueryEditorFunctionParameterExpression]
    }

    #QueryEditorFunctionParameterExpression: {
        type: #QueryEditorExpressionType & "functionParameter"
        name?: string
    }

    #QueryEditorGroupByExpression: {
        type: #QueryEditorExpressionType & "groupBy"
        property: #QueryEditorProperty
    }

    #QueryEditorProperty: {
        type: #QueryEditorPropertyType
        name?: string
    }

    #QueryEditorPropertyType: "string" @cog(kind="enum")

    #QueryEditorPropertyExpression: {
        type: #QueryEditorExpressionType & "property"
        property: #QueryEditorProperty
    }

    #OrderByDirection: "ASC" | "DESC"
}
