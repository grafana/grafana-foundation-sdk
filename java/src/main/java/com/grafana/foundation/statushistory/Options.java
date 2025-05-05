// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statushistory;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.VisibilityMode;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import java.util.List;

public class Options {
    // Set the height of the rows
    @JsonProperty("rowHeight")
    public Float rowHeight;
    // Show values on the columns
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("showValue")
    public VisibilityMode showValue;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timezone")
    public List<String> timezone;
    // Controls the column width
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("colWidth")
    public Double colWidth;
    public Options() {
        this.rowHeight = 0.9f;
        this.showValue = VisibilityMode.AUTO;
        this.legend = new com.grafana.foundation.common.VizLegendOptions();
        this.tooltip = new com.grafana.foundation.common.VizTooltipOptions();
        this.colWidth = 0.9;
    }
    public Options(Float rowHeight,VisibilityMode showValue,VizLegendOptions legend,VizTooltipOptions tooltip,List<String> timezone,Double colWidth) {
        this.rowHeight = rowHeight;
        this.showValue = showValue;
        this.legend = legend;
        this.tooltip = tooltip;
        this.timezone = timezone;
        this.colWidth = colWidth;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
