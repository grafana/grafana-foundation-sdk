// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class XychartXYSeriesConfigX {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("matcher")
    public MatcherConfig matcher;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<XychartXYSeriesConfigX> {
        protected final XychartXYSeriesConfigX internal;
        
        public Builder() {
            this.internal = new XychartXYSeriesConfigX();
        }
    public Builder matcher(com.grafana.foundation.cog.Builder<MatcherConfig> matcher) {
    this.internal.matcher = matcher.build();
        return this;
    }
    public XychartXYSeriesConfigX build() {
            return this.internal;
        }
    }
}
