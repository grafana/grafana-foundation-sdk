// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.timeseries;

import com.grafana.foundation.common.LegendDisplayMode;
import com.grafana.foundation.common.LegendPlacement;
import com.grafana.foundation.common.VizLegendOptions;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.grafana.foundation.common.VizTooltipOptions;
import com.grafana.foundation.common.VizOrientation;

public class Options {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timezone")
    public List<String> timezone;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("orientation")
    public VizOrientation orientation;
    public Options() {
        this.legend = new VizLegendOptions(LegendDisplayMode.LIST, LegendPlacement.BOTTOM, false, false, false, "", false, 0.0, List.of());
    }
    
    public Options(List<String> timezone,VizLegendOptions legend,VizTooltipOptions tooltip,VizOrientation orientation) {
        this.timezone = timezone;
        this.legend = legend;
        this.tooltip = tooltip;
        this.orientation = orientation;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
