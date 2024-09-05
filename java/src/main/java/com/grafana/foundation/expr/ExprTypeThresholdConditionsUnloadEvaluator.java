// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import java.util.List;

public class ExprTypeThresholdConditionsUnloadEvaluator {
    @JsonProperty("params")
    public List<Double> params;
    // e.g. "gt"
    @JsonProperty("type")
    public TypeThresholdType type;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ExprTypeThresholdConditionsUnloadEvaluator> {
        private final ExprTypeThresholdConditionsUnloadEvaluator internal;
        
        public Builder() {
            this.internal = new ExprTypeThresholdConditionsUnloadEvaluator();
        }
    public Builder params(List<Double> params) {
    this.internal.params = params;
        return this;
    }
    
    public Builder type(TypeThresholdType type) {
    this.internal.type = type;
        return this;
    }
    public ExprTypeThresholdConditionsUnloadEvaluator build() {
            return this.internal;
        }
    }
}
