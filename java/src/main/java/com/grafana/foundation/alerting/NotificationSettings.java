// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class NotificationSettings {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_by")
    public List<String> groupBy;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_interval")
    public String groupInterval;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_wait")
    public String groupWait;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mute_time_intervals")
    public List<String> muteTimeIntervals;
    @JsonProperty("receiver")
    public String receiver;
    @JsonInclude(JsonInclude.Include.NON_NULL)
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
