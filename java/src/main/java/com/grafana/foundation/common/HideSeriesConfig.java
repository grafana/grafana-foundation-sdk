// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

// TODO docs
public class HideSeriesConfig {
    @JsonProperty("tooltip")
    public Boolean tooltip;
    @JsonProperty("legend")
    public Boolean legend;
    @JsonProperty("viz")
    public Boolean viz;
    public HideSeriesConfig() {
    }
    public HideSeriesConfig(Boolean tooltip,Boolean legend,Boolean viz) {
        this.tooltip = tooltip;
        this.legend = legend;
        this.viz = viz;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
