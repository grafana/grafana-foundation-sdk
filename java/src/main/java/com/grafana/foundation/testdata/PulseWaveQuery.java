// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class PulseWaveQuery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("offCount")
    public Long offCount;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("offValue")
    public Double offValue;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onCount")
    public Long onCount;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("onValue")
    public Double onValue;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeStep")
    public Long timeStep;
    public PulseWaveQuery() {
    }
    
    public PulseWaveQuery(Long offCount,Double offValue,Long onCount,Double onValue,Long timeStep) {
        this.offCount = offCount;
        this.offValue = offValue;
        this.onCount = onCount;
        this.onValue = onValue;
        this.timeStep = timeStep;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
