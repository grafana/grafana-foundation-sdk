// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprTypeThresholdConditionsBuilder implements com.grafana.foundation.cog.Builder<ExprTypeThresholdConditions> {
    protected final ExprTypeThresholdConditions internal;
    
    public ExprTypeThresholdConditionsBuilder() {
        this.internal = new ExprTypeThresholdConditions();
    }
    public ExprTypeThresholdConditionsBuilder evaluator(com.grafana.foundation.cog.Builder<ExprTypeThresholdConditionsEvaluator> evaluator) {
        this.internal.evaluator = evaluator.build();
        return this;
    }
    
    public ExprTypeThresholdConditionsBuilder loadedDimensions(Object loadedDimensions) {
        this.internal.loadedDimensions = loadedDimensions;
        return this;
    }
    
    public ExprTypeThresholdConditionsBuilder unloadEvaluator(com.grafana.foundation.cog.Builder<ExprTypeThresholdConditionsUnloadEvaluator> unloadEvaluator) {
        this.internal.unloadEvaluator = unloadEvaluator.build();
        return this;
    }
    public ExprTypeThresholdConditions build() {
        return this.internal;
    }
}
