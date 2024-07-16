// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;

public class NotificationSettings { 
    @JsonProperty("group_by")
    public List<String> groupBy; 
    @JsonProperty("group_interval")
    public String groupInterval; 
    @JsonProperty("group_wait")
    public String groupWait; 
    @JsonProperty("mute_time_intervals")
    public List<String> muteTimeIntervals; 
    @JsonProperty("receiver")
    public String receiver; 
    @JsonProperty("repeat_interval")
    public String repeatInterval;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<NotificationSettings> {
        private final NotificationSettings internal;
        
        public Builder() {
            this.internal = new NotificationSettings();
        this.groupBy(List.of("alertname", "grafana_folder"));
        }
    public Builder groupBy(List<String> groupBy) {
    this.internal.groupBy = groupBy;
        return this;
    }
    
    public Builder groupInterval(String groupInterval) {
    this.internal.groupInterval = groupInterval;
        return this;
    }
    
    public Builder groupWait(String groupWait) {
    this.internal.groupWait = groupWait;
        return this;
    }
    
    public Builder muteTimeIntervals(List<String> muteTimeIntervals) {
    this.internal.muteTimeIntervals = muteTimeIntervals;
        return this;
    }
    
    public Builder receiver(String receiver) {
    this.internal.receiver = receiver;
        return this;
    }
    
    public Builder repeatInterval(String repeatInterval) {
    this.internal.repeatInterval = repeatInterval;
        return this;
    }
    public NotificationSettings build() {
            return this.internal;
        }
    }
}
