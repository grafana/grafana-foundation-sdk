// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv0alpha1;

import com.fasterxml.jackson.core.JsonGenerator;
import com.fasterxml.jackson.databind.JsonSerializer;
import com.fasterxml.jackson.databind.SerializerProvider;

import java.io.IOException;

public class NotificationSettingsSerializer extends JsonSerializer<NotificationSettings> {

    @Override
    public void serialize(NotificationSettings value, JsonGenerator gen, SerializerProvider serializerProvider) throws IOException {
        if (value.simplifiedRouting != null) {
            gen.writeObject(value.simplifiedRouting);
        }
        else if (value.namedRoutingTree != null) {
            gen.writeObject(value.namedRoutingTree);
        }
    }
}
