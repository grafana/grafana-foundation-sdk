// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import java.util.List;

public class MuteTiming {
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("name")
    public String name;
    @JsonSerialize(include = JsonSerialize.Inclusion.NON_NULL)
    @JsonProperty("time_intervals")
    public List<TimeInterval> timeIntervals;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<MuteTiming> {
        private final MuteTiming internal;
        
        public Builder() {
            this.internal = new MuteTiming();
        }
    public Builder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public Builder timeIntervals(com.grafana.foundation.cog.Builder<List<TimeInterval>> timeIntervals) {
    this.internal.timeIntervals = timeIntervals.build();
        return this;
    }
    public MuteTiming build() {
            return this.internal;
        }
    }
}
