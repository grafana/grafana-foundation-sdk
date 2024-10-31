// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as text from '../text';

export class CodeOptionsBuilder implements cog.Builder<text.CodeOptions> {
    protected readonly internal: text.CodeOptions;

    constructor() {
        this.internal = text.defaultCodeOptions();
    }

    /**
     * Builds the object.
     */
    build(): text.CodeOptions {
        return this.internal;
    }

    // The language passed to monaco code editor
    language(language: text.CodeLanguage): this {
        this.internal.language = language;
        return this;
    }

    showLineNumbers(showLineNumbers: boolean): this {
        this.internal.showLineNumbers = showLineNumbers;
        return this;
    }

    showMiniMap(showMiniMap: boolean): this {
        this.internal.showMiniMap = showMiniMap;
        return this;
    }
}
