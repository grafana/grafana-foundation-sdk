// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


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
