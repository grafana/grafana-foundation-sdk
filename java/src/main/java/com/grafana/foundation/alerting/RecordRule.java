// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class RecordRule {
    @JsonProperty("from")
    public String from;
    @JsonProperty("metric")
    public String metric;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<RecordRule> {
        protected final RecordRule internal;
        
        public Builder() {
            this.internal = new RecordRule();
        }
    public Builder from(String from) {
    this.internal.from = from;
        return this;
    }
    
    public Builder metric(String metric) {
    this.internal.metric = metric;
        return this;
    }
    public RecordRule build() {
            return this.internal;
        }
    }
}
