// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class AdhocFilters {
    @JsonProperty("key")
    public String key;
    @JsonProperty("operator")
    public String operator;
    @JsonProperty("value")
    public String value;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("values")
    public List<String> values;
    public AdhocFilters() {
        this.key = "";
        this.operator = "";
        this.value = "";
    }
    public AdhocFilters(String key,String operator,String value,List<String> values) {
        this.key = key;
        this.operator = operator;
        this.value = value;
        this.values = values;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
