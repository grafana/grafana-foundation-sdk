// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ExprTypeThresholdConditionsUnloadEvaluator {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("params")
    public List<Double> params;
    // e.g. "gt"
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public TypeThresholdType type;
    public ExprTypeThresholdConditionsUnloadEvaluator() {
    }
    
    public ExprTypeThresholdConditionsUnloadEvaluator(List<Double> params,TypeThresholdType type) {
        this.params = params;
        this.type = type;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
