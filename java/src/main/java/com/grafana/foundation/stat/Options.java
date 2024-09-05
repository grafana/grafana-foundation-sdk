// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.stat;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.grafana.foundation.common.BigValueGraphMode;
import com.grafana.foundation.common.BigValueColorMode;
import com.grafana.foundation.common.BigValueJustifyMode;
import com.grafana.foundation.common.ReduceDataOptions;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.grafana.foundation.common.VizTextDisplayOptions;
import com.grafana.foundation.common.BigValueTextMode;
import com.grafana.foundation.common.VizOrientation;

public class Options {
    @JsonProperty("graphMode")
    public BigValueGraphMode graphMode;
    @JsonProperty("colorMode")
    public BigValueColorMode colorMode;
    @JsonProperty("justifyMode")
    public BigValueJustifyMode justifyMode;
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    @JsonProperty("textMode")
    public BigValueTextMode textMode;
    @JsonProperty("orientation")
    public VizOrientation orientation;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
