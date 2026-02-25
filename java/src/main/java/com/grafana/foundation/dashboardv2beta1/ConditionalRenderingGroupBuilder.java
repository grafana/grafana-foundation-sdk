// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;
import java.util.LinkedList;

public class ConditionalRenderingGroupBuilder implements com.grafana.foundation.cog.Builder<ConditionalRenderingGroupKind> {
    protected final ConditionalRenderingGroupKind internal;
    
    public ConditionalRenderingGroupBuilder() {
        this.internal = new ConditionalRenderingGroupKind();
        this.internal.kind = "ConditionalRenderingGroup";
    }
    public ConditionalRenderingGroupBuilder visibility(ConditionalRenderingGroupSpecVisibility visibility) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConditionalRenderingGroupSpec();
		}
        this.internal.spec.visibility = visibility;
        return this;
    }
    
    public ConditionalRenderingGroupBuilder condition(ConditionalRenderingGroupSpecCondition condition) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConditionalRenderingGroupSpec();
		}
        this.internal.spec.condition = condition;
        return this;
    }
    
    public ConditionalRenderingGroupBuilder items(List<ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind> items) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConditionalRenderingGroupSpec();
		}
        this.internal.spec.items = items;
        return this;
    }
    
    public ConditionalRenderingGroupBuilder item(ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind item) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.ConditionalRenderingGroupSpec();
		}
		if (this.internal.spec.items == null) {
			this.internal.spec.items = new LinkedList<>();
		}
        this.internal.spec.items.add(item);
        return this;
    }
    public ConditionalRenderingGroupKind build() {
        return this.internal;
    }
}
