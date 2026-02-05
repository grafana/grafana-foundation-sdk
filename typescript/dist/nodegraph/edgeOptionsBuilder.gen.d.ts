import * as cog from '../cog';
import * as nodegraph from '../nodegraph';
export declare class EdgeOptionsBuilder implements cog.Builder<nodegraph.EdgeOptions> {
    protected readonly internal: nodegraph.EdgeOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): nodegraph.EdgeOptions;
    mainStatUnit(mainStatUnit: string): this;
    secondaryStatUnit(secondaryStatUnit: string): this;
}
