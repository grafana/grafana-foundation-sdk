// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statetimeline;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.TimelineValueAlignment;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import java.util.List;

public class Options {
    // Show timeline values on chart
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("showValue")
    public VisibilityMode showValue;
    // Controls the row height
    @JsonProperty("rowHeight")
    public Double rowHeight;
    // Merge equal consecutive values
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mergeValues")
    public Boolean mergeValues;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timezone")
    public List<String> timezone;
    // Controls value alignment on the timelines
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("alignValue")
    public TimelineValueAlignment alignValue;
    public Options() {
        this.showValue = VisibilityMode.AUTO;
        this.rowHeight = 0.9;
        this.mergeValues = true;
        this.legend = new com.grafana.foundation.common.VizLegendOptions();
        this.tooltip = new com.grafana.foundation.common.VizTooltipOptions();
        this.alignValue = TimelineValueAlignment.LEFT;
    }
    public Options(VisibilityMode showValue,Double rowHeight,Boolean mergeValues,VizLegendOptions legend,VizTooltipOptions tooltip,List<String> timezone,TimelineValueAlignment alignValue) {
        this.showValue = showValue;
        this.rowHeight = rowHeight;
        this.mergeValues = mergeValues;
        this.legend = legend;
        this.tooltip = tooltip;
        this.timezone = timezone;
        this.alignValue = alignValue;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
