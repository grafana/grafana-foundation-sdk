// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.text;


public class CodeOptionsBuilder implements com.grafana.foundation.cog.Builder<CodeOptions> {
    protected final CodeOptions internal;
    
    public CodeOptionsBuilder() {
        this.internal = new CodeOptions();
    }
    public CodeOptionsBuilder language(CodeLanguage language) {
    this.internal.language = language;
        return this;
    }
    
    public CodeOptionsBuilder showLineNumbers(Boolean showLineNumbers) {
    this.internal.showLineNumbers = showLineNumbers;
        return this;
    }
    
    public CodeOptionsBuilder showMiniMap(Boolean showMiniMap) {
    this.internal.showMiniMap = showMiniMap;
        return this;
    }
    public CodeOptions build() {
        return this.internal;
    }
}
