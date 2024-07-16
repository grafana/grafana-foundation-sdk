// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class PulseWaveQuery { 
    @JsonProperty("timeStep")
    public Long timeStep; 
    @JsonProperty("onCount")
    public Long onCount; 
    @JsonProperty("offCount")
    public Long offCount; 
    @JsonProperty("onValue")
    public Double onValue; 
    @JsonProperty("offValue")
    public Double offValue;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<PulseWaveQuery> {
        private final PulseWaveQuery internal;
        
        public Builder() {
            this.internal = new PulseWaveQuery();
        }
    public Builder timeStep(Long timeStep) {
    this.internal.timeStep = timeStep;
        return this;
    }
    
    public Builder onCount(Long onCount) {
    this.internal.onCount = onCount;
        return this;
    }
    
    public Builder offCount(Long offCount) {
    this.internal.offCount = offCount;
        return this;
    }
    
    public Builder onValue(Double onValue) {
    this.internal.onValue = onValue;
        return this;
    }
    
    public Builder offValue(Double offValue) {
    this.internal.offValue = offValue;
        return this;
    }
    public PulseWaveQuery build() {
            return this.internal;
        }
    }
}
