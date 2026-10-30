// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;


public class PanelRefBuilder implements com.grafana.foundation.cog.Builder<PanelRef> {
    protected final PanelRef internal;
    
    public PanelRefBuilder() {
        this.internal = new PanelRef();
    }
    public PanelRefBuilder dashboardUID(String dashboardUID) {
        if (!(dashboardUID.length() >= 1)) {
            throw new IllegalArgumentException("dashboardUID.length() must be >= 1");
        }
        if (!dashboardUID.matches("^[a-zA-Z0-9_-]+$")) {
            throw new IllegalArgumentException("dashboardUID must match regex ^[a-zA-Z0-9_-]+$");
        }
        this.internal.dashboardUID = dashboardUID;
        return this;
    }
    
    public PanelRefBuilder panelID(Long panelID) {
        if (!(panelID > 0)) {
            throw new IllegalArgumentException("panelID must be > 0");
        }
        this.internal.panelID = panelID;
        return this;
    }
    public PanelRef build() {
        return this.internal;
    }
}
