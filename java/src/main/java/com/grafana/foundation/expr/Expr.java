// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.annotation.JsonInclude;

@JsonDeserialize(using = ExprDeserializer.class)
public class Expr implements com.grafana.foundation.cog.variants.Dataquery {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TypeMath typeMath;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TypeReduce typeReduce;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TypeResample typeResample;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TypeClassicConditions typeClassicConditions;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TypeThreshold typeThreshold;
    @JsonInclude(JsonInclude.Include.NON_NULL)
    @JsonUnwrapped
    protected TypeSql typeSql;
    protected Expr() {}
    public static Expr createTypeMath(TypeMath typeMath) {
        Expr expr = new Expr();
        expr.typeMath = typeMath;
        return expr;
    }
    public static Expr createTypeReduce(TypeReduce typeReduce) {
        Expr expr = new Expr();
        expr.typeReduce = typeReduce;
        return expr;
    }
    public static Expr createTypeResample(TypeResample typeResample) {
        Expr expr = new Expr();
        expr.typeResample = typeResample;
        return expr;
    }
    public static Expr createTypeClassicConditions(TypeClassicConditions typeClassicConditions) {
        Expr expr = new Expr();
        expr.typeClassicConditions = typeClassicConditions;
        return expr;
    }
    public static Expr createTypeThreshold(TypeThreshold typeThreshold) {
        Expr expr = new Expr();
        expr.typeThreshold = typeThreshold;
        return expr;
    }
    public static Expr createTypeSql(TypeSql typeSql) {
        Expr expr = new Expr();
        expr.typeSql = typeSql;
        return expr;
    }
    public String dataqueryName() {
        return "__expr__";
    }
    
    public String toJSON() throws JsonProcessingException {
        ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
        return ow.writeValueAsString(this);
    }

}
