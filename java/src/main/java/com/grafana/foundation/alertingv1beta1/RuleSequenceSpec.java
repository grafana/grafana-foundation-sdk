// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.LinkedList;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class RuleSequenceSpec {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("trigger")
    public IntervalTrigger trigger;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("recordingRules")
    public List<RuleRef> recordingRules;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alertRules")
    public List<RuleRef> alertRules;
    public RuleSequenceSpec() {
        this.trigger = new com.grafana.foundation.alertingv1beta1.IntervalTrigger();
        this.recordingRules = new LinkedList<>();
    }
    public RuleSequenceSpec(IntervalTrigger trigger,List<RuleRef> recordingRules,List<RuleRef> alertRules) {
        this.trigger = trigger;
        this.recordingRules = recordingRules;
        this.alertRules = alertRules;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
