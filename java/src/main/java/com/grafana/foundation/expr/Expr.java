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
    public static Expr createTypeMath(com.grafana.foundation.cog.Builder<TypeMath> typeMath) {
        Expr expr = new Expr();
        expr.typeMath = typeMath.build();
        return expr;
    }
    public static Expr createTypeReduce(com.grafana.foundation.cog.Builder<TypeReduce> typeReduce) {
        Expr expr = new Expr();
        expr.typeReduce = typeReduce.build();
        return expr;
    }
    public static Expr createTypeResample(com.grafana.foundation.cog.Builder<TypeResample> typeResample) {
        Expr expr = new Expr();
        expr.typeResample = typeResample.build();
        return expr;
    }
    public static Expr createTypeClassicConditions(com.grafana.foundation.cog.Builder<TypeClassicConditions> typeClassicConditions) {
        Expr expr = new Expr();
        expr.typeClassicConditions = typeClassicConditions.build();
        return expr;
    }
    public static Expr createTypeThreshold(com.grafana.foundation.cog.Builder<TypeThreshold> typeThreshold) {
        Expr expr = new Expr();
        expr.typeThreshold = typeThreshold.build();
        return expr;
    }
    public static Expr createTypeSql(com.grafana.foundation.cog.Builder<TypeSql> typeSql) {
        Expr expr = new Expr();
        expr.typeSql = typeSql.build();
        return expr;
    }
    public String dataqueryName() {
        return "__expr__";
    }
    
    public String toJSON() throws JsonProcessingException {
        if (typeMath != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(typeMath);
        }
        if (typeReduce != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(typeReduce);
        }
        if (typeResample != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(typeResample);
        }
        if (typeClassicConditions != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(typeClassicConditions);
        }
        if (typeThreshold != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(typeThreshold);
        }
        if (typeSql != null) {
            ObjectWriter ow = new ObjectMapper().writer().withDefaultPrettyPrinter();
            return ow.writeValueAsString(typeSql);
        }
        
        return null;
    }

}
