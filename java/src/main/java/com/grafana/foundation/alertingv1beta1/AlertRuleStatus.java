// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class AlertRuleStatus {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("health")
    public AlertRuleHealth health;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("state")
    public AlertRuleState state;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("stateReason")
    public AlertRuleStateReason stateReason;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lastEvaluationTime")
    public String lastEvaluationTime;
    // duration of the last evaluation in seconds
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("evaluationDuration")
    public Double evaluationDuration;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("lastError")
    public String lastError;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("totals")
    public AlertRuleInstanceTotals totals;
    public AlertRuleStatus() {
    }
    public AlertRuleStatus(AlertRuleHealth health,AlertRuleState state,AlertRuleStateReason stateReason,String lastEvaluationTime,Double evaluationDuration,String lastError,AlertRuleInstanceTotals totals) {
        this.health = health;
        this.state = state;
        this.stateReason = stateReason;
        this.lastEvaluationTime = lastEvaluationTime;
        this.evaluationDuration = evaluationDuration;
        this.lastError = lastError;
        this.totals = totals;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
