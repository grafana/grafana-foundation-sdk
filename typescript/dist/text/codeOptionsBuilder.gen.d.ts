import * as cog from '../cog';
import * as text from '../text';
export declare class CodeOptionsBuilder implements cog.Builder<text.CodeOptions> {
    protected readonly internal: text.CodeOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): text.CodeOptions;
    language(language: text.CodeLanguage): this;
    showLineNumbers(showLineNumbers: boolean): this;
    showMiniMap(showMiniMap: boolean): this;
}
