// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class CSVWave {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timeStep")
    public Long timeStep;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("name")
    public String name;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("valuesCSV")
    public String valuesCSV;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("labels")
    public String labels;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<CSVWave> {
        protected final CSVWave internal;
        
        public Builder() {
            this.internal = new CSVWave();
        }
    public Builder timeStep(Long timeStep) {
    this.internal.timeStep = timeStep;
        return this;
    }
    
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder valuesCSV(String valuesCSV) {
    this.internal.valuesCSV = valuesCSV;
        return this;
    }
    
    public Builder labels(String labels) {
    this.internal.labels = labels;
        return this;
    }
    public CSVWave build() {
            return this.internal;
        }
    }
}
