// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;

public class ExprTypeThresholdConditions {
    @JsonProperty("evaluator")
    public ExprTypeThresholdConditionsEvaluator evaluator;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("loadedDimensions")
    public Object loadedDimensions;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("unloadEvaluator")
    public ExprTypeThresholdConditionsUnloadEvaluator unloadEvaluator;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ExprTypeThresholdConditions> {
        private final ExprTypeThresholdConditions internal;
        
        public Builder() {
            this.internal = new ExprTypeThresholdConditions();
        }
    public Builder evaluator(com.grafana.foundation.cog.Builder<ExprTypeThresholdConditionsEvaluator> evaluator) {
    this.internal.evaluator = evaluator.build();
        return this;
    }
    
    public Builder loadedDimensions(Object loadedDimensions) {
    this.internal.loadedDimensions = loadedDimensions;
        return this;
    }
    
    public Builder unloadEvaluator(com.grafana.foundation.cog.Builder<ExprTypeThresholdConditionsUnloadEvaluator> unloadEvaluator) {
    this.internal.unloadEvaluator = unloadEvaluator.build();
        return this;
    }
    public ExprTypeThresholdConditions build() {
            return this.internal;
        }
    }
}
