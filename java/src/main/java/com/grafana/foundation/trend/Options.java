// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.trend;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;

// Identical to timeseries... except it does not have timezone settings
public class Options {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    // Name of the x field to use (defaults to first number)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("xField")
    public String xField;
    public Options() {
        this.legend = new com.grafana.foundation.common.VizLegendOptions();
        this.tooltip = new com.grafana.foundation.common.VizTooltipOptions();
    }
    public Options(VizLegendOptions legend,VizTooltipOptions tooltip,String xField) {
        this.legend = legend;
        this.tooltip = tooltip;
        this.xField = xField;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
