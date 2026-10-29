// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import java.util.Map;
import java.util.HashMap;

public class AlertRuleBuilder implements com.grafana.foundation.cog.Builder<AlertRuleSpec> {
    protected final AlertRuleSpec internal;
    
    public AlertRuleBuilder() {
        this.internal = new AlertRuleSpec();
    }
    public AlertRuleBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public AlertRuleBuilder paused(Boolean paused) {
        this.internal.paused = paused;
        return this;
    }
    
    public AlertRuleBuilder trigger(com.grafana.foundation.cog.Builder<IntervalTrigger> trigger) {
    IntervalTrigger triggerResource = trigger.build();
        this.internal.trigger = triggerResource;
        return this;
    }
    
    public AlertRuleBuilder labels(Map<String, String> labels) {
        this.internal.labels = labels;
        return this;
    }
    
    public AlertRuleBuilder annotations(Map<String, String> annotations) {
        this.internal.annotations = annotations;
        return this;
    }
    
    public AlertRuleBuilder forArg(String forArg) {
        this.internal.forArg = forArg;
        return this;
    }
    
    public AlertRuleBuilder keepFiringFor(String keepFiringFor) {
        this.internal.keepFiringFor = keepFiringFor;
        return this;
    }
    
    public AlertRuleBuilder missingSeriesEvalsToResolve(Long missingSeriesEvalsToResolve) {
        if (!(missingSeriesEvalsToResolve >= 0)) {
            throw new IllegalArgumentException("missingSeriesEvalsToResolve must be >= 0");
        }
        this.internal.missingSeriesEvalsToResolve = missingSeriesEvalsToResolve;
        return this;
    }
    
    public AlertRuleBuilder noDataState(NoDataState noDataState) {
        this.internal.noDataState = noDataState;
        return this;
    }
    
    public AlertRuleBuilder execErrState(ExecErrState execErrState) {
        this.internal.execErrState = execErrState;
        return this;
    }
    
    public AlertRuleBuilder notificationSettings(NotificationSettings notificationSettings) {
        this.internal.notificationSettings = notificationSettings;
        return this;
    }
    
    public AlertRuleBuilder expressions(Map<String, com.grafana.foundation.cog.Builder<Expression>> expressions) {
        Map<String, Expression> expressionsResources = new HashMap<>();
        for (var entry1 : expressions.entrySet()) {
                Expression expressionsDepth1 = entry1.getValue().build();
                expressionsResources.put(entry1.getKey(), expressionsDepth1);           
        }
        this.internal.expressions = expressionsResources;
        return this;
    }
    
    public AlertRuleBuilder panelRef(com.grafana.foundation.cog.Builder<PanelRef> panelRef) {
    PanelRef panelRefResource = panelRef.build();
        this.internal.panelRef = panelRefResource;
        return this;
    }
    public AlertRuleSpec build() {
        return this.internal;
    }
}
