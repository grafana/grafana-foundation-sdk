// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class TextDimensionConfig {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public TextDimensionMode mode;
    // fixed: T -- will be added by each element
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("field")
    public String field;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("fixed")
    public String fixed;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TextDimensionConfig> {
        protected final TextDimensionConfig internal;
        
        public Builder() {
            this.internal = new TextDimensionConfig();
        }
    public Builder mode(TextDimensionMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder field(String field) {
    this.internal.field = field;
        return this;
    }
    
    public Builder fixed(String fixed) {
    this.internal.fixed = fixed;
        return this;
    }
    public TextDimensionConfig build() {
            return this.internal;
        }
    }
}
