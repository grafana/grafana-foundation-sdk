// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export interface Options {
	onlyFromThisDashboard: boolean;
	onlyInTimeRange: boolean;
	tags: string[];
	limit: number;
	showUser: boolean;
	showTime: boolean;
	showTags: boolean;
	navigateToPanel: boolean;
	navigateBefore: string;
	navigateAfter: string;
}

export const defaultOptions = (): Options => ({
	onlyFromThisDashboard: false,
	onlyInTimeRange: false,
	tags: [],
	limit: 10,
	showUser: true,
	showTime: true,
	showTags: true,
	navigateToPanel: true,
	navigateBefore: "10m",
	navigateAfter: "10m",
});

