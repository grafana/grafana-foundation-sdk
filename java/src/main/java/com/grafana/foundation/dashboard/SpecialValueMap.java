// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
// See SpecialValueMatch to see the list of special values.
// For example, you can configure a special value mapping so that null values appear as N/A.
public class SpecialValueMap {
    @JsonProperty("type")
    public String type;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("options")
    public DashboardSpecialValueMapOptions options;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<SpecialValueMap> {
        protected final SpecialValueMap internal;
        
        public Builder() {
            this.internal = new SpecialValueMap();
    this.internal.type = "special";
        }
    public Builder options(com.grafana.foundation.cog.Builder<DashboardSpecialValueMapOptions> options) {
    this.internal.options = options.build();
        return this;
    }
    public SpecialValueMap build() {
            return this.internal;
        }
    }
}
