// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import java.util.List;
import java.util.LinkedList;

public class RuleSequenceBuilder implements com.grafana.foundation.cog.Builder<RuleSequenceSpec> {
    protected final RuleSequenceSpec internal;
    
    public RuleSequenceBuilder() {
        this.internal = new RuleSequenceSpec();
    }
    public RuleSequenceBuilder trigger(com.grafana.foundation.cog.Builder<IntervalTrigger> trigger) {
    IntervalTrigger triggerResource = trigger.build();
        this.internal.trigger = triggerResource;
        return this;
    }
    
    public RuleSequenceBuilder recordingRules(List<com.grafana.foundation.cog.Builder<RuleRef>> recordingRules) {
        List<RuleRef> recordingRulesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<RuleRef> r1 : recordingRules) {
                RuleRef recordingRulesDepth1 = r1.build();
                recordingRulesResources.add(recordingRulesDepth1); 
        }
        this.internal.recordingRules = recordingRulesResources;
        return this;
    }
    
    public RuleSequenceBuilder alertRules(List<com.grafana.foundation.cog.Builder<RuleRef>> alertRules) {
        List<RuleRef> alertRulesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<RuleRef> r1 : alertRules) {
                RuleRef alertRulesDepth1 = r1.build();
                alertRulesResources.add(alertRulesDepth1); 
        }
        this.internal.alertRules = alertRulesResources;
        return this;
    }
    public RuleSequenceSpec build() {
        return this.internal;
    }
}
