// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class TimeInterval {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("times")
    public List<TimeRange> times;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("weekdays")
    public List<WeekdayRange> weekdays;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("days_of_month")
    public List<DayOfMonthRange> daysOfMonth;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("months")
    public List<MonthRange> months;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("years")
    public List<YearRange> years;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("location")
    public String location;
    public TimeInterval() {
    }
    public TimeInterval(List<TimeRange> times,List<WeekdayRange> weekdays,List<DayOfMonthRange> daysOfMonth,List<MonthRange> months,List<YearRange> years,String location) {
        this.times = times;
        this.weekdays = weekdays;
        this.daysOfMonth = daysOfMonth;
        this.months = months;
        this.years = years;
        this.location = location;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
