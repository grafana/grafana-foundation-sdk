// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.testdata;

import java.util.List;

public class USAQueryBuilder implements com.grafana.foundation.cog.Builder<USAQuery> {
    protected final USAQuery internal;
    
    public USAQueryBuilder() {
        this.internal = new USAQuery();
    }
    public USAQueryBuilder mode(String mode) {
    this.internal.mode = mode;
        return this;
    }
    
    public USAQueryBuilder period(String period) {
    this.internal.period = period;
        return this;
    }
    
    public USAQueryBuilder fields(List<String> fields) {
    this.internal.fields = fields;
        return this;
    }
    
    public USAQueryBuilder states(List<String> states) {
    this.internal.states = states;
        return this;
    }
    public USAQuery build() {
        return this.internal;
    }
}
