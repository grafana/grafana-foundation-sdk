// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ExprTypeThresholdConditions {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("evaluator")
    public ExprTypeThresholdConditionsEvaluator evaluator;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("loadedDimensions")
    public Object loadedDimensions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("unloadEvaluator")
    public ExprTypeThresholdConditionsUnloadEvaluator unloadEvaluator;
    public ExprTypeThresholdConditions() {
    }
    
    public ExprTypeThresholdConditions(ExprTypeThresholdConditionsEvaluator evaluator,Object loadedDimensions,ExprTypeThresholdConditionsUnloadEvaluator unloadEvaluator) {
        this.evaluator = evaluator;
        this.loadedDimensions = loadedDimensions;
        this.unloadEvaluator = unloadEvaluator;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
