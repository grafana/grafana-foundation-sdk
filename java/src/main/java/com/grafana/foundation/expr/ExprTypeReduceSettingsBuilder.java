// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprTypeReduceSettingsBuilder implements com.grafana.foundation.cog.Builder<ExprTypeReduceSettings> {
    protected final ExprTypeReduceSettings internal;
    
    public ExprTypeReduceSettingsBuilder() {
        this.internal = new ExprTypeReduceSettings();
    }
    public ExprTypeReduceSettingsBuilder mode(ExprTypeReduceSettingsMode mode) {
        this.internal.mode = mode;
        return this;
    }
    
    public ExprTypeReduceSettingsBuilder replaceWithValue(Double replaceWithValue) {
        this.internal.replaceWithValue = replaceWithValue;
        return this;
    }
    public ExprTypeReduceSettings build() {
        return this.internal;
    }
}
