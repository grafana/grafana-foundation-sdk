// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as preferencesv1alpha1 from '../preferencesv1alpha1';

export class CookiePreferencesBuilder implements cog.Builder<preferencesv1alpha1.CookiePreferences> {
    protected readonly internal: preferencesv1alpha1.CookiePreferences;

    constructor() {
        this.internal = preferencesv1alpha1.defaultCookiePreferences();
    }

    /**
     * Builds the object.
     */
    build(): preferencesv1alpha1.CookiePreferences {
        return this.internal;
    }

    analytics(analytics: any): this {
        this.internal.analytics = analytics;
        return this;
    }

    performance(performance: any): this {
        this.internal.performance = performance;
        return this;
    }

    functional(functional: any): this {
        this.internal.functional = functional;
        return this;
    }
}

