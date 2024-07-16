// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class TimePicker {
    // Whether timepicker is visible or not. 
    @JsonProperty("hidden")
    public Boolean hidden;
    // Interval options available in the refresh picker dropdown. 
    @JsonProperty("refresh_intervals")
    public List<String> refreshIntervals;
    // Whether timepicker is collapsed or not. Has no effect on provisioned dashboard. 
    @JsonProperty("collapse")
    public Boolean collapse;
    // Whether timepicker is enabled or not. Has no effect on provisioned dashboard. 
    @JsonProperty("enable")
    public Boolean enable;
    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard. 
    @JsonProperty("time_options")
    public List<String> timeOptions;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<TimePicker> {
        private final TimePicker internal;
        
        public Builder() {
            this.internal = new TimePicker();
        this.hidden(false);
        this.refreshIntervals(List.of("5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"));
        this.collapse(false);
        this.enable(true);
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
    
    public Builder collapse(Boolean collapse) {
    this.internal.collapse = collapse;
        return this;
    }
    
    public Builder enable(Boolean enable) {
    this.internal.enable = enable;
        return this;
    }
    
    public Builder timeOptions(List<String> timeOptions) {
    this.internal.timeOptions = timeOptions;
        return this;
    }
    public TimePicker build() {
            return this.internal;
        }
    }
}
