// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export enum PyroscopeQueryType {
	Metrics = "metrics",
	Profile = "profile",
	Both = "both",
}

export const defaultPyroscopeQueryType = (): PyroscopeQueryType => (PyroscopeQueryType.Both);

export interface dataquery {
	// Specifies the query label selectors.
	labelSelector?: string;
	// Specifies the query span selectors.
	spanSelector?: string[];
	// Specifies the type of profile to query.
	profileTypeId?: string;
	// Allows to group the results.
	groupBy?: string[];
	// Sets the maximum number of nodes in the flamegraph.
	maxNodes?: number;
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	refId?: string;
	// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
	hide?: boolean;
	// Specify the query flavor
	// TODO make this required and give it a default
	queryType?: string;
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	datasource?: any;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): dataquery => ({
	labelSelector: "{}",
	_implementsDataqueryVariant: () => {},
});

