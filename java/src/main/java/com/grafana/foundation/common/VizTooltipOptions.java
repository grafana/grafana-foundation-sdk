// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class VizTooltipOptions {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("mode")
    public TooltipDisplayMode mode;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("sort")
    public SortOrder sort;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxWidth")
    public Double maxWidth;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("maxHeight")
    public Double maxHeight;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<VizTooltipOptions> {
        protected final VizTooltipOptions internal;
        
        public Builder() {
            this.internal = new VizTooltipOptions();
        }
    public Builder mode(TooltipDisplayMode mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public Builder sort(SortOrder sort) {
    this.internal.sort = sort;
        return this;
    }
    
    public Builder maxWidth(Double maxWidth) {
    this.internal.maxWidth = maxWidth;
        return this;
    }
    
    public Builder maxHeight(Double maxHeight) {
    this.internal.maxHeight = maxHeight;
        return this;
    }
    public VizTooltipOptions build() {
            return this.internal;
        }
    }
}
