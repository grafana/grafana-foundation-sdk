// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statushistory;

import com.grafana.foundation.common.VisibilityMode;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
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
    // Controls the column width
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("colWidth")
    public Double colWidth;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timezone")
    public List<String> timezone;
    // Enables pagination when > 0
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("perPage")
    public Double perPage;
    public Options() {
        this.rowHeight = 0.9f;
        this.showValue = VisibilityMode.AUTO;
        this.colWidth = 0.9;
        this.perPage = 20.0;
    }
    
    public Options(Float rowHeight,VisibilityMode showValue,Double colWidth,VizLegendOptions legend,VizTooltipOptions tooltip,List<String> timezone,Double perPage) {
        this.rowHeight = rowHeight;
        this.showValue = showValue;
        this.colWidth = colWidth;
        this.legend = legend;
        this.tooltip = tooltip;
        this.timezone = timezone;
        this.perPage = perPage;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
