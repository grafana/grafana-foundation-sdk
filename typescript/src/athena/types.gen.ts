// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as dashboard from '../dashboard';


// Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts
export interface Dataquery {
	format: FormatOptions;
	connectionArgs: ConnectionArgs;
	table?: string;
	column?: string;
	queryID?: string;
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId: string;
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	rawSQL: string;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: dashboard.DataSourceRef;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): Dataquery => ({
	format: FormatOptions.TimeSeries,
	connectionArgs: defaultConnectionArgs(),
	refId: "",
	rawSQL: "",
	_implementsDataqueryVariant: () => {},
});

export const defaultKey = "__default";

export interface ConnectionArgs {
	region?: string;
	catalog?: string;
	database?: string;
	resultReuseEnabled?: boolean;
	resultReuseMaxAgeInMinutes?: number;
}

export const defaultConnectionArgs = (): ConnectionArgs => ({
	region: "__default",
	catalog: "__default",
	database: "__default",
	resultReuseEnabled: false,
	resultReuseMaxAgeInMinutes: 60,
});

export enum FormatOptions {
	TimeSeries = 0,
	Table = 1,
	Logs = 2,
}

export const defaultFormatOptions = (): FormatOptions => (FormatOptions.TimeSeries);

