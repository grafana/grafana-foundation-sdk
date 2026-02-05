import * as cog from '../cog';
import * as preferences from '../preferences';
export declare class NavbarPreferenceBuilder implements cog.Builder<preferences.NavbarPreference> {
    protected readonly internal: preferences.NavbarPreference;
    constructor();
    /**
     * Builds the object.
     */
    build(): preferences.NavbarPreference;
    bookmarkUrls(bookmarkUrls: string[]): this;
}
