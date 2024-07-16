// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class CSVWave { 
    @JsonProperty("labels")
    public String labels; 
    @JsonProperty("name")
    public String name; 
    @JsonProperty("timeStep")
    public Long timeStep; 
    @JsonProperty("valuesCSV")
    public String valuesCSV;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CSVWave> {
        private final CSVWave internal;
        
        public Builder() {
            this.internal = new CSVWave();
        }
    public Builder labels(String labels) {
    this.internal.labels = labels;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder timeStep(Long timeStep) {
    this.internal.timeStep = timeStep;
        return this;
    }
    
    public Builder valuesCSV(String valuesCSV) {
    this.internal.valuesCSV = valuesCSV;
        return this;
    }
    public CSVWave build() {
            return this.internal;
        }
    }
}
