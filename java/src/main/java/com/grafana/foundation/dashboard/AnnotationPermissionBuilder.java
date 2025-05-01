// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class AnnotationPermissionBuilder implements com.grafana.foundation.cog.Builder<AnnotationPermission> {
    protected final AnnotationPermission internal;
    
    public AnnotationPermissionBuilder() {
        this.internal = new AnnotationPermission();
    }
    public AnnotationPermissionBuilder dashboard(com.grafana.foundation.cog.Builder<AnnotationActions> dashboard) {
    AnnotationActions dashboardResource = dashboard.build();
        this.internal.dashboard = dashboardResource;
        return this;
    }
    
    public AnnotationPermissionBuilder organization(com.grafana.foundation.cog.Builder<AnnotationActions> organization) {
    AnnotationActions organizationResource = organization.build();
        this.internal.organization = organizationResource;
        return this;
    }
    public AnnotationPermission build() {
        return this.internal;
    }
}
