// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statushistory;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.VisibilityMode;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import java.util.List;

public class Options {
    // Set the height of the rows
    @JsonProperty("rowHeight")
    public Float rowHeight;
    // Show values on the columns
    @JsonProperty("showValue")
    public VisibilityMode showValue;
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("timezone")
    public List<String> timezone;
    // Controls the column width
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("colWidth")
    public Double colWidth;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
