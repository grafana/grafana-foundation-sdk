// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboardv2 from '../dashboardv2';

// Dashboard specific preferences (applied per dashboard = all users using the dashboard)
export class PreferencesBuilder implements cog.Builder<dashboardv2.Preferences> {
    protected readonly internal: dashboardv2.Preferences;

    constructor() {
        this.internal = dashboardv2.defaultPreferences();
    }

    /**
     * Builds the object.
     */
    build(): dashboardv2.Preferences {
        return this.internal;
    }

    // default layout template to be used when new containers are created
    layout(layout: cog.Builder<dashboardv2.AutoGridLayoutKind> | cog.Builder<dashboardv2.GridLayoutKind>): this {
        const layoutResource = layout.build();
        this.internal.layout = layoutResource;
        return this;
    }
}

