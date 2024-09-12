// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class OptionsWithTooltip {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<OptionsWithTooltip> {
        private final OptionsWithTooltip internal;
        
        public Builder() {
            this.internal = new OptionsWithTooltip();
        }
    public Builder tooltip(com.grafana.foundation.cog.Builder<VizTooltipOptions> tooltip) {
    this.internal.tooltip = tooltip.build();
        return this;
    }
    public OptionsWithTooltip build() {
            return this.internal;
        }
    }
}
