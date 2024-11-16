// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.xychart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class XychartFieldConfigPointSize {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fixed")
    public Integer fixed;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("min")
    public Integer min;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("max")
    public Integer max;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<XychartFieldConfigPointSize> {
        protected final XychartFieldConfigPointSize internal;
        
        public Builder() {
            this.internal = new XychartFieldConfigPointSize();
        }
    public Builder fixed(Integer fixed) {
        if (!(fixed >= 0)) {
            throw new IllegalArgumentException("fixed must be >= 0");
        }
    this.internal.fixed = fixed;
        return this;
    }
    
    public Builder min(Integer min) {
        if (!(min >= 0)) {
            throw new IllegalArgumentException("min must be >= 0");
        }
    this.internal.min = min;
        return this;
    }
    
    public Builder max(Integer max) {
        if (!(max >= 0)) {
            throw new IllegalArgumentException("max must be >= 0");
        }
    this.internal.max = max;
        return this;
    }
    public XychartFieldConfigPointSize build() {
            return this.internal;
        }
    }
}
