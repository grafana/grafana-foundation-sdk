// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;

public class ExprTypeClassicConditionsConditionsEvaluatorBuilder implements com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsEvaluator> {
    protected final ExprTypeClassicConditionsConditionsEvaluator internal;
    
    public ExprTypeClassicConditionsConditionsEvaluatorBuilder() {
        this.internal = new ExprTypeClassicConditionsConditionsEvaluator();
    }
    public ExprTypeClassicConditionsConditionsEvaluatorBuilder params(List<Double> params) {
    this.internal.params = params;
        return this;
    }
    
    public ExprTypeClassicConditionsConditionsEvaluatorBuilder type(String type) {
    this.internal.type = type;
        return this;
    }
    public ExprTypeClassicConditionsConditionsEvaluator build() {
        return this.internal;
    }
}
