// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as preferences from '../preferences';

export class CookiePreferencesBuilder implements cog.Builder<preferences.CookiePreferences> {
    protected readonly internal: preferences.CookiePreferences;

    constructor() {
        this.internal = preferences.defaultCookiePreferences();
    }

    /**
     * Builds the object.
     */
    build(): preferences.CookiePreferences {
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
