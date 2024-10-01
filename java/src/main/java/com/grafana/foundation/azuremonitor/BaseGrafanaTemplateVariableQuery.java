// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class BaseGrafanaTemplateVariableQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("rawQuery")
    public String rawQuery;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<BaseGrafanaTemplateVariableQuery> {
        protected final BaseGrafanaTemplateVariableQuery internal;
        
        public Builder() {
            this.internal = new BaseGrafanaTemplateVariableQuery();
        }
    public Builder rawQuery(String rawQuery) {
    this.internal.rawQuery = rawQuery;
        return this;
    }
    public BaseGrafanaTemplateVariableQuery build() {
            return this.internal;
        }
    }
}
