// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

// A library panel is a reusable panel that you can use in any dashboard.
// When you make a change to a library panel, that change propagates to all instances of where the panel is used.
// Library panels streamline reuse of panels across multiple dashboards.
export class LibraryPanelRefBuilder implements cog.Builder<dashboardv2.LibraryPanelRef> {
    protected readonly internal: dashboardv2.LibraryPanelRef;

    constructor() {
        this.internal = dashboardv2.defaultLibraryPanelRef();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.LibraryPanelRef {
        return this.internal;
    }

    // Library panel name
    name(name: string): this {
        this.internal.name = name;
        return this;
    }

    // Library panel uid
    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }
}

