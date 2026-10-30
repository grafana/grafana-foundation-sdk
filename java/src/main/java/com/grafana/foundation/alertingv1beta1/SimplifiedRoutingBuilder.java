// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;

import java.util.List;

public class SimplifiedRoutingBuilder implements com.grafana.foundation.cog.Builder<SimplifiedRouting> {
    protected final SimplifiedRouting internal;
    
    public SimplifiedRoutingBuilder() {
        this.internal = new SimplifiedRouting();
    }
    public SimplifiedRoutingBuilder receiver(String receiver) {
        this.internal.receiver = receiver;
        return this;
    }
    
    public SimplifiedRoutingBuilder groupBy(List<String> groupBy) {
        this.internal.groupBy = groupBy;
        return this;
    }
    
    public SimplifiedRoutingBuilder groupWait(String groupWait) {
        if (!groupWait.matches("^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$")) {
            throw new IllegalArgumentException("groupWait must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$");
        }
        if (groupWait.matches("hmuµn")) {
            throw new IllegalArgumentException("groupWait must not match regex hmuµn");
        }
        this.internal.groupWait = groupWait;
        return this;
    }
    
    public SimplifiedRoutingBuilder groupInterval(String groupInterval) {
        if (!groupInterval.matches("^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$")) {
            throw new IllegalArgumentException("groupInterval must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$");
        }
        if (groupInterval.matches("hmuµn")) {
            throw new IllegalArgumentException("groupInterval must not match regex hmuµn");
        }
        this.internal.groupInterval = groupInterval;
        return this;
    }
    
    public SimplifiedRoutingBuilder repeatInterval(String repeatInterval) {
        if (!repeatInterval.matches("^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$")) {
            throw new IllegalArgumentException("repeatInterval must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$");
        }
        if (repeatInterval.matches("hmuµn")) {
            throw new IllegalArgumentException("repeatInterval must not match regex hmuµn");
        }
        this.internal.repeatInterval = repeatInterval;
        return this;
    }
    
    public SimplifiedRoutingBuilder muteTimeIntervals(List<String> muteTimeIntervals) {
        this.internal.muteTimeIntervals = muteTimeIntervals;
        return this;
    }
    
    public SimplifiedRoutingBuilder activeTimeIntervals(List<String> activeTimeIntervals) {
        this.internal.activeTimeIntervals = activeTimeIntervals;
        return this;
    }
    public SimplifiedRouting build() {
        return this.internal;
    }
}
