// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class RequestDeserializer extends JsonDeserializer<Request> {

    @Override
    public Request deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        Request request = new Request();
        if (!root.has("queryMode")) {
            throw new IOException("Cannot find discriminator for Request");
        }
        String discriminator = root.get("queryMode").asText();  
        
        switch (discriminator) {
        case "Annotations":
            request.annotationQuery = mapper.convertValue(root, AnnotationQuery.class);
            break;
        case "Logs":
            request.logsQuery = mapper.convertValue(root, LogsQuery.class);
            break;
        case "Metrics":
            request.metricsQuery = mapper.convertValue(root, MetricsQuery.class);
            break;
        }
        
        return request;
    }
}
