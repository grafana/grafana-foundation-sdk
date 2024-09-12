// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ExprTypeClassicConditionsConditionsOperator {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public TypeClassicConditionsType type;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsOperator> {
        private final ExprTypeClassicConditionsConditionsOperator internal;
        
        public Builder() {
            this.internal = new ExprTypeClassicConditionsConditionsOperator();
        }
    public Builder type(TypeClassicConditionsType type) {
    this.internal.type = type;
        return this;
    }
    public ExprTypeClassicConditionsConditionsOperator build() {
            return this.internal;
        }
    }
}
