// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprTypeThresholdConditionsBuilder implements com.grafana.foundation.cog.Builder<ExprTypeThresholdConditions> {
    protected final ExprTypeThresholdConditions internal;
    
    public ExprTypeThresholdConditionsBuilder() {
        this.internal = new ExprTypeThresholdConditions();
    }
    public ExprTypeThresholdConditionsBuilder evaluator(com.grafana.foundation.cog.Builder<ExprTypeThresholdConditionsEvaluator> evaluator) {
    ExprTypeThresholdConditionsEvaluator evaluatorResource = evaluator.build();
        this.internal.evaluator = evaluatorResource;
        return this;
    }
    
    public ExprTypeThresholdConditionsBuilder loadedDimensions(Object loadedDimensions) {
        this.internal.loadedDimensions = loadedDimensions;
        return this;
    }
    
    public ExprTypeThresholdConditionsBuilder unloadEvaluator(com.grafana.foundation.cog.Builder<ExprTypeThresholdConditionsUnloadEvaluator> unloadEvaluator) {
    ExprTypeThresholdConditionsUnloadEvaluator unloadEvaluatorResource = unloadEvaluator.build();
        this.internal.unloadEvaluator = unloadEvaluatorResource;
        return this;
    }
    public ExprTypeThresholdConditions build() {
        return this.internal;
    }
}
