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
    folderId?: number;
    folderUID?: string;
}
export declare const defaultOptions: () => Options;
