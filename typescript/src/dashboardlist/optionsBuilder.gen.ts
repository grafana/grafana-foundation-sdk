// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardlist from '../dashboardlist';

export class OptionsBuilder implements cog.Builder<dashboardlist.Options> {
    protected readonly internal: dashboardlist.Options;

    constructor() {
        this.internal = dashboardlist.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): dashboardlist.Options {
        return this.internal;
    }

    keepTime(keepTime: boolean): this {
        this.internal.keepTime = keepTime;
        return this;
    }

    includeVars(includeVars: boolean): this {
        this.internal.includeVars = includeVars;
        return this;
    }

    showStarred(showStarred: boolean): this {
        this.internal.showStarred = showStarred;
        return this;
    }

    showRecentlyViewed(showRecentlyViewed: boolean): this {
        this.internal.showRecentlyViewed = showRecentlyViewed;
        return this;
    }

    showSearch(showSearch: boolean): this {
        this.internal.showSearch = showSearch;
        return this;
    }

    showHeadings(showHeadings: boolean): this {
        this.internal.showHeadings = showHeadings;
        return this;
    }

    maxItems(maxItems: number): this {
        this.internal.maxItems = maxItems;
        return this;
    }

    query(query: string): this {
        this.internal.query = query;
        return this;
    }

    tags(tags: string[]): this {
        this.internal.tags = tags;
        return this;
    }

    // folderId is deprecated, and migrated to folderUid on panel init
    folderId(folderId: number): this {
        this.internal.folderId = folderId;
        return this;
    }

    folderUID(folderUID: string): this {
        this.internal.folderUID = folderUID;
        return this;
    }
}

