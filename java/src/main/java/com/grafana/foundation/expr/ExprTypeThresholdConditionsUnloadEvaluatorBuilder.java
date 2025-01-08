// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;

public class ExprTypeThresholdConditionsUnloadEvaluatorBuilder implements com.grafana.foundation.cog.Builder<ExprTypeThresholdConditionsUnloadEvaluator> {
    protected final ExprTypeThresholdConditionsUnloadEvaluator internal;
    
    public ExprTypeThresholdConditionsUnloadEvaluatorBuilder() {
        this.internal = new ExprTypeThresholdConditionsUnloadEvaluator();
    }
    public ExprTypeThresholdConditionsUnloadEvaluatorBuilder params(List<Double> params) {
    this.internal.params = params;
        return this;
    }
    
    public ExprTypeThresholdConditionsUnloadEvaluatorBuilder type(TypeThresholdType type) {
    this.internal.type = type;
        return this;
    }
    public ExprTypeThresholdConditionsUnloadEvaluator build() {
        return this.internal;
    }
}
