// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as preferencesv1alpha1 from '../preferencesv1alpha1';

export class NavbarPreferenceBuilder implements cog.Builder<preferencesv1alpha1.NavbarPreference> {
    protected readonly internal: preferencesv1alpha1.NavbarPreference;

    constructor() {
        this.internal = preferencesv1alpha1.defaultNavbarPreference();
    }

    /**
     * Builds the object.
     */
    build(): preferencesv1alpha1.NavbarPreference {
        return this.internal;
    }

    bookmarkUrls(bookmarkUrls: string[]): this {
        this.internal.bookmarkUrls = bookmarkUrls;
        return this;
    }
}

