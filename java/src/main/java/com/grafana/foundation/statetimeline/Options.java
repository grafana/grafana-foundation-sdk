// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.statetimeline;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.VisibilityMode;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.grafana.foundation.common.TimelineValueAlignment;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.VizTooltipOptions;
import java.util.List;

public class Options {
    // Show timeline values on chart
    @JsonProperty("showValue")
    public VisibilityMode showValue;
    // Controls the row height
    @JsonProperty("rowHeight")
    public Double rowHeight;
    // Merge equal consecutive values
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("mergeValues")
    public Boolean mergeValues;
    // Controls value alignment on the timelines
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("alignValue")
    public TimelineValueAlignment alignValue;
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonProperty("tooltip")
    public VizTooltipOptions tooltip;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("timezone")
    public List<String> timezone;
    // Enables pagination when > 0
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("perPage")
    public Double perPage;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
