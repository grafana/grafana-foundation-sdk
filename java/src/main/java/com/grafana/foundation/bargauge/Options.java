// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.bargauge;

import com.grafana.foundation.common.BarGaugeDisplayMode;
import com.grafana.foundation.common.BarGaugeValueMode;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
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
    @JsonProperty("showUnfilled")
    public Boolean showUnfilled;
    @JsonProperty("minVizWidth")
    public Integer minVizWidth;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    @JsonProperty("minVizHeight")
    public Integer minVizHeight;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("orientation")
    public VizOrientation orientation;
    public Options() {
        this.displayMode = BarGaugeDisplayMode.GRADIENT;
        this.valueMode = BarGaugeValueMode.COLOR;
        this.showUnfilled = true;
        this.minVizWidth = 0;
        this.minVizHeight = 10;
    }
    
    public Options(BarGaugeDisplayMode displayMode,BarGaugeValueMode valueMode,Boolean showUnfilled,Integer minVizWidth,ReduceDataOptions reduceOptions,VizTextDisplayOptions text,Integer minVizHeight,VizOrientation orientation) {
        this.displayMode = displayMode;
        this.valueMode = valueMode;
        this.showUnfilled = showUnfilled;
        this.minVizWidth = minVizWidth;
        this.reduceOptions = reduceOptions;
        this.text = text;
        this.minVizHeight = minVizHeight;
        this.orientation = orientation;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
