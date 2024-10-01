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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<PulseWaveQuery> {
        protected final PulseWaveQuery internal;
        
        public Builder() {
            this.internal = new PulseWaveQuery();
        }
    public Builder offCount(Long offCount) {
    this.internal.offCount = offCount;
        return this;
    }
    
    public Builder offValue(Double offValue) {
    this.internal.offValue = offValue;
        return this;
    }
    
    public Builder onCount(Long onCount) {
    this.internal.onCount = onCount;
        return this;
    }
    
    public Builder onValue(Double onValue) {
    this.internal.onValue = onValue;
        return this;
    }
    
    public Builder timeStep(Long timeStep) {
    this.internal.timeStep = timeStep;
        return this;
    }
    public PulseWaveQuery build() {
            return this.internal;
        }
    }
}
