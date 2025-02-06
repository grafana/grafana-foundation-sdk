// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprTypeClassicConditionsConditionsBuilder implements com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditions> {
    protected final ExprTypeClassicConditionsConditions internal;
    
    public ExprTypeClassicConditionsConditionsBuilder() {
        this.internal = new ExprTypeClassicConditionsConditions();
    }
    public ExprTypeClassicConditionsConditionsBuilder evaluator(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsEvaluator> evaluator) {
        this.internal.evaluator = evaluator.build();
        return this;
    }
    
    public ExprTypeClassicConditionsConditionsBuilder operator(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsOperator> operator) {
        this.internal.operator = operator.build();
        return this;
    }
    
    public ExprTypeClassicConditionsConditionsBuilder query(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsQuery> query) {
        this.internal.query = query.build();
        return this;
    }
    
    public ExprTypeClassicConditionsConditionsBuilder reducer(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsReducer> reducer) {
        this.internal.reducer = reducer.build();
        return this;
    }
    public ExprTypeClassicConditionsConditions build() {
        return this.internal;
    }
}
