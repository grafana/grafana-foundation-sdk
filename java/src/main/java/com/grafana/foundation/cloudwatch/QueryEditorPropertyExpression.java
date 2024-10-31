// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryEditorPropertyExpression {
    @JsonProperty("type")
    public String type;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("property")
    public QueryEditorProperty property;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<QueryEditorPropertyExpression> {
        protected final QueryEditorPropertyExpression internal;
        
        public Builder() {
            this.internal = new QueryEditorPropertyExpression();
    this.internal.type = "property";
        }
    public Builder property(com.grafana.foundation.cog.Builder<QueryEditorProperty> property) {
    this.internal.property = property.build();
        return this;
    }
    public QueryEditorPropertyExpression build() {
            return this.internal;
        }
    }
}
