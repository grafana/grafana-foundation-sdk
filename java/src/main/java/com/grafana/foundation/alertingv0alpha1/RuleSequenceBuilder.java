// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

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
    
    public RuleSequenceBuilder alertingRules(List<com.grafana.foundation.cog.Builder<RuleRef>> alertingRules) {
        List<RuleRef> alertingRulesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<RuleRef> r1 : alertingRules) {
                RuleRef alertingRulesDepth1 = r1.build();
                alertingRulesResources.add(alertingRulesDepth1); 
        }
        this.internal.alertingRules = alertingRulesResources;
        return this;
    }
    public RuleSequenceSpec build() {
        return this.internal;
    }
}
