export interface Preferences {
    homeDashboardUID?: string;
    timezone?: string;
    weekStart?: string;
    theme?: string;
    language?: string;
    queryHistory?: QueryHistoryPreference;
    cookiePreferences?: CookiePreferences;
    navbar?: NavbarPreference;
}
export declare const defaultPreferences: () => Preferences;
export interface QueryHistoryPreference {
    homeTab?: string;
}
export declare const defaultQueryHistoryPreference: () => QueryHistoryPreference;
export interface CookiePreferences {
    analytics?: any;
    performance?: any;
    functional?: any;
}
export declare const defaultCookiePreferences: () => CookiePreferences;
export interface NavbarPreference {
    bookmarkUrls: string[];
}
export declare const defaultNavbarPreference: () => NavbarPreference;
