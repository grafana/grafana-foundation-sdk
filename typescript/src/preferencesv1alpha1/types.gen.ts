// Code generated - EDITING IS FUTILE. DO NOT EDIT.

export interface QueryHistoryPreference {
	// one of: '' | 'query' | 'starred';
	homeTab?: string;
}

export const defaultQueryHistoryPreference = (): QueryHistoryPreference => ({
});

export interface NavbarPreference {
	bookmarkUrls: string[];
}

export const defaultNavbarPreference = (): NavbarPreference => ({
	bookmarkUrls: [],
});

export interface Preferences {
	// Explicit home URL (NOTE: this can only be modified in the system settings)
	homeURL?: string;
	// UID for the home dashboard
	homeDashboardUID?: string;
	// The timezone selection
	timezone?: string;
	// day of the week (sunday, monday, etc)
	weekStart?: string;
	// user interface theme
	theme?: string;
	// Selected language
	language?: string;
	// Explore query history preferences
	queryHistory?: QueryHistoryPreference;
	// Navigation preferences
	navbar?: NavbarPreference;
}

export const defaultPreferences = (): Preferences => ({
});

export const PreferencesApiVersion = "preferences.grafana.app/v1alpha1";

export const PreferencesKind = "Preferences";

