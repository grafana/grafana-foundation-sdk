// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

// Time configuration
// It defines the default time config for the time picker, the refresh picker for the specific dashboard.
public class TimeSettingsSpec {
    // Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("timezone")
    public String timezone;
    // Start time range for dashboard.
    // Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    @JsonProperty("from")
    public String from;
    // End time range for dashboard.
    // Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    @JsonProperty("to")
    public String to;
    // Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
    // v1: refresh
    @JsonProperty("autoRefresh")
    public String autoRefresh;
    // Interval options available in the refresh picker dropdown.
    // v1: timepicker.refresh_intervals
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("autoRefreshIntervals")
    public List<String> autoRefreshIntervals;
    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    // v1: timepicker.quick_ranges , not exposed in the UI
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("quickRanges")
    public List<TimeRangeOption> quickRanges;
    // Whether timepicker is visible or not.
    // v1: timepicker.hidden
    @JsonProperty("hideTimepicker")
    public Boolean hideTimepicker;
    // Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("weekStart")
    public TimeSettingsSpecWeekStart weekStart;
    // The month that the fiscal year starts on. 0 = January, 11 = December
    @JsonProperty("fiscalYearStartMonth")
    public Long fiscalYearStartMonth;
    // Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
    // v1: timepicker.nowDelay
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("nowDelay")
    public String nowDelay;
    public TimeSettingsSpec() {
        this.timezone = "browser";
        this.from = "now-6h";
        this.to = "now";
        this.autoRefresh = "";
        this.autoRefreshIntervals = List.of("5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d");
        this.hideTimepicker = false;
        this.fiscalYearStartMonth = 0L;
    }
    public TimeSettingsSpec(String timezone,String from,String to,String autoRefresh,List<String> autoRefreshIntervals,List<TimeRangeOption> quickRanges,Boolean hideTimepicker,TimeSettingsSpecWeekStart weekStart,Long fiscalYearStartMonth,String nowDelay) {
        this.timezone = timezone;
        this.from = from;
        this.to = to;
        this.autoRefresh = autoRefresh;
        this.autoRefreshIntervals = autoRefreshIntervals;
        this.quickRanges = quickRanges;
        this.hideTimepicker = hideTimepicker;
        this.weekStart = weekStart;
        this.fiscalYearStartMonth = fiscalYearStartMonth;
        this.nowDelay = nowDelay;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
