// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class PrometheusDataqueryScope {
    @JsonProperty("matchers")
    public String matchers;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<PrometheusDataqueryScope> {
        protected final PrometheusDataqueryScope internal;
        
        public Builder() {
            this.internal = new PrometheusDataqueryScope();
        }
    public Builder matchers(String matchers) {
    this.internal.matchers = matchers;
        return this;
    }
    public PrometheusDataqueryScope build() {
            return this.internal;
        }
    }
}
