// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class ExprTypeResampleTimeRange {
    // From is the start time of the query. 
    @JsonProperty("from")
    public String from;
    // To is the end time of the query. 
    @JsonProperty("to")
    public String to;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ExprTypeResampleTimeRange> {
        private final ExprTypeResampleTimeRange internal;
        
        public Builder() {
            this.internal = new ExprTypeResampleTimeRange();
        this.from("now-6h");
        this.to("now");
        }
    public Builder from(String from) {
    this.internal.from = from;
        return this;
    }
    
    public Builder to(String to) {
    this.internal.to = to;
        return this;
    }
    public ExprTypeResampleTimeRange build() {
            return this.internal;
        }
    }
}
