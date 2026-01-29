// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class LibraryPanelBuilder implements com.grafana.foundation.cog.Builder<LibraryPanelKind> {
    protected final LibraryPanelKind internal;
    
    public LibraryPanelBuilder() {
        this.internal = new LibraryPanelKind();
        this.internal.kind = "LibraryPanel";
    }
    public LibraryPanelBuilder spec(com.grafana.foundation.cog.Builder<LibraryPanelKindSpec> spec) {
    LibraryPanelKindSpec specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    
    public LibraryPanelBuilder id(Double id) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.LibraryPanelKindSpecBuilder().build();
		}
        this.internal.spec.id = id;
        return this;
    }
    
    public LibraryPanelBuilder title(String title) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.LibraryPanelKindSpecBuilder().build();
		}
        this.internal.spec.title = title;
        return this;
    }
    
    public LibraryPanelBuilder libraryPanel(com.grafana.foundation.cog.Builder<LibraryPanelRef> libraryPanel) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.LibraryPanelKindSpecBuilder().build();
		}
    LibraryPanelRef libraryPanelResource = libraryPanel.build();
        this.internal.spec.libraryPanel = libraryPanelResource;
        return this;
    }
    public LibraryPanelKind build() {
        return this.internal;
    }
}
