// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class PrometheusDataqueryScope {
    @JsonProperty("matchers")
    public String matchers;
    public PrometheusDataqueryScope() {
        this.matchers = "";
    }
    public PrometheusDataqueryScope(String matchers) {
        this.matchers = matchers;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
