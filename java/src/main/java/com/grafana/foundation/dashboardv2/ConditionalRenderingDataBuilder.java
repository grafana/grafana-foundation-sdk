// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2;


public class ConditionalRenderingDataBuilder implements com.grafana.foundation.cog.Builder<ConditionalRenderingDataKind> {
    protected final ConditionalRenderingDataKind internal;
    
    public ConditionalRenderingDataBuilder() {
        this.internal = new ConditionalRenderingDataKind();
        this.internal.kind = "ConditionalRenderingData";
    }
    public ConditionalRenderingDataBuilder value(Boolean value) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2.ConditionalRenderingDataSpec();
		}
        this.internal.spec.value = value;
        return this;
    }
    public ConditionalRenderingDataKind build() {
        return this.internal;
    }
}
