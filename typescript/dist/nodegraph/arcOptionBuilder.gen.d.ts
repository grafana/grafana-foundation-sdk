import * as cog from '../cog';
import * as nodegraph from '../nodegraph';
export declare class ArcOptionBuilder implements cog.Builder<nodegraph.ArcOption> {
    protected readonly internal: nodegraph.ArcOption;
    constructor();
    /**
     * Builds the object.
     */
    build(): nodegraph.ArcOption;
    field(field: string): this;
    color(color: string): this;
}
