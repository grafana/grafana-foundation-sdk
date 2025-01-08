// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// TimeInterval describes intervals of time. ContainsTime will tell you if a golang time is contained
// within the interval.
public class TimeInterval {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("days_of_month")
    public List<String> daysOfMonth;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("location")
    public String location;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("months")
    public List<String> months;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("times")
    public List<TimeRange> times;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("weekdays")
    public List<String> weekdays;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("years")
    public List<String> years;
    public TimeInterval() {
    }
    
    public TimeInterval(List<String> daysOfMonth,String location,List<String> months,List<TimeRange> times,List<String> weekdays,List<String> years) {
        this.daysOfMonth = daysOfMonth;
        this.location = location;
        this.months = months;
        this.times = times;
        this.weekdays = weekdays;
        this.years = years;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
