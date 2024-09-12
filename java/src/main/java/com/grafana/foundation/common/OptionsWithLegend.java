// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class OptionsWithLegend {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public VizLegendOptions legend;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<OptionsWithLegend> {
        private final OptionsWithLegend internal;
        
        public Builder() {
            this.internal = new OptionsWithLegend();
        }
    public Builder legend(com.grafana.foundation.cog.Builder<VizLegendOptions> legend) {
    this.internal.legend = legend.build();
        return this;
    }
    public OptionsWithLegend build() {
            return this.internal;
        }
    }
}
