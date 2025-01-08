// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.stat;

import com.grafana.foundation.common.BigValueGraphMode;
import com.grafana.foundation.common.BigValueColorMode;
import com.grafana.foundation.common.BigValueJustifyMode;
import com.grafana.foundation.common.BigValueTextMode;
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
    @JsonProperty("graphMode")
    public BigValueGraphMode graphMode;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("colorMode")
    public BigValueColorMode colorMode;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("justifyMode")
    public BigValueJustifyMode justifyMode;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("textMode")
    public BigValueTextMode textMode;
    @JsonProperty("wideLayout")
    public Boolean wideLayout;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    @JsonProperty("showPercentChange")
    public Boolean showPercentChange;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("orientation")
    public VizOrientation orientation;
    public Options() {
        this.graphMode = BigValueGraphMode.AREA;
        this.colorMode = BigValueColorMode.VALUE;
        this.justifyMode = BigValueJustifyMode.AUTO;
        this.textMode = BigValueTextMode.AUTO;
        this.wideLayout = true;
        this.showPercentChange = false;
    }
    
    public Options(BigValueGraphMode graphMode,BigValueColorMode colorMode,BigValueJustifyMode justifyMode,BigValueTextMode textMode,Boolean wideLayout,ReduceDataOptions reduceOptions,VizTextDisplayOptions text,Boolean showPercentChange,VizOrientation orientation) {
        this.graphMode = graphMode;
        this.colorMode = colorMode;
        this.justifyMode = justifyMode;
        this.textMode = textMode;
        this.wideLayout = wideLayout;
        this.reduceOptions = reduceOptions;
        this.text = text;
        this.showPercentChange = showPercentChange;
        this.orientation = orientation;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
