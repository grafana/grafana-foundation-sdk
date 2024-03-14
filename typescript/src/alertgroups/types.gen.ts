// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export interface Options {
	// Comma-separated list of values used to filter alert results
	labels: string;
	// Name of the alertmanager used as a source for alerts
	alertmanager: string;
	// Expand all alert groups by default
	expandAll: boolean;
}

export const defaultOptions = (): Options => ({
	labels: "",
	alertmanager: "",
	expandAll: false,
});

