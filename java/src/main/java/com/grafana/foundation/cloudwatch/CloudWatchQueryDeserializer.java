// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import com.fasterxml.jackson.core.JsonParser;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationContext;
import com.fasterxml.jackson.databind.JsonDeserializer;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.io.IOException;

public class CloudWatchQueryDeserializer extends JsonDeserializer<CloudWatchQuery> {

    @Override
    public CloudWatchQuery deserialize(JsonParser jp, DeserializationContext cxt) throws IOException {
        ObjectMapper mapper = (ObjectMapper) jp.getCodec();
        JsonNode root = mapper.readTree(jp);
        
        CloudWatchQuery cloudWatchQuery = new CloudWatchQuery();
        if (!root.has("queryMode")) {
            throw new IOException("Cannot find discriminator for CloudWatchQuery");
        }
        String discriminator = root.get("queryMode").asText();  
        
        switch (discriminator) {
        case "Annotations":
            cloudWatchQuery.cloudWatchAnnotationQuery = mapper.convertValue(root, CloudWatchAnnotationQuery.class);
            break;
        case "Logs":
            cloudWatchQuery.cloudWatchLogsQuery = mapper.convertValue(root, CloudWatchLogsQuery.class);
            break;
        case "Metrics":
            cloudWatchQuery.cloudWatchMetricsQuery = mapper.convertValue(root, CloudWatchMetricsQuery.class);
            break;
        }
        
        return cloudWatchQuery;
    }
}
