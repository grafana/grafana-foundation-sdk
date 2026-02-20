// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as preferencesv1alpha1 from '../preferencesv1alpha1';
import * as resource from '../resource';

export class PreferencesBuilder implements cog.Builder<preferencesv1alpha1.Preferences> {
    protected readonly internal: preferencesv1alpha1.Preferences;

    constructor() {
        this.internal = preferencesv1alpha1.defaultPreferences();
    }

    /**
     * Builds the object.
     */
    build(): preferencesv1alpha1.Preferences {
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

    // Selected locale (beta)
    regionalFormat(regionalFormat: string): this {
        this.internal.regionalFormat = regionalFormat;
        return this;
    }

    // Explore query history preferences
    queryHistory(queryHistory: cog.Builder<preferencesv1alpha1.QueryHistoryPreference>): this {
        const queryHistoryResource = queryHistory.build();
        this.internal.queryHistory = queryHistoryResource;
        return this;
    }

    // Cookie preferences
    cookiePreferences(cookiePreferences: cog.Builder<preferencesv1alpha1.CookiePreferences>): this {
        const cookiePreferencesResource = cookiePreferences.build();
        this.internal.cookiePreferences = cookiePreferencesResource;
        return this;
    }

    // Navigation preferences
    navbar(navbar: cog.Builder<preferencesv1alpha1.NavbarPreference>): this {
        const navbarResource = navbar.build();
        this.internal.navbar = navbarResource;
        return this;
    }
}

/**
 * Creates a resource manifest from Preferences.
 */
export function manifest(name: string,preferences: cog.Builder<preferencesv1alpha1.Preferences>): cog.Builder<resource.Manifest> {
	const builder = new resource.ManifestBuilder();
	builder.apiVersion("preferences.grafana.app/v1alpha1");
	builder.kind("Preferences");
	builder.metadata(resource.named(name));
	builder.spec(preferences.build());

	return builder;
}

