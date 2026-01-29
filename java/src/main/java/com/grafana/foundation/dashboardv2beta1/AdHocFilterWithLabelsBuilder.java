// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.List;

public class AdHocFilterWithLabelsBuilder implements com.grafana.foundation.cog.Builder<AdHocFilterWithLabels> {
    protected final AdHocFilterWithLabels internal;
    
    public AdHocFilterWithLabelsBuilder() {
        this.internal = new AdHocFilterWithLabels();
    }
    public AdHocFilterWithLabelsBuilder key(String key) {
        this.internal.key = key;
        return this;
    }
    
    public AdHocFilterWithLabelsBuilder operator(String operator) {
        this.internal.operator = operator;
        return this;
    }
    
    public AdHocFilterWithLabelsBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    
    public AdHocFilterWithLabelsBuilder values(List<String> values) {
        this.internal.values = values;
        return this;
    }
    
    public AdHocFilterWithLabelsBuilder keyLabel(String keyLabel) {
        this.internal.keyLabel = keyLabel;
        return this;
    }
    
    public AdHocFilterWithLabelsBuilder valueLabels(List<String> valueLabels) {
        this.internal.valueLabels = valueLabels;
        return this;
    }
    
    public AdHocFilterWithLabelsBuilder forceEdit(Boolean forceEdit) {
        this.internal.forceEdit = forceEdit;
        return this;
    }
    
    public AdHocFilterWithLabelsBuilder origin(String origin) {
        this.internal.origin = origin;
        return this;
    }
    
    public AdHocFilterWithLabelsBuilder condition(String condition) {
        this.internal.condition = condition;
        return this;
    }
    public AdHocFilterWithLabels build() {
        return this.internal;
    }
}
