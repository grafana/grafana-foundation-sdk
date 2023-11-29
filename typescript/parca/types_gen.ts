// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export enum ParcaQueryType {
	Metrics = "metrics",
	Profile = "profile",
	Both = "both",
}

export const defaultParcaQueryType = (): ParcaQueryType => (ParcaQueryType.Both);

export interface dataquery {
	labelSelector?: string;
	profileTypeId?: string;
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

