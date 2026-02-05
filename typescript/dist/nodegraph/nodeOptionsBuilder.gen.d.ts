import * as cog from '../cog';
import * as nodegraph from '../nodegraph';
export declare class NodeOptionsBuilder implements cog.Builder<nodegraph.NodeOptions> {
    protected readonly internal: nodegraph.NodeOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): nodegraph.NodeOptions;
    mainStatUnit(mainStatUnit: string): this;
    secondaryStatUnit(secondaryStatUnit: string): this;
    arcs(arcs: cog.Builder<nodegraph.ArcOption>[]): this;
}
