// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as preferences from '../preferences';

export class NavbarPreferenceBuilder implements cog.Builder<preferences.NavbarPreference> {
    protected readonly internal: preferences.NavbarPreference;

    constructor() {
        this.internal = preferences.defaultNavbarPreference();
    }

    /**
     * Builds the object.
     */
    build(): preferences.NavbarPreference {
        return this.internal;
    }

    bookmarkUrls(bookmarkUrls: string[]): this {
        this.internal.bookmarkUrls = bookmarkUrls;
        return this;
    }
}
