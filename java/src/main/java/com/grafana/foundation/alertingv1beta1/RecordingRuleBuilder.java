// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import java.util.Map;
import java.util.HashMap;

public class RecordingRuleBuilder implements com.grafana.foundation.cog.Builder<RecordingRuleSpec> {
    protected final RecordingRuleSpec internal;
    
    public RecordingRuleBuilder() {
        this.internal = new RecordingRuleSpec();
    }
    public RecordingRuleBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public RecordingRuleBuilder paused(Boolean paused) {
        this.internal.paused = paused;
        return this;
    }
    
    public RecordingRuleBuilder trigger(com.grafana.foundation.cog.Builder<IntervalTrigger> trigger) {
    IntervalTrigger triggerResource = trigger.build();
        this.internal.trigger = triggerResource;
        return this;
    }
    
    public RecordingRuleBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public RecordingRuleBuilder metric(String metric) {
        if (!metric.matches("^[a-zA-Z_:][a-zA-Z0-9_:]*$")) {
            throw new IllegalArgumentException("metric must match regex ^[a-zA-Z_:][a-zA-Z0-9_:]*$");
        }
        this.internal.metric = metric;
        return this;
    }
    
    public RecordingRuleBuilder expressions(Map<String, com.grafana.foundation.cog.Builder<Expression>> expressions) {
        Map<String, Expression> expressionsResources = new HashMap<>();
        for (var entry1 : expressions.entrySet()) {
                Expression expressionsDepth1 = entry1.getValue().build();
                expressionsResources.put(entry1.getKey(), expressionsDepth1);           
        }
        this.internal.expressions = expressionsResources;
        return this;
    }
    
    public RecordingRuleBuilder targetDatasourceUID(String targetDatasourceUID) {
        if (!targetDatasourceUID.matches("^[a-zA-Z0-9_-]+$")) {
            throw new IllegalArgumentException("targetDatasourceUID must match regex ^[a-zA-Z0-9_-]+$");
        }
        this.internal.targetDatasourceUID = targetDatasourceUID;
        return this;
    }
    public RecordingRuleSpec build() {
        return this.internal;
    }
}
