// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// Configuration for the Table/Auto mode
public class XYDimensionConfig {
    @JsonProperty("frame")
    public Integer frame;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("x")
    public String x;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("exclude")
    public List<String> exclude;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<XYDimensionConfig> {
        protected final XYDimensionConfig internal;
        
        public Builder() {
            this.internal = new XYDimensionConfig();
        }
    public Builder frame(Integer frame) {
        if (!(frame >= 0)) {
            throw new IllegalArgumentException("frame must be >= 0");
        }
    this.internal.frame = frame;
        return this;
    }
    
    public Builder x(String x) {
    this.internal.x = x;
        return this;
    }
    
    public Builder exclude(List<String> exclude) {
    this.internal.exclude = exclude;
        return this;
    }
    public XYDimensionConfig build() {
            return this.internal;
        }
    }
}
