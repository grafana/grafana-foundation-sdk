// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export enum PhlareQueryType {
	Metrics = "metrics",
	Profile = "profile",
	Both = "both",
}

export const defaultPhlareQueryType = (): PhlareQueryType => (PhlareQueryType.Both);

export interface dataquery {
	labelSelector?: string;
	profileTypeId?: string;
	groupBy?: string[];
	maxNodes?: number;
	refId?: string;
	hide?: boolean;
	queryType?: string;
	datasource?: any;
	_implementsDataqueryVariant(): void;
}

export const defaultDataquery = (): dataquery => ({
	labelSelector: "{}",
	_implementsDataqueryVariant: () => {},
});

