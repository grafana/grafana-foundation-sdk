package com.grafana.foundation.cog.variants;
{{ if .Data.Config.GenerateConverters }}
import com.grafana.foundation.cog.Converter;
{{- end }}

public class DataqueryConfig {
    private final Class<? extends Dataquery> dataquery;
    {{ if .Data.Config.GenerateConverters }}private final Converter<Dataquery> converter;{{ end }}

    public DataqueryConfig(Class<? extends Dataquery> dataquery{{ if .Data.Config.GenerateConverters }}, Converter<Dataquery> converter{{ end }}) {
        this.dataquery = dataquery;
        {{ if .Data.Config.GenerateConverters }}this.converter = converter;{{ end }}
    }

    public Class<? extends Dataquery> getDataquery() {
        return dataquery;
    }
    
    {{- if .Data.Config.GenerateConverters }}
    public Converter<Dataquery> getConverter() {
        return converter;
    }
    {{- end }}
}
