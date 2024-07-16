// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.stat;

import com.grafana.foundation.common.BigValueGraphMode;
import com.grafana.foundation.common.BigValueColorMode;
import com.grafana.foundation.common.BigValueJustifyMode;
import com.grafana.foundation.common.BigValueTextMode;
import com.grafana.foundation.common.ReduceDataOptions;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.PercentChangeColorMode;
import com.grafana.foundation.common.VizOrientation;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Options { 
    @JsonProperty("graphMode")
    public BigValueGraphMode graphMode; 
    @JsonProperty("colorMode")
    public BigValueColorMode colorMode; 
    @JsonProperty("justifyMode")
    public BigValueJustifyMode justifyMode; 
    @JsonProperty("textMode")
    public BigValueTextMode textMode; 
    @JsonProperty("wideLayout")
    public Boolean wideLayout; 
    @JsonProperty("showPercentChange")
    public Boolean showPercentChange; 
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions; 
    @JsonProperty("text")
    public VizTextDisplayOptions text; 
    @JsonProperty("percentChangeColorMode")
    public PercentChangeColorMode percentChangeColorMode; 
    @JsonProperty("orientation")
    public VizOrientation orientation;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
