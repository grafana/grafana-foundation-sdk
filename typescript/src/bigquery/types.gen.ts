// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as dashboard from '../dashboard';


// Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75
export interface Dataquery {
	dataset?: string;
	table?: string;
	project?: string;
	format: QueryFormat;
	rawQuery?: boolean;
	rawSql: string;
	location?: string;
	partitioned?: boolean;
	partitionedField?: string;
	convertToUTC?: boolean;
	sharded?: boolean;
	queryPriority?: QueryPriority;
	timeShift?: string;
	editorMode?: EditorMode;
	sql?: SQLExpression;
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): Dataquery => ({
	format: QueryFormat.Timeseries,
	rawSql: "",
	refId: "",
	_implementsDataqueryVariant: () => {},
});

export enum QueryFormat {
	Timeseries = 0,
	Table = 1,
}

export const defaultQueryFormat = (): QueryFormat => (QueryFormat.Timeseries);

export enum QueryPriority {
	Interactive = "INTERACTIVE",
	Batch = "BATCH",
}

export const defaultQueryPriority = (): QueryPriority => (QueryPriority.Interactive);

export enum EditorMode {
	Code = "code",
	Builder = "builder",
}

export const defaultEditorMode = (): EditorMode => (EditorMode.Code);

export interface SQLExpression {
	columns?: QueryEditorFunctionExpression[];
	from?: string;
	// whereJsonTree?: _
	whereString?: string;
	groupBy?: QueryEditorGroupByExpression[];
	orderBy?: QueryEditorPropertyExpression;
	orderByDirection?: OrderByDirection;
	limit?: number;
	offset?: number;
}

export const defaultSQLExpression = (): SQLExpression => ({
});

export interface QueryEditorFunctionExpression {
	type: "function";
	name?: string;
	parameters?: QueryEditorFunctionParameterExpression[];
}

export const defaultQueryEditorFunctionExpression = (): QueryEditorFunctionExpression => ({
	type: "function",
});

export interface QueryEditorFunctionParameterExpression {
	type: "functionParameter";
	name?: string;
}

export const defaultQueryEditorFunctionParameterExpression = (): QueryEditorFunctionParameterExpression => ({
	type: "functionParameter",
});

export interface QueryEditorGroupByExpression {
	type: "groupBy";
	property: QueryEditorProperty;
}

export const defaultQueryEditorGroupByExpression = (): QueryEditorGroupByExpression => ({
	type: "groupBy",
	property: defaultQueryEditorProperty(),
});

export interface QueryEditorProperty {
	type: QueryEditorPropertyType;
	name?: string;
}

export const defaultQueryEditorProperty = (): QueryEditorProperty => ({
	type: QueryEditorPropertyType.String,
});

export enum QueryEditorPropertyType {
	String = "string",
}

export const defaultQueryEditorPropertyType = (): QueryEditorPropertyType => (QueryEditorPropertyType.String);

export interface QueryEditorPropertyExpression {
	type: "property";
	property: QueryEditorProperty;
}

export const defaultQueryEditorPropertyExpression = (): QueryEditorPropertyExpression => ({
	type: "property",
	property: defaultQueryEditorProperty(),
});

export enum OrderByDirection {
	ASC = "ASC",
	DESC = "DESC",
}

export const defaultOrderByDirection = (): OrderByDirection => (OrderByDirection.ASC);

export enum QueryEditorExpressionType {
	Property = "property",
	Operator = "operator",
	Or = "or",
	And = "and",
	GroupBy = "groupBy",
	Function = "function",
	FunctionParameter = "functionParameter",
}

export const defaultQueryEditorExpressionType = (): QueryEditorExpressionType => (QueryEditorExpressionType.Property);

