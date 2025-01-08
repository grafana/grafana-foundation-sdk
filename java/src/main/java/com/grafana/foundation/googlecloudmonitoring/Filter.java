// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

// Query filter representation.
public class Filter {
    // Filter key.
    @JsonProperty("key")
    public String key;
    // Filter operator.
    @JsonProperty("operator")
    public String operator;
    // Filter value.
    @JsonProperty("value")
    public String value;
    // Filter condition.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("condition")
    public String condition;
    public Filter() {
    }
    
    public Filter(String key,String operator,String value,String condition) {
        this.key = key;
        this.operator = operator;
        this.value = value;
        this.condition = condition;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
