// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class QueryEditorGroupByExpression {
    @JsonProperty("type")
    public String type;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("property")
    public QueryEditorProperty property;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<QueryEditorGroupByExpression> {
        protected final QueryEditorGroupByExpression internal;
        
        public Builder() {
            this.internal = new QueryEditorGroupByExpression();
    this.internal.type = "groupBy";
        }
    public Builder property(com.grafana.foundation.cog.Builder<QueryEditorProperty> property) {
    this.internal.property = property.build();
        return this;
    }
    public QueryEditorGroupByExpression build() {
            return this.internal;
        }
    }
}
