// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

export class AnnotationPermissionBuilder implements cog.Builder<dashboard.AnnotationPermission> {
    protected readonly internal: dashboard.AnnotationPermission;

    constructor() {
        this.internal = dashboard.defaultAnnotationPermission();
    }

    /**
     * Builds the object.
     */
    build(): dashboard.AnnotationPermission {
        return this.internal;
    }

    dashboard(dashboard: cog.Builder<dashboard.AnnotationActions>): this {
        const dashboardResource = dashboard.build();
        this.internal.dashboard = dashboardResource;
        return this;
    }

    organization(organization: cog.Builder<dashboard.AnnotationActions>): this {
        const organizationResource = organization.build();
        this.internal.organization = organizationResource;
        return this;
    }
}
