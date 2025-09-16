// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class ExprSerializer extends JsonSerializer<Expr> {

    @Override
    public void serialize(Expr value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.typeMath != null) {
            gen.writeObject(value.typeMath);
        }
        else if (value.typeReduce != null) {
            gen.writeObject(value.typeReduce);
        }
        else if (value.typeResample != null) {
            gen.writeObject(value.typeResample);
        }
        else if (value.typeClassicConditions != null) {
            gen.writeObject(value.typeClassicConditions);
        }
        else if (value.typeThreshold != null) {
            gen.writeObject(value.typeThreshold);
        }
        else if (value.typeSql != null) {
            gen.writeObject(value.typeSql);
        }
    }
}
