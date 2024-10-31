// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;
import java.util.List;

public class NotificationSettings {
    // Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for
    // cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels
    // use the special value '...' as the sole label name.
    // This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what
    // you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
    // Must include 'alertname' and 'grafana_folder' if not using '...'.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_by")
    public List<String> groupBy;
    // Override how long to wait before sending a notification about new alerts that are added to a group of alerts for
    // which an initial notification has already been sent. (Usually ~5m or more.)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_interval")
    public String groupInterval;
    // Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an
    // inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("group_wait")
    public String groupWait;
    // Override the times when notifications should be muted. These must match the name of a mute time interval defined
    // in the alertmanager configuration mute_time_intervals section. When muted it will not send any notifications, but
    // otherwise acts normally.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("mute_time_intervals")
    public List<String> muteTimeIntervals;
    // Name of the receiver to send notifications to.
    @JsonProperty("receiver")
    public String receiver;
    // Override how long to wait before sending a notification again if it has already been sent successfully for an
    // alert. (Usually ~3h or more).
    // Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.
    // Notifications will be resent after either repeat_interval or the data retention period have passed, whichever
    // occurs first. `repeat_interval` should not be less than `group_interval`.
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonProperty("repeat_interval")
    public String repeatInterval;
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<NotificationSettings> {
        protected final NotificationSettings internal;
        
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
