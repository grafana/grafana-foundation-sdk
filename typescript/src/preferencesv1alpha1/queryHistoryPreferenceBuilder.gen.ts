// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as preferencesv1alpha1 from '../preferencesv1alpha1';

export class QueryHistoryPreferenceBuilder implements cog.Builder<preferencesv1alpha1.QueryHistoryPreference> {
    protected readonly internal: preferencesv1alpha1.QueryHistoryPreference;

    constructor() {
        this.internal = preferencesv1alpha1.defaultQueryHistoryPreference();
    }

    /**
     * Builds the object.
     */
    build(): preferencesv1alpha1.QueryHistoryPreference {
        return this.internal;
    }

    // one of: '' | 'query' | 'starred';
    homeTab(homeTab: string): this {
        this.internal.homeTab = homeTab;
        return this;
    }
}

