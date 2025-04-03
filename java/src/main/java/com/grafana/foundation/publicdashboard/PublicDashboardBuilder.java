// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.publicdashboard;


public class PublicDashboardBuilder implements com.grafana.foundation.cog.Builder<PublicDashboard> {
    protected final PublicDashboard internal;
    
    public PublicDashboardBuilder() {
        this.internal = new PublicDashboard();
    }
    public PublicDashboardBuilder uid(String uid) {
        this.internal.uid = uid;
        return this;
    }
    
    public PublicDashboardBuilder dashboardUid(String dashboardUid) {
        this.internal.dashboardUid = dashboardUid;
        return this;
    }
    
    public PublicDashboardBuilder accessToken(String accessToken) {
        this.internal.accessToken = accessToken;
        return this;
    }
    
    public PublicDashboardBuilder isEnabled(Boolean isEnabled) {
        this.internal.isEnabled = isEnabled;
        return this;
    }
    
    public PublicDashboardBuilder annotationsEnabled(Boolean annotationsEnabled) {
        this.internal.annotationsEnabled = annotationsEnabled;
        return this;
    }
    
    public PublicDashboardBuilder timeSelectionEnabled(Boolean timeSelectionEnabled) {
        this.internal.timeSelectionEnabled = timeSelectionEnabled;
        return this;
    }
    public PublicDashboard build() {
        return this.internal;
    }
}
