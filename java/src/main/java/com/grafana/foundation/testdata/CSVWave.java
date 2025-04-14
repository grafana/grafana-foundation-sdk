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
    public CSVWave() {
    }
    public CSVWave(Long timeStep,String name,String valuesCSV,String labels) {
        this.timeStep = timeStep;
        this.name = name;
        this.valuesCSV = valuesCSV;
        this.labels = labels;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
