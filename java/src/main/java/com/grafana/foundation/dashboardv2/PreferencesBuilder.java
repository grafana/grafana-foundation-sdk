// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;


public class PreferencesBuilder implements com.grafana.foundation.cog.Builder<Preferences> {
    protected final Preferences internal;
    
    public PreferencesBuilder() {
        this.internal = new Preferences();
    }
    public PreferencesBuilder layout(AutoGridLayoutKindOrGridLayoutKind layout) {
        this.internal.layout = layout;
        return this;
    }
    public Preferences build() {
        return this.internal;
    }
}
