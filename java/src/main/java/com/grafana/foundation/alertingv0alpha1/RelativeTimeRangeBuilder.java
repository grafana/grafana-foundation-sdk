// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;


public class RelativeTimeRangeBuilder implements com.grafana.foundation.cog.Builder<RelativeTimeRange> {
    protected final RelativeTimeRange internal;
    
    public RelativeTimeRangeBuilder() {
        this.internal = new RelativeTimeRange();
    }
    public RelativeTimeRangeBuilder from(String from) {
        if (!from.matches("^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)$")) {
            throw new IllegalArgumentException("from must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)$");
        }
        this.internal.from = from;
        return this;
    }
    
    public RelativeTimeRangeBuilder to(String to) {
        if (!to.matches("^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)$")) {
            throw new IllegalArgumentException("to must match regex ^((([0-9]+)y)?(([0-9]+)w)?(([0-9]+)d)?(([0-9]+)h)?(([0-9]+)m)?(([0-9]+)s)?(([0-9]+)ms)?|0)$");
        }
        this.internal.to = to;
        return this;
    }
    public RelativeTimeRange build() {
        return this.internal;
    }
}
