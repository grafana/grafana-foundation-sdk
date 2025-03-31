// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as dashboard from '../dashboard';

// +k8s:deepcopy-gen=true
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

    // +k8s:deepcopy-gen=true
    dashboard(dashboard: cog.Builder<dashboard.AnnotationActions>): this {
        const dashboardResource = dashboard.build();
        this.internal.dashboard = dashboardResource;
        return this;
    }

    // +k8s:deepcopy-gen=true
    organization(organization: cog.Builder<dashboard.AnnotationActions>): this {
        const organizationResource = organization.build();
        this.internal.organization = organizationResource;
        return this;
    }
}
