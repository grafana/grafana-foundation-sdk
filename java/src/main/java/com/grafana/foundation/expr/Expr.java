// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.annotation.JsonUnwrapped;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.ObjectWriter;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;

@JsonDeserialize(using = ExprDeserializer.class)
public class Expr implements com.grafana.foundation.cog.variants.Dataquery { 
    @JsonUnwrapped
    public TypeMath typeMath; 
    @JsonUnwrapped
    public TypeReduce typeReduce; 
    @JsonUnwrapped
    public TypeResample typeResample; 
    @JsonUnwrapped
    public TypeClassicConditions typeClassicConditions; 
    @JsonUnwrapped
    public TypeThreshold typeThreshold; 
    @JsonUnwrapped
    public TypeSql typeSql;
    
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
