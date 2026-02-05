// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.datagrid;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class Options {
    @JsonProperty("selectedSeries")
    public Integer selectedSeries;
    public Options() {
        this.selectedSeries = 0;
    }
    public Options(Integer selectedSeries) {
        this.selectedSeries = selectedSeries;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
