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
    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("time_options")
    public List<String> timeOptions;
    // Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("nowDelay")
    public String nowDelay;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TimePickerConfig> {
        protected final TimePickerConfig internal;
        
        public Builder() {
            this.internal = new TimePickerConfig();
        this.hidden(false);
        this.refreshIntervals(List.of("5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"));
        this.timeOptions(List.of("5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d"));
        }
    public Builder hidden(Boolean hidden) {
    this.internal.hidden = hidden;
        return this;
    }
    
    public Builder refreshIntervals(List<String> refreshIntervals) {
    this.internal.refreshIntervals = refreshIntervals;
        return this;
    }
    
    public Builder timeOptions(List<String> timeOptions) {
    this.internal.timeOptions = timeOptions;
        return this;
    }
    
    public Builder nowDelay(String nowDelay) {
    this.internal.nowDelay = nowDelay;
        return this;
    }
    public TimePickerConfig build() {
            return this.internal;
        }
    }
}
