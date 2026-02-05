// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class PanelBuilder implements com.grafana.foundation.cog.Builder<PanelKind> {
    protected final PanelKind internal;
    
    public PanelBuilder() {
        this.internal = new PanelKind();
        this.internal.kind = "Panel";
    }
    public PanelBuilder id(Double id) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.PanelSpec();
		}
        this.internal.spec.id = id;
        return this;
    }
    
    public PanelBuilder title(String title) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.PanelSpec();
		}
        this.internal.spec.title = title;
        return this;
    }
    
    public PanelBuilder description(String description) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.PanelSpec();
		}
        this.internal.spec.description = description;
        return this;
    }
    
    public PanelBuilder links(List<com.grafana.foundation.cog.Builder<DataLink>> links) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.PanelSpec();
		}
        List<DataLink> linksResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<DataLink> r1 : links) {
                DataLink linksDepth1 = r1.build();
                linksResources.add(linksDepth1); 
        }
        this.internal.spec.links = linksResources;
        return this;
    }
    
    public PanelBuilder data(com.grafana.foundation.cog.Builder<QueryGroupKind> data) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.PanelSpec();
		}
    QueryGroupKind dataResource = data.build();
        this.internal.spec.data = dataResource;
        return this;
    }
    
    public PanelBuilder visualization(com.grafana.foundation.cog.Builder<VizConfigKind> vizConfig) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.PanelSpec();
		}
    VizConfigKind vizConfigResource = vizConfig.build();
        this.internal.spec.vizConfig = vizConfigResource;
        return this;
    }
    
    public PanelBuilder transparent(Boolean transparent) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.PanelSpec();
		}
        this.internal.spec.transparent = transparent;
        return this;
    }
    public PanelKind build() {
        return this.internal;
    }
}
