// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import java.util.Map;
import java.util.HashMap;

public class RuleSpecBuilder implements com.grafana.foundation.cog.Builder<RuleSpec> {
    protected final RuleSpec internal;
    
    public RuleSpecBuilder() {
        this.internal = new RuleSpec();
    }
    public RuleSpecBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public RuleSpecBuilder paused(Boolean paused) {
        this.internal.paused = paused;
        return this;
    }
    
    public RuleSpecBuilder trigger(com.grafana.foundation.cog.Builder<IntervalTrigger> trigger) {
    IntervalTrigger triggerResource = trigger.build();
        this.internal.trigger = triggerResource;
        return this;
    }
    
    public RuleSpecBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public RuleSpecBuilder expressions(Map<String, com.grafana.foundation.cog.Builder<Expression>> expressions) {
        Map<String, Expression> expressionsResources = new HashMap<>();
        for (var entry1 : expressions.entrySet()) {
                Expression expressionsDepth1 = entry1.getValue().build();
                expressionsResources.put(entry1.getKey(), expressionsDepth1);           
        }
        this.internal.expressions = expressionsResources;
        return this;
    }
    public RuleSpec build() {
        return this.internal;
    }
}
