// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class ConditionalRenderingTimeRangeSizeBuilder implements com.grafana.foundation.cog.Builder<ConditionalRenderingTimeRangeSizeKind> {
    protected final ConditionalRenderingTimeRangeSizeKind internal;
    
    public ConditionalRenderingTimeRangeSizeBuilder() {
        this.internal = new ConditionalRenderingTimeRangeSizeKind();
        this.internal.kind = "ConditionalRenderingTimeRangeSize";
    }
    public ConditionalRenderingTimeRangeSizeBuilder value(String value) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec();
		}
        this.internal.spec.value = value;
        return this;
    }
    public ConditionalRenderingTimeRangeSizeKind build() {
        return this.internal;
    }
}
