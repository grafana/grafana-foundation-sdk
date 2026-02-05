// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class SingleStatBaseOptions {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("reduceOptions")
    public ReduceDataOptions reduceOptions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("text")
    public VizTextDisplayOptions text;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("orientation")
    public VizOrientation orientation;
    public SingleStatBaseOptions() {
        this.reduceOptions = new com.grafana.foundation.common.ReduceDataOptions();
        this.orientation = VizOrientation.AUTO;
    }
    public SingleStatBaseOptions(ReduceDataOptions reduceOptions,VizTextDisplayOptions text,VizOrientation orientation) {
        this.reduceOptions = reduceOptions;
        this.text = text;
        this.orientation = orientation;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
