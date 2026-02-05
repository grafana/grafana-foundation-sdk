import * as common from '../common';
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
    refId?: string;
    hide?: boolean;
    queryType?: string;
    datasource?: common.DataSourceRef;
    _implementsDataqueryVariant(): void;
}
export declare const defaultDataquery: () => Dataquery;
export declare enum QueryFormat {
    Timeseries = 0,
    Table = 1
}
export declare const defaultQueryFormat: () => QueryFormat;
export declare enum QueryPriority {
    Interactive = "INTERACTIVE",
    Batch = "BATCH"
}
export declare const defaultQueryPriority: () => QueryPriority;
export declare enum EditorMode {
    Code = "code",
    Builder = "builder"
}
export declare const defaultEditorMode: () => EditorMode;
export interface SQLExpression {
    columns?: QueryEditorFunctionExpression[];
    from?: string;
    whereString?: string;
    groupBy?: QueryEditorGroupByExpression[];
    orderBy?: QueryEditorPropertyExpression;
    orderByDirection?: OrderByDirection;
    limit?: number;
    offset?: number;
}
export declare const defaultSQLExpression: () => SQLExpression;
export interface QueryEditorFunctionExpression {
    type: QueryEditorExpressionType.Function;
    name?: string;
    parameters?: QueryEditorFunctionParameterExpression[];
}
export declare const defaultQueryEditorFunctionExpression: () => QueryEditorFunctionExpression;
export declare enum QueryEditorExpressionType {
    Property = "property",
    Operator = "operator",
    Or = "or",
    And = "and",
    GroupBy = "groupBy",
    Function = "function",
    FunctionParameter = "functionParameter"
}
export declare const defaultQueryEditorExpressionType: () => QueryEditorExpressionType;
export interface QueryEditorFunctionParameterExpression {
    type: QueryEditorExpressionType.FunctionParameter;
    name?: string;
}
export declare const defaultQueryEditorFunctionParameterExpression: () => QueryEditorFunctionParameterExpression;
export interface QueryEditorGroupByExpression {
    type: QueryEditorExpressionType.GroupBy;
    property: QueryEditorProperty;
}
export declare const defaultQueryEditorGroupByExpression: () => QueryEditorGroupByExpression;
export interface QueryEditorProperty {
    type: QueryEditorPropertyType.String;
    name?: string;
}
export declare const defaultQueryEditorProperty: () => QueryEditorProperty;
export declare enum QueryEditorPropertyType {
    String = "string"
}
export declare const defaultQueryEditorPropertyType: () => QueryEditorPropertyType;
export interface QueryEditorPropertyExpression {
    type: QueryEditorExpressionType.Property;
    property: QueryEditorProperty;
}
export declare const defaultQueryEditorPropertyExpression: () => QueryEditorPropertyExpression;
export declare enum OrderByDirection {
    ASC = "ASC",
    DESC = "DESC"
}
export declare const defaultOrderByDirection: () => OrderByDirection;
