// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonSetter;
import com.fasterxml.jackson.annotation.Nulls;
import java.util.List;

public class TimePicker {
    // Whether timepicker is visible or not.
    @JsonProperty("hidden")
    public Boolean hidden;
    // Interval options available in the refresh picker dropdown.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("refresh_intervals")
    public List<String> refreshIntervals;
    // Whether timepicker is collapsed or not. Has no effect on provisioned dashboard.
    @JsonProperty("collapse")
    public Boolean collapse;
    // Whether timepicker is enabled or not. Has no effect on provisioned dashboard.
    @JsonProperty("enable")
    public Boolean enable;
    // Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    @JsonSetter(nulls = Nulls.AS_EMPTY)
    @JsonProperty("time_options")
    public List<String> timeOptions;
    public TimePicker() {
        this.hidden = false;
        this.refreshIntervals = List.of("5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d");
        this.collapse = false;
        this.enable = true;
        this.timeOptions = List.of("5m", "15m", "1h", "6h", "12h", "24h", "2d", "7d", "30d");
    }
    public TimePicker(Boolean hidden,List<String> refreshIntervals,Boolean collapse,Boolean enable,List<String> timeOptions) {
        this.hidden = hidden;
        this.refreshIntervals = refreshIntervals;
        this.collapse = collapse;
        this.enable = enable;
        this.timeOptions = timeOptions;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
