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
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
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
	type: QueryEditorExpressionType.Function;
	name?: string;
	parameters?: QueryEditorFunctionParameterExpression[];
}

export const defaultQueryEditorFunctionExpression = (): QueryEditorFunctionExpression => ({
	type: QueryEditorExpressionType.Function,
});

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

export interface QueryEditorFunctionParameterExpression {
	type: QueryEditorExpressionType.FunctionParameter;
	name?: string;
}

export const defaultQueryEditorFunctionParameterExpression = (): QueryEditorFunctionParameterExpression => ({
	type: QueryEditorExpressionType.FunctionParameter,
});

export interface QueryEditorGroupByExpression {
	type: QueryEditorExpressionType.GroupBy;
	property: QueryEditorProperty;
}

export const defaultQueryEditorGroupByExpression = (): QueryEditorGroupByExpression => ({
	type: QueryEditorExpressionType.GroupBy,
	property: defaultQueryEditorProperty(),
});

export interface QueryEditorProperty {
	type: QueryEditorPropertyType.String;
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
	type: QueryEditorExpressionType.Property;
	property: QueryEditorProperty;
}

export const defaultQueryEditorPropertyExpression = (): QueryEditorPropertyExpression => ({
	type: QueryEditorExpressionType.Property,
	property: defaultQueryEditorProperty(),
});

export enum OrderByDirection {
	ASC = "ASC",
	DESC = "DESC",
}

export const defaultOrderByDirection = (): OrderByDirection => (OrderByDirection.ASC);

