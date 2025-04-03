// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;

public class NotificationSettingsBuilder implements com.grafana.foundation.cog.Builder<NotificationSettings> {
    protected final NotificationSettings internal;
    
    public NotificationSettingsBuilder() {
        this.internal = new NotificationSettings();
    }
    public NotificationSettingsBuilder groupBy(List<String> groupBy) {
        this.internal.groupBy = groupBy;
        return this;
    }
    
    public NotificationSettingsBuilder groupInterval(String groupInterval) {
        this.internal.groupInterval = groupInterval;
        return this;
    }
    
    public NotificationSettingsBuilder groupWait(String groupWait) {
        this.internal.groupWait = groupWait;
        return this;
    }
    
    public NotificationSettingsBuilder muteTimeIntervals(List<String> muteTimeIntervals) {
        this.internal.muteTimeIntervals = muteTimeIntervals;
        return this;
    }
    
    public NotificationSettingsBuilder receiver(String receiver) {
        this.internal.receiver = receiver;
        return this;
    }
    
    public NotificationSettingsBuilder repeatInterval(String repeatInterval) {
        this.internal.repeatInterval = repeatInterval;
        return this;
    }
    public NotificationSettings build() {
        return this.internal;
    }
}
