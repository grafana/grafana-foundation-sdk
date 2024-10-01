// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as dashboard from '../dashboard';


export enum PhlareQueryType {
	Metrics = "metrics",
	Profile = "profile",
	Both = "both",
}

export const defaultPhlareQueryType = (): PhlareQueryType => (PhlareQueryType.Both);

export interface dataquery {
	// Specifies the query label selectors.
	labelSelector: string;
	// Specifies the type of profile to query.
	profileTypeId: string;
	// Allows to group the results.
	groupBy: string[];
	// Sets the maximum number of nodes in the flamegraph.
	maxNodes?: number;
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

export const defaultDataquery = (): dataquery => ({
	labelSelector: "{}",
	profileTypeId: "",
	groupBy: [],
	refId: "",
	_implementsDataqueryVariant: () => {},
});

