// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;


public class RecordingRuleStatusBuilder implements com.grafana.foundation.cog.Builder<RecordingRuleStatus> {
    protected final RecordingRuleStatus internal;
    
    public RecordingRuleStatusBuilder() {
        this.internal = new RecordingRuleStatus();
    }
    public RecordingRuleStatusBuilder health(RecordingRuleHealth health) {
        this.internal.health = health;
        return this;
    }
    
    public RecordingRuleStatusBuilder lastEvaluationTime(String lastEvaluationTime) {
        this.internal.lastEvaluationTime = lastEvaluationTime;
        return this;
    }
    
    public RecordingRuleStatusBuilder evaluationDuration(Double evaluationDuration) {
        this.internal.evaluationDuration = evaluationDuration;
        return this;
    }
    
    public RecordingRuleStatusBuilder lastError(String lastError) {
        this.internal.lastError = lastError;
        return this;
    }
    public RecordingRuleStatus build() {
        return this.internal;
    }
}
