// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Define the MetricFindValue type
public class MetricFindValue {
    @JsonProperty("text")
    public String text;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("value")
    public StringOrFloat64 value;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group")
    public String group;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("expandable")
    public Boolean expandable;
    public MetricFindValue() {
        this.text = "";
    }
    public MetricFindValue(String text,StringOrFloat64 value,String group,Boolean expandable) {
        this.text = text;
        this.value = value;
        this.group = group;
        this.expandable = expandable;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
