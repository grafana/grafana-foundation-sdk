package com.grafana.foundation.cog.variants;
{{ if .Data.Config.GenerateConverters }}
import com.grafana.foundation.dashboard.Panel;
import com.grafana.foundation.cog.Converter;
{{- end }}

public class PanelConfig {
    private final Class<?> optionsClass;
    private final Class<?> fieldConfigClass;
    {{- if .Data.Config.GenerateConverters }}
    private final Converter<Panel> converter;
    {{- end }}

    public PanelConfig(Class<?> optionsClass, Class<?> fieldConfigClass{{ if .Data.Config.GenerateConverters }}, Converter<Panel> converter{{ end }}) {
        this.optionsClass = optionsClass;
        this.fieldConfigClass = fieldConfigClass;
        {{ if .Data.Config.GenerateConverters -}}
        this.converter = converter;
        {{- end }}
    }

    public Class<?> getOptionsClass() {
        return optionsClass;
    }

    public Class<?> getFieldConfigClass() {
        return fieldConfigClass;
    }
    
    {{- if .Data.Config.GenerateConverters }}
    public Converter<Panel> getConverter() {
        return converter;
    }
    {{- end }}
}
