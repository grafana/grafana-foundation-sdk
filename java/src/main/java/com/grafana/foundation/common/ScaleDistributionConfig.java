// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// TODO docs
public class ScaleDistributionConfig {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("type")
    public ScaleDistribution type;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("log")
    public Double log;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("linearThreshold")
    public Double linearThreshold;
    public ScaleDistributionConfig() {
        this.type = ScaleDistribution.LINEAR;
    }
    public ScaleDistributionConfig(ScaleDistribution type,Double log,Double linearThreshold) {
        this.type = type;
        this.log = log;
        this.linearThreshold = linearThreshold;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
