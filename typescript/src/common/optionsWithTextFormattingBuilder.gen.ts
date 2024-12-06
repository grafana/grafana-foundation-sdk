// Code generated - EDITING IS FUTILE. DO NOT EDIT.

import * as cog from '../cog';
import * as common from '../common';

// TODO docs
export class OptionsWithTextFormattingBuilder implements cog.Builder<common.OptionsWithTextFormatting> {
    protected readonly internal: common.OptionsWithTextFormatting;

    constructor() {
        this.internal = common.defaultOptionsWithTextFormatting();
    }

    /**
     * Builds the object.
     */
    build(): common.OptionsWithTextFormatting {
        return this.internal;
    }

    text(text: cog.Builder<common.VizTextDisplayOptions>): this {
        const textResource = text.build();
        this.internal.text = textResource;
        return this;
    }
}
