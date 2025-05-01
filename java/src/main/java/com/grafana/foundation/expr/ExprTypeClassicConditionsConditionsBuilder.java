// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprTypeClassicConditionsConditionsBuilder implements com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditions> {
    protected final ExprTypeClassicConditionsConditions internal;
    
    public ExprTypeClassicConditionsConditionsBuilder() {
        this.internal = new ExprTypeClassicConditionsConditions();
    }
    public ExprTypeClassicConditionsConditionsBuilder evaluator(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsEvaluator> evaluator) {
    ExprTypeClassicConditionsConditionsEvaluator evaluatorResource = evaluator.build();
        this.internal.evaluator = evaluatorResource;
        return this;
    }
    
    public ExprTypeClassicConditionsConditionsBuilder operator(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsOperator> operator) {
    ExprTypeClassicConditionsConditionsOperator operatorResource = operator.build();
        this.internal.operator = operatorResource;
        return this;
    }
    
    public ExprTypeClassicConditionsConditionsBuilder query(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsQuery> query) {
    ExprTypeClassicConditionsConditionsQuery queryResource = query.build();
        this.internal.query = queryResource;
        return this;
    }
    
    public ExprTypeClassicConditionsConditionsBuilder reducer(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsReducer> reducer) {
    ExprTypeClassicConditionsConditionsReducer reducerResource = reducer.build();
        this.internal.reducer = reducerResource;
        return this;
    }
    public ExprTypeClassicConditionsConditions build() {
        return this.internal;
    }
}
