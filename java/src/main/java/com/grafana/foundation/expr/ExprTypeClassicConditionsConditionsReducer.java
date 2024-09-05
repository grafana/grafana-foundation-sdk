// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ExprTypeClassicConditionsConditionsReducer {
    @JsonProperty("type")
    public String type;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsReducer> {
        private final ExprTypeClassicConditionsConditionsReducer internal;
        
        public Builder() {
            this.internal = new ExprTypeClassicConditionsConditionsReducer();
        }
    public Builder type(String type) {
    this.internal.type = type;
        return this;
    }
    public ExprTypeClassicConditionsConditionsReducer build() {
            return this.internal;
        }
    }
}
