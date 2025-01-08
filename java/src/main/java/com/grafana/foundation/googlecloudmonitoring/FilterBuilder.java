// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;


public class FilterBuilder implements com.grafana.foundation.cog.Builder<Filter> {
    protected final Filter internal;
    
    public FilterBuilder() {
        this.internal = new Filter();
    }
    public FilterBuilder key(String key) {
    this.internal.key = key;
        return this;
    }
    
    public FilterBuilder operator(String operator) {
    this.internal.operator = operator;
        return this;
    }
    
    public FilterBuilder value(String value) {
    this.internal.value = value;
        return this;
    }
    
    public FilterBuilder condition(String condition) {
    this.internal.condition = condition;
        return this;
    }
    public Filter build() {
        return this.internal;
    }
}
