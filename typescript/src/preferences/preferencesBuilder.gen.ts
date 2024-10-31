// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as preferences from '../preferences';

export class PreferencesBuilder implements cog.Builder<preferences.Preferences> {
    protected readonly internal: preferences.Preferences;

    constructor() {
        this.internal = preferences.defaultPreferences();
    }

    /**
     * Builds the object.
     */
    build(): preferences.Preferences {
        return this.internal;
    }

    // UID for the home dashboard
    homeDashboardUID(homeDashboardUID: string): this {
        this.internal.homeDashboardUID = homeDashboardUID;
        return this;
    }

    // The timezone selection
    // TODO: this should use the timezone defined in common
    timezone(timezone: string): this {
        this.internal.timezone = timezone;
        return this;
    }

    // day of the week (sunday, monday, etc)
    weekStart(weekStart: string): this {
        this.internal.weekStart = weekStart;
        return this;
    }

    // light, dark, empty is default
    theme(theme: string): this {
        this.internal.theme = theme;
        return this;
    }

    // Selected language (beta)
    language(language: string): this {
        this.internal.language = language;
        return this;
    }

    // Explore query history preferences
    queryHistory(queryHistory: cog.Builder<preferences.QueryHistoryPreference>): this {
        const queryHistoryResource = queryHistory.build();
        this.internal.queryHistory = queryHistoryResource;
        return this;
    }

    // Cookie preferences
    cookiePreferences(cookiePreferences: cog.Builder<preferences.CookiePreferences>): this {
        const cookiePreferencesResource = cookiePreferences.build();
        this.internal.cookiePreferences = cookiePreferencesResource;
        return this;
    }
}
