import * as cog from '../cog';
import * as common from '../common';
export declare class LineStyleBuilder implements cog.Builder<common.LineStyle> {
    protected readonly internal: common.LineStyle;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.LineStyle;
    fill(fill: "solid" | "dash" | "dot" | "square"): this;
    dash(dash: number[]): this;
}
