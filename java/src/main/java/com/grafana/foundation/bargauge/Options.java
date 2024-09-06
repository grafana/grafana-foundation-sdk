// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bargauge;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.grafana.foundation.common.BarGaugeDisplayMode;
import com.grafana.foundation.common.BarGaugeValueMode;
import com.grafana.foundation.common.BarGaugeNamePlacement;
import com.grafana.foundation.common.BarGaugeSizing;
import com.grafana.foundation.common.VizLegendOptions;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.VizOrientation;

public class Options {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("displayMode")
    public BarGaugeDisplayMode displayMode;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("valueMode")
    public BarGaugeValueMode valueMode;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("namePlacement")
    public BarGaugeNamePlacement namePlacement;
    @JsonProperty("showUnfilled")
    public Boolean showUnfilled;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("sizing")
    public BarGaugeSizing sizing;
    @JsonProperty("minVizWidth")
    public Integer minVizWidth;
    @JsonProperty("minVizHeight")
    public Integer minVizHeight;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("legend")
    public VizLegendOptions legend;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    @JsonProperty("maxVizHeight")
    public Integer maxVizHeight;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("orientation")
    public VizOrientation orientation;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
