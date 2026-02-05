import * as cog from '../cog';
import * as common from '../common';
export declare class OptionsWithTextFormattingBuilder implements cog.Builder<common.OptionsWithTextFormatting> {
    protected readonly internal: common.OptionsWithTextFormatting;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.OptionsWithTextFormatting;
    text(text: cog.Builder<common.VizTextDisplayOptions>): this;
}
