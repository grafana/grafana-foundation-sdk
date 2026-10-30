// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class RecordingRuleStatus {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("health")
    public RecordingRuleHealth health;
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
    public RecordingRuleStatus() {
    }
    public RecordingRuleStatus(RecordingRuleHealth health,String lastEvaluationTime,Double evaluationDuration,String lastError) {
        this.health = health;
        this.lastEvaluationTime = lastEvaluationTime;
        this.evaluationDuration = evaluationDuration;
        this.lastError = lastError;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
