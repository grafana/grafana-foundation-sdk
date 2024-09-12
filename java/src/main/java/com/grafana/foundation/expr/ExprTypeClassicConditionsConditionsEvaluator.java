// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class ExprTypeClassicConditionsConditionsEvaluator {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("params")
    public List<Double> params;
    // e.g. "gt"
    @JsonProperty("type")
    public String type;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsEvaluator> {
        private final ExprTypeClassicConditionsConditionsEvaluator internal;
        
        public Builder() {
            this.internal = new ExprTypeClassicConditionsConditionsEvaluator();
        }
    public Builder params(List<Double> params) {
    this.internal.params = params;
        return this;
    }
    
    public Builder type(String type) {
    this.internal.type = type;
        return this;
    }
    public ExprTypeClassicConditionsConditionsEvaluator build() {
            return this.internal;
        }
    }
}
