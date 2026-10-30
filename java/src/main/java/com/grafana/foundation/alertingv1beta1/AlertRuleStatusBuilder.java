// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;


public class AlertRuleStatusBuilder implements com.grafana.foundation.cog.Builder<AlertRuleStatus> {
    protected final AlertRuleStatus internal;
    
    public AlertRuleStatusBuilder() {
        this.internal = new AlertRuleStatus();
    }
    public AlertRuleStatusBuilder health(AlertRuleHealth health) {
        this.internal.health = health;
        return this;
    }
    
    public AlertRuleStatusBuilder state(AlertRuleState state) {
        this.internal.state = state;
        return this;
    }
    
    public AlertRuleStatusBuilder stateReason(AlertRuleStateReason stateReason) {
        this.internal.stateReason = stateReason;
        return this;
    }
    
    public AlertRuleStatusBuilder lastEvaluationTime(String lastEvaluationTime) {
        this.internal.lastEvaluationTime = lastEvaluationTime;
        return this;
    }
    
    public AlertRuleStatusBuilder evaluationDuration(Double evaluationDuration) {
        this.internal.evaluationDuration = evaluationDuration;
        return this;
    }
    
    public AlertRuleStatusBuilder lastError(String lastError) {
        this.internal.lastError = lastError;
        return this;
    }
    
    public AlertRuleStatusBuilder totals(com.grafana.foundation.cog.Builder<AlertRuleInstanceTotals> totals) {
    AlertRuleInstanceTotals totalsResource = totals.build();
        this.internal.totals = totalsResource;
        return this;
    }
    public AlertRuleStatus build() {
        return this.internal;
    }
}
