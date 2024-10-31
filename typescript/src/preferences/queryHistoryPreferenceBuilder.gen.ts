// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as preferences from '../preferences';

export class QueryHistoryPreferenceBuilder implements cog.Builder<preferences.QueryHistoryPreference> {
    protected readonly internal: preferences.QueryHistoryPreference;

    constructor() {
        this.internal = preferences.defaultQueryHistoryPreference();
    }

    /**
     * Builds the object.
     */
    build(): preferences.QueryHistoryPreference {
        return this.internal;
    }

    // one of: '' | 'query' | 'starred';
    homeTab(homeTab: string): this {
        this.internal.homeTab = homeTab;
        return this;
    }
}
