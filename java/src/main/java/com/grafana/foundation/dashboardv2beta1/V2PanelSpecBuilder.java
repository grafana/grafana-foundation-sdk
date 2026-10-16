// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class V2PanelSpecBuilder implements com.grafana.foundation.cog.Builder<V2PanelSpec> {
    protected final V2PanelSpec internal;
    
    public V2PanelSpecBuilder() {
        this.internal = new V2PanelSpec();
    }
    public V2PanelSpecBuilder id(Double id) {
        this.internal.id = id;
        return this;
    }
    
    public V2PanelSpecBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public V2PanelSpecBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public V2PanelSpecBuilder subtitle(String subtitle) {
        this.internal.subtitle = subtitle;
        return this;
    }
    
    public V2PanelSpecBuilder links(List<com.grafana.foundation.cog.Builder<DataLink>> links) {
        List<DataLink> linksResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<DataLink> r1 : links) {
                DataLink linksDepth1 = r1.build();
                linksResources.add(linksDepth1); 
        }
        this.internal.links = linksResources;
        return this;
    }
    
    public V2PanelSpecBuilder data(com.grafana.foundation.cog.Builder<V2QueryGroupKind> data) {
    V2QueryGroupKind dataResource = data.build();
        this.internal.data = dataResource;
        return this;
    }
    
    public V2PanelSpecBuilder vizConfig(com.grafana.foundation.cog.Builder<VizConfigKind> vizConfig) {
    VizConfigKind vizConfigResource = vizConfig.build();
        this.internal.vizConfig = vizConfigResource;
        return this;
    }
    
    public V2PanelSpecBuilder transparent(Boolean transparent) {
        this.internal.transparent = transparent;
        return this;
    }
    public V2PanelSpec build() {
        return this.internal;
    }
}
