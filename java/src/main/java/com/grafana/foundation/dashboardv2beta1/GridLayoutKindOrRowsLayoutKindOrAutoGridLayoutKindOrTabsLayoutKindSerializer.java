// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindSerializer extends JsonSerializer<GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind> {

    @Override
    public void serialize(GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.gridLayoutKind != null) {
            gen.writeObject(value.gridLayoutKind);
        }
        else if (value.rowsLayoutKind != null) {
            gen.writeObject(value.rowsLayoutKind);
        }
        else if (value.autoGridLayoutKind != null) {
            gen.writeObject(value.autoGridLayoutKind);
        }
        else if (value.tabsLayoutKind != null) {
            gen.writeObject(value.tabsLayoutKind);
        }
    }
}
