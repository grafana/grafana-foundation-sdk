// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export interface Options {
	keepTime: boolean;
	includeVars: boolean;
	showStarred: boolean;
	showRecentlyViewed: boolean;
	showSearch: boolean;
	showHeadings: boolean;
	showFolderNames: boolean;
	maxItems: number;
	query: string;
	tags: string[];
	// folderId is deprecated, and migrated to folderUid on panel init
	folderId?: number;
	folderUID?: string;
}

export const defaultOptions = (): Options => ({
	keepTime: false,
	includeVars: false,
	showStarred: true,
	showRecentlyViewed: false,
	showSearch: false,
	showHeadings: true,
	showFolderNames: true,
	maxItems: 10,
	query: "",
	tags: [],
});

