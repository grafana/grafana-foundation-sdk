// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as text from '../text';

export class OptionsBuilder implements cog.Builder<text.Options> {
    protected readonly internal: text.Options;

    constructor() {
        this.internal = text.defaultOptions();
    }

    /**
     * Builds the object.
     */
    build(): text.Options {
        return this.internal;
    }

    mode(mode: text.TextMode): this {
        this.internal.mode = mode;
        return this;
    }

    code(code: cog.Builder<text.CodeOptions>): this {
        const codeResource = code.build();
        this.internal.code = codeResource;
        return this;
    }

    content(content: string): this {
        this.internal.content = content;
        return this;
    }
}

