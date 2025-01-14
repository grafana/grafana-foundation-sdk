package com.grafana.foundation.cog.variants;
{{- $panelSchemas := .Context.SchemasForVariant "panelcfg" }}
{{- $dataquerySchemas := .Context.SchemasForVariant "dataquery" }}

import java.util.HashMap;
import java.util.Map;
{{ if .Data.Config.GenerateConverters -}}
import java.lang.reflect.InvocationTargetException;
import com.grafana.foundation.dashboard.Panel;
import com.grafana.foundation.cog.Converter;
{{- end }}

public class Registry {
    private static final Map<String, PanelConfig> panelRegistry = new HashMap<>();
    private static final Map<String, DataqueryConfig> dataqueryRegistry = new HashMap<>();
    
    static {
        {{- range $schema := $panelSchemas }}
        registerPanel("{{ $schema.Metadata.Identifier|lower }}", com.grafana.foundation.{{ $schema.Package }}.Options.class, {{ if $schema.HasObject "FieldConfig" }}com.grafana.foundation.{{ $schema.Package }}.FieldConfig.class{{ else }}null{{ end }}{{ if $.Data.Config.GenerateConverters }}, new com.grafana.foundation.{{ $schema.Package }}.PanelConverter.class{{ end }});
        {{- end }}

        {{- range $schema := $dataquerySchemas }}
        registerDataquery("{{ $schema.Metadata.Identifier|lower }}", com.grafana.foundation.{{ $schema.Package }}.{{ $schema.EntryPoint|formatObjectName }}.class{{ if $.Data.Config.GenerateConverters }}, new com.grafana.foundation.{{ $schema.Package }}.{{ $schema.EntryPoint|formatObjectName }}MapperConverter.class{{ end }});
        {{- end }}
    }

    public static void registerDataquery(String type, Class<? extends Dataquery> clazz{{ if .Data.Config.GenerateConverters }}, Converter<Dataquery> converter{{ end }}) {
        dataqueryRegistry.put(type, new DataqueryConfig(clazz{{ if .Data.Config.GenerateConverters }}, converter{{ end }}));
    }

    public static Class<? extends Dataquery> getDataquery(String type) {
        DataqueryConfig config = dataqueryRegistry.get(type);
        if (config != null) {
            return config.getDataquery();
        }

        return null;
    }
    
    public static void registerPanel(String type, Class<?> options, Class<?> fieldConfig{{ if .Data.Config.GenerateConverters }}, Converter<Panel> converter{{ end }}) {
        panelRegistry.put(type, new PanelConfig(options, fieldConfig{{ if .Data.Config.GenerateConverters }}, converter{{ end }}));
    }
    
    public static PanelConfig getPanel(String type) {
        return panelRegistry.get(type);
    }    

    {{- if .Data.Config.GenerateConverters }}
    public static String convertPanelToCode(Panel panel, String type) {
        return panelRegistry.get(type).getConverter().convert(panel);
    }
    
    public static String convertDataqueryToCode(Dataquery dataquery) {
        DataqueryConfig config = dataqueryRegistry.get(dataquery.dataqueryName());
        if (config != null) {
            return config.getConverter().convert(dataquery);
        }
        
        return null;
    }
    {{- end }}
}
