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
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

    
    public static class Builder implements com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditions> {
        protected final ExprTypeClassicConditionsConditions internal;
        
        public Builder() {
            this.internal = new ExprTypeClassicConditionsConditions();
        }
    public Builder evaluator(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsEvaluator> evaluator) {
    this.internal.evaluator = evaluator.build();
        return this;
    }
    
    public Builder operator(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsOperator> operator) {
    this.internal.operator = operator.build();
        return this;
    }
    
    public Builder query(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsQuery> query) {
    this.internal.query = query.build();
        return this;
    }
    
    public Builder reducer(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditionsReducer> reducer) {
    this.internal.reducer = reducer.build();
        return this;
    }
    public ExprTypeClassicConditionsConditions build() {
            return this.internal;
        }
    }
}
