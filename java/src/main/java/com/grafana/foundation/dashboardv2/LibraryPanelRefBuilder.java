// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;


/**
 * A library panel is a reusable panel that you can use in any dashboard.
 * When you make a change to a library panel, that change propagates to all instances of where the panel is used.
 * Library panels streamline reuse of panels across multiple dashboards.
 */
public class LibraryPanelRefBuilder implements com.grafana.foundation.cog.Builder<LibraryPanelRef> {
    protected final LibraryPanelRef internal;
    
    public LibraryPanelRefBuilder() {
        this.internal = new LibraryPanelRef();
    }
    public LibraryPanelRefBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public LibraryPanelRefBuilder uid(String uid) {
        this.internal.uid = uid;
        return this;
    }
    public LibraryPanelRef build() {
        return this.internal;
    }
}
