// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.annotation.JsonInclude;

public class ExprTypeClassicConditionsConditions {
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("evaluator")
    public ExprTypeClassicConditionsConditionsEvaluator evaluator;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("operator")
    public ExprTypeClassicConditionsConditionsOperator operator;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("query")
    public ExprTypeClassicConditionsConditionsQuery query;
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    @JsonProperty("reducer")
    public ExprTypeClassicConditionsConditionsReducer reducer;
    public ExprTypeClassicConditionsConditions() {
        this.evaluator = new com.grafana.foundation.expr.ExprTypeClassicConditionsConditionsEvaluatorBuilder().build();
        this.operator = new com.grafana.foundation.expr.ExprTypeClassicConditionsConditionsOperatorBuilder().build();
        this.query = new com.grafana.foundation.expr.ExprTypeClassicConditionsConditionsQueryBuilder().build();
        this.reducer = new com.grafana.foundation.expr.ExprTypeClassicConditionsConditionsReducerBuilder().build();
    }
    public ExprTypeClassicConditionsConditions(ExprTypeClassicConditionsConditionsEvaluator evaluator,ExprTypeClassicConditionsConditionsOperator operator,ExprTypeClassicConditionsConditionsQuery query,ExprTypeClassicConditionsConditionsReducer reducer) {
        this.evaluator = evaluator;
        this.operator = operator;
        this.query = query;
        this.reducer = reducer;
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
