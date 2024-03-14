// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export interface Options {
	keepTime: boolean;
	includeVars: boolean;
	showStarred: boolean;
	showRecentlyViewed: boolean;
	showSearch: boolean;
	showHeadings: boolean;
	maxItems: number;
	query: string;
	folderId?: number;
	tags: string[];
}

export const defaultOptions = (): Options => ({
	keepTime: false,
	includeVars: false,
	showStarred: true,
	showRecentlyViewed: false,
	showSearch: false,
	showHeadings: true,
	maxItems: 10,
	query: "",
	tags: [],
});

