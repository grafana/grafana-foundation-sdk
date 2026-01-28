// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class LibraryPanelKindSpecBuilder implements com.grafana.foundation.cog.Builder<LibraryPanelKindSpec> {
    protected final LibraryPanelKindSpec internal;
    
    public LibraryPanelKindSpecBuilder() {
        this.internal = new LibraryPanelKindSpec();
    }
    public LibraryPanelKindSpecBuilder id(Double id) {
        this.internal.id = id;
        return this;
    }
    
    public LibraryPanelKindSpecBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public LibraryPanelKindSpecBuilder libraryPanel(com.grafana.foundation.cog.Builder<LibraryPanelRef> libraryPanel) {
    LibraryPanelRef libraryPanelResource = libraryPanel.build();
        this.internal.libraryPanel = libraryPanelResource;
        return this;
    }
    public LibraryPanelKindSpec build() {
        return this.internal;
    }
}
