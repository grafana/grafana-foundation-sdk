// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as annotationslist from '../annotationslist';

export class OptionsBuilder implements cog.Builder<annotationslist.Options> {
    protected readonly internal: annotationslist.Options;

    constructor() {
        this.internal = annotationslist.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): annotationslist.Options {
        return this.internal;
    }

    onlyFromThisDashboard(onlyFromThisDashboard: boolean): this {
        this.internal.onlyFromThisDashboard = onlyFromThisDashboard;
        return this;
    }

    onlyInTimeRange(onlyInTimeRange: boolean): this {
        this.internal.onlyInTimeRange = onlyInTimeRange;
        return this;
    }

    tags(tags: string[]): this {
        this.internal.tags = tags;
        return this;
    }

    limit(limit: number): this {
        this.internal.limit = limit;
        return this;
    }

    showUser(showUser: boolean): this {
        this.internal.showUser = showUser;
        return this;
    }

    showTime(showTime: boolean): this {
        this.internal.showTime = showTime;
        return this;
    }

    showTags(showTags: boolean): this {
        this.internal.showTags = showTags;
        return this;
    }

    navigateToPanel(navigateToPanel: boolean): this {
        this.internal.navigateToPanel = navigateToPanel;
        return this;
    }

    navigateBefore(navigateBefore: string): this {
        this.internal.navigateBefore = navigateBefore;
        return this;
    }

    navigateAfter(navigateAfter: string): this {
        this.internal.navigateAfter = navigateAfter;
        return this;
    }
}

