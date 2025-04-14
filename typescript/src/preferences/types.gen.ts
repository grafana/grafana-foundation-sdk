// Code generated - EDITING IS FUTILE. DO NOT EDIT.

// Spec defines user, team or org Grafana preferences
// swagger:model Preferences
export interface Preferences {
	// UID for the home dashboard
	homeDashboardUID?: string;
	// The timezone selection
	// TODO: this should use the timezone defined in common
	timezone?: string;
	// day of the week (sunday, monday, etc)
	weekStart?: string;
	// light, dark, empty is default
	theme?: string;
	// Selected language (beta)
	language?: string;
	// Selected locale (beta)
	locale?: string;
	// Explore query history preferences
	queryHistory?: QueryHistoryPreference;
	// Cookie preferences
	cookiePreferences?: CookiePreferences;
	// Navigation preferences
	navbar?: NavbarPreference;
}

export const defaultPreferences = (): Preferences => ({
});

export interface QueryHistoryPreference {
	// one of: '' | 'query' | 'starred';
	homeTab?: string;
}

export const defaultQueryHistoryPreference = (): QueryHistoryPreference => ({
});

export interface CookiePreferences {
	analytics?: any;
	performance?: any;
	functional?: any;
}

export const defaultCookiePreferences = (): CookiePreferences => ({
});

export interface NavbarPreference {
	bookmarkUrls: string[];
}

export const defaultNavbarPreference = (): NavbarPreference => ({
	bookmarkUrls: [],
});

