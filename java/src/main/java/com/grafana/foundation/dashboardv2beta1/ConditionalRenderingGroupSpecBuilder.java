// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class ConditionalRenderingGroupSpecBuilder implements com.grafana.foundation.cog.Builder<ConditionalRenderingGroupSpec> {
    protected final ConditionalRenderingGroupSpec internal;
    
    public ConditionalRenderingGroupSpecBuilder() {
        this.internal = new ConditionalRenderingGroupSpec();
    }
    public ConditionalRenderingGroupSpecBuilder visibility(ConditionalRenderingGroupSpecVisibility visibility) {
        this.internal.visibility = visibility;
        return this;
    }
    
    public ConditionalRenderingGroupSpecBuilder condition(ConditionalRenderingGroupSpecCondition condition) {
        this.internal.condition = condition;
        return this;
    }
    
    public ConditionalRenderingGroupSpecBuilder items(List<ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind> items) {
        this.internal.items = items;
        return this;
    }
    public ConditionalRenderingGroupSpec build() {
        return this.internal;
    }
}
