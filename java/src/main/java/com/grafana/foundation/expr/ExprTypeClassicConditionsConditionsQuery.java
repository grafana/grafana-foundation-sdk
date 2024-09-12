// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class ExprTypeClassicConditionsConditionsQuery {
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("params")
    public List<String> params;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsQuery> {
        private final ExprTypeClassicConditionsConditionsQuery internal;
        
        public Builder() {
            this.internal = new ExprTypeClassicConditionsConditionsQuery();
        }
    public Builder params(List<String> params) {
    this.internal.params = params;
        return this;
    }
    public ExprTypeClassicConditionsConditionsQuery build() {
            return this.internal;
        }
    }
}
