// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;

public class IntervalVariableBuilder implements com.grafana.foundation.cog.Builder<VariableModel> {
    protected final VariableModel internal;
    
    public IntervalVariableBuilder(String name) {
        this.internal = new VariableModel();
    this.internal.name = name;
    this.internal.type = VariableType.INTERVAL;
    }
    public IntervalVariableBuilder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public IntervalVariableBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public IntervalVariableBuilder label(String label) {
    this.internal.label = label;
        return this;
    }
    
    public IntervalVariableBuilder hide(VariableHide hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public IntervalVariableBuilder description(String description) {
    this.internal.description = description;
        return this;
    }
    
    public IntervalVariableBuilder values(StringOrMap query) {
    this.internal.query = query;
        return this;
    }
    
    public IntervalVariableBuilder allFormat(String allFormat) {
    this.internal.allFormat = allFormat;
        return this;
    }
    
    public IntervalVariableBuilder current(VariableOption current) {
    this.internal.current = current;
        return this;
    }
    
    public IntervalVariableBuilder options(List<VariableOption> options) {
    this.internal.options = options;
        return this;
    }
    
    public IntervalVariableBuilder auto(Boolean auto) {
    this.internal.auto = auto;
        return this;
    }
    
    public IntervalVariableBuilder minInterval(String autoMin) {
    this.internal.autoMin = autoMin;
        return this;
    }
    
    public IntervalVariableBuilder stepCount(Integer autoCount) {
        if (!(autoCount > 0)) {
            throw new IllegalArgumentException("autoCount must be > 0");
        }
    this.internal.autoCount = autoCount;
        return this;
    }
    public VariableModel build() {
        return this.internal;
    }
}
