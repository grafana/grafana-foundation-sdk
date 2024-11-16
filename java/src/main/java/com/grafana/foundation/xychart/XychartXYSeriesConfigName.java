// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class XychartXYSeriesConfigName {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fixed")
    public String fixed;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<XychartXYSeriesConfigName> {
        protected final XychartXYSeriesConfigName internal;
        
        public Builder() {
            this.internal = new XychartXYSeriesConfigName();
        }
    public Builder fixed(String fixed) {
    this.internal.fixed = fixed;
        return this;
    }
    public XychartXYSeriesConfigName build() {
            return this.internal;
        }
    }
}
