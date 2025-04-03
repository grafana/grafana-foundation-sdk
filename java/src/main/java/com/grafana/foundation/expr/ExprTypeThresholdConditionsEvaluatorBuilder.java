// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;

public class ExprTypeThresholdConditionsEvaluatorBuilder implements com.grafana.foundation.cog.Builder<ExprTypeThresholdConditionsEvaluator> {
    protected final ExprTypeThresholdConditionsEvaluator internal;
    
    public ExprTypeThresholdConditionsEvaluatorBuilder() {
        this.internal = new ExprTypeThresholdConditionsEvaluator();
    }
    public ExprTypeThresholdConditionsEvaluatorBuilder params(List<Double> params) {
        this.internal.params = params;
        return this;
    }
    
    public ExprTypeThresholdConditionsEvaluatorBuilder type(ExprTypeThresholdConditionsEvaluatorType type) {
        this.internal.type = type;
        return this;
    }
    public ExprTypeThresholdConditionsEvaluator build() {
        return this.internal;
    }
}
