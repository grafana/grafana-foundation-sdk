// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.HashMap;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.Map;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;

public class RecordingRuleSpec {
    @JsonProperty("title")
    public String title;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("paused")
    public Boolean paused;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("trigger")
    public IntervalTrigger trigger;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labels")
    public Map<String, String> labels;
    @JsonProperty("metric")
    public String metric;
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("expressions")
    public Map<String, Expression> expressions;
    @JsonProperty("targetDatasourceUID")
    public String targetDatasourceUID;
    public RecordingRuleSpec() {
        this.title = "";
        this.trigger = new com.grafana.foundation.alertingv1beta1.IntervalTrigger();
        this.metric = "";
        this.expressions = new HashMap<>();
        this.targetDatasourceUID = "";
    }
    public RecordingRuleSpec(String title,Boolean paused,IntervalTrigger trigger,Map<String, String> labels,String metric,Map<String, Expression> expressions,String targetDatasourceUID) {
        this.title = title;
        this.paused = paused;
        this.trigger = trigger;
        this.labels = labels;
        this.metric = metric;
        this.expressions = expressions;
        this.targetDatasourceUID = targetDatasourceUID;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
