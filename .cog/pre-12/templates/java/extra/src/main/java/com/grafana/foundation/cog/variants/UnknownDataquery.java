package com.grafana.foundation.cog.variants;

import com.fasterxml.jackson.databind.annotation.JsonSerialize;

import java.util.HashMap;
import java.util.Map;

@JsonSerialize(using = UnknownDataquerySerializer.class)
public class UnknownDataquery implements Dataquery {
    public final Map<String, Object> genericFields = new HashMap<>();
    
    public String dataqueryName() {
        return "unknown";
    }
}
