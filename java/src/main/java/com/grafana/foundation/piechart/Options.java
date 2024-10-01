// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.piechart;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;
import com.grafana.foundation.common.VizTooltipOptions;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;

public class Options {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("pieType")
    public PieChartType pieType;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("displayLabels")
    public List<PieChartLabels> displayLabels;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public PieChartLegendOptions legend;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("orientation")
    public VizOrientation orientation;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
