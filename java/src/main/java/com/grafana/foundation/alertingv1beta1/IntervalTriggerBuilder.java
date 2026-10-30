// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;


public class IntervalTriggerBuilder implements com.grafana.foundation.cog.Builder<IntervalTrigger> {
    protected final IntervalTrigger internal;
    
    public IntervalTriggerBuilder() {
        this.internal = new IntervalTrigger();
    }
    public IntervalTriggerBuilder interval(String interval) {
        if (!interval.matches("^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$")) {
            throw new IllegalArgumentException("interval must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?|0)$");
        }
        if (interval.matches("hmuµn")) {
            throw new IllegalArgumentException("interval must not match regex hmuµn");
        }
        this.internal.interval = interval;
        return this;
    }
    public IntervalTrigger build() {
        return this.internal;
    }
}
