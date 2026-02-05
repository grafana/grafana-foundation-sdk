// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

// Time picker configuration
// It defines the default config for the time picker and the refresh picker for the specific dashboard.
public class TimePickerConfig {
    // Whether timepicker is visible or not.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("hidden")
    public Boolean hidden;
    // Interval options available in the refresh picker dropdown.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("refresh_intervals")
    public List<String> refreshIntervals;
    // Quick ranges for time picker.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("quick_ranges")
    public List<TimeOption> quickRanges;
    // Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("nowDelay")
    public String nowDelay;
    public TimePickerConfig() {
        this.hidden = false;
        this.refreshIntervals = List.of("5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d");
    }
    public TimePickerConfig(Boolean hidden,List<String> refreshIntervals,List<TimeOption> quickRanges,String nowDelay) {
        this.hidden = hidden;
        this.refreshIntervals = refreshIntervals;
        this.quickRanges = quickRanges;
        this.nowDelay = nowDelay;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
