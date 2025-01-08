// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import java.util.List;

public class OptionsWithTimezonesBuilder implements com.grafana.foundation.cog.Builder<OptionsWithTimezones> {
    protected final OptionsWithTimezones internal;
    
    public OptionsWithTimezonesBuilder() {
        this.internal = new OptionsWithTimezones();
    }
    public OptionsWithTimezonesBuilder timezone(List<String> timezone) {
    this.internal.timezone = timezone;
        return this;
    }
    public OptionsWithTimezones build() {
        return this.internal;
    }
}
