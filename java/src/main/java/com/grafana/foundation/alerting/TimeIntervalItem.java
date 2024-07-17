// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class TimeIntervalItem { 
    @JsonProperty("days_of_month")
    public List<String> daysOfMonth; 
    @JsonProperty("location")
    public String location; 
    @JsonProperty("months")
    public List<String> months; 
    @JsonProperty("times")
    public List<TimeIntervalTimeRange> times; 
    @JsonProperty("weekdays")
    public List<String> weekdays; 
    @JsonProperty("years")
    public List<String> years;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TimeIntervalItem> {
        private final TimeIntervalItem internal;
        
        public Builder() {
            this.internal = new TimeIntervalItem();
        }
    public Builder daysOfMonth(List<String> daysOfMonth) {
    this.internal.daysOfMonth = daysOfMonth;
        return this;
    }
    
    public Builder location(String location) {
    this.internal.location = location;
        return this;
    }
    
    public Builder months(List<String> months) {
    this.internal.months = months;
        return this;
    }
    
    public Builder times(com.grafana.foundation.cog.Builder<List<TimeIntervalTimeRange>> times) {
    this.internal.times = times.build();
        return this;
    }
    
    public Builder weekdays(List<String> weekdays) {
    this.internal.weekdays = weekdays;
        return this;
    }
    
    public Builder years(List<String> years) {
    this.internal.years = years;
        return this;
    }
    public TimeIntervalItem build() {
            return this.internal;
        }
    }
}
