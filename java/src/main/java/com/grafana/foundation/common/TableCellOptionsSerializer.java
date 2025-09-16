// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class TableCellOptionsSerializer extends JsonSerializer<TableCellOptions> {

    @Override
    public void serialize(TableCellOptions value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.tableAutoCellOptions != null) {
            gen.writeObject(value.tableAutoCellOptions);
        }
        else if (value.tableSparklineCellOptions != null) {
            gen.writeObject(value.tableSparklineCellOptions);
        }
        else if (value.tableBarGaugeCellOptions != null) {
            gen.writeObject(value.tableBarGaugeCellOptions);
        }
        else if (value.tableColoredBackgroundCellOptions != null) {
            gen.writeObject(value.tableColoredBackgroundCellOptions);
        }
        else if (value.tableColorTextCellOptions != null) {
            gen.writeObject(value.tableColorTextCellOptions);
        }
        else if (value.tableImageCellOptions != null) {
            gen.writeObject(value.tableImageCellOptions);
        }
        else if (value.tableJsonViewCellOptions != null) {
            gen.writeObject(value.tableJsonViewCellOptions);
        }
    }
}
