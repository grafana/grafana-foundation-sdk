// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as publicdashboard from '../publicdashboard';

export class PublicDashboardBuilder implements cog.Builder<publicdashboard.PublicDashboard> {
    protected readonly internal: publicdashboard.PublicDashboard;

    constructor() {
        this.internal = publicdashboard.defaultPublicDashboard();
    }

    /**
     * Builds the object.
     */
    build(): publicdashboard.PublicDashboard {
        return this.internal;
    }

    // Unique public dashboard identifier
    uid(uid: string): this {
        this.internal.uid = uid;
        return this;
    }

    // Dashboard unique identifier referenced by this public dashboard
    dashboardUid(dashboardUid: string): this {
        this.internal.dashboardUid = dashboardUid;
        return this;
    }

    // Unique public access token
    accessToken(accessToken: string): this {
        this.internal.accessToken = accessToken;
        return this;
    }

    // Flag that indicates if the public dashboard is enabled
    isEnabled(isEnabled: boolean): this {
        this.internal.isEnabled = isEnabled;
        return this;
    }

    // Flag that indicates if annotations are enabled
    annotationsEnabled(annotationsEnabled: boolean): this {
        this.internal.annotationsEnabled = annotationsEnabled;
        return this;
    }

    // Flag that indicates if the time range picker is enabled
    timeSelectionEnabled(timeSelectionEnabled: boolean): this {
        this.internal.timeSelectionEnabled = timeSelectionEnabled;
        return this;
    }
}
