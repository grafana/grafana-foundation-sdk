// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class AzureResourceGraphQuery {
    // Azure Resource Graph KQL query to be executed. 
    @JsonProperty("query")
    public String query;
    // Specifies the format results should be returned as. Defaults to table. 
    @JsonProperty("resultFormat")
    public String resultFormat;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<AzureResourceGraphQuery> {
        private final AzureResourceGraphQuery internal;
        
        public Builder() {
            this.internal = new AzureResourceGraphQuery();
        }
    public Builder query(String query) {
    this.internal.query = query;
        return this;
    }
    
    public Builder resultFormat(String resultFormat) {
    this.internal.resultFormat = resultFormat;
        return this;
    }
    public AzureResourceGraphQuery build() {
            return this.internal;
        }
    }
}
