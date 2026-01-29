// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class ConditionalRenderingGroupBuilder implements com.grafana.foundation.cog.Builder<ConditionalRenderingGroupKind> {
    protected final ConditionalRenderingGroupKind internal;
    
    public ConditionalRenderingGroupBuilder() {
        this.internal = new ConditionalRenderingGroupKind();
        this.internal.kind = "ConditionalRenderingGroup";
    }
    public ConditionalRenderingGroupBuilder spec(com.grafana.foundation.cog.Builder<ConditionalRenderingGroupSpec> spec) {
    ConditionalRenderingGroupSpec specResource = spec.build();
        this.internal.spec = specResource;
        return this;
    }
    
    public ConditionalRenderingGroupBuilder visibility(ConditionalRenderingGroupSpecVisibility visibility) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConditionalRenderingGroupSpecBuilder().build();
		}
        this.internal.spec.visibility = visibility;
        return this;
    }
    
    public ConditionalRenderingGroupBuilder condition(ConditionalRenderingGroupSpecCondition condition) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConditionalRenderingGroupSpecBuilder().build();
		}
        this.internal.spec.condition = condition;
        return this;
    }
    
    public ConditionalRenderingGroupBuilder items(List<ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind> items) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConditionalRenderingGroupSpecBuilder().build();
		}
        this.internal.spec.items = items;
        return this;
    }
    public ConditionalRenderingGroupKind build() {
        return this.internal;
    }
}
