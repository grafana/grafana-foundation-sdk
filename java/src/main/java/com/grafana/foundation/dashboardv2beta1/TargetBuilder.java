// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class TargetBuilder implements com.grafana.foundation.cog.Builder<PanelQueryKind> {
    protected final PanelQueryKind internal;
    
    public TargetBuilder() {
        this.internal = new PanelQueryKind();
        this.internal.kind = "PanelQuery";
    }
    public TargetBuilder query(com.grafana.foundation.cog.Builder<DataQueryKind> query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.PanelQuerySpec();
		}
    DataQueryKind queryResource = query.build();
        this.internal.spec.query = queryResource;
        return this;
    }
    
    public TargetBuilder refId(String refId) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.PanelQuerySpec();
		}
        this.internal.spec.refId = refId;
        return this;
    }
    
    public TargetBuilder hidden(Boolean hidden) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.PanelQuerySpec();
		}
        this.internal.spec.hidden = hidden;
        return this;
    }
    public PanelQueryKind build() {
        return this.internal;
    }
}
