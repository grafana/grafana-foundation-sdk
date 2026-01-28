// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.text;


public class OptionsBuilder implements com.grafana.foundation.cog.Builder<Options> {
    protected final Options internal;
    
    public OptionsBuilder() {
        this.internal = new Options();
    }
    public OptionsBuilder mode(TextMode mode) {
        this.internal.mode = mode;
        return this;
    }
    
    public OptionsBuilder code(com.grafana.foundation.cog.Builder<CodeOptions> code) {
    CodeOptions codeResource = code.build();
        this.internal.code = codeResource;
        return this;
    }
    
    public OptionsBuilder content(String content) {
        this.internal.content = content;
        return this;
    }
    public Options build() {
        return this.internal;
    }
}
