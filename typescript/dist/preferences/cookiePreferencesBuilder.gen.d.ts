import * as cog from '../cog';
import * as preferences from '../preferences';
export declare class CookiePreferencesBuilder implements cog.Builder<preferences.CookiePreferences> {
    protected readonly internal: preferences.CookiePreferences;
    constructor();
    /**
     * Builds the object.
     */
    build(): preferences.CookiePreferences;
    analytics(analytics: any): this;
    performance(performance: any): this;
    functional(functional: any): this;
}
