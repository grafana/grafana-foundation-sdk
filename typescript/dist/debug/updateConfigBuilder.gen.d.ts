import * as cog from '../cog';
import * as debug from '../debug';
export declare class UpdateConfigBuilder implements cog.Builder<debug.UpdateConfig> {
    protected readonly internal: debug.UpdateConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): debug.UpdateConfig;
    render(render: boolean): this;
    dataChanged(dataChanged: boolean): this;
    schemaChanged(schemaChanged: boolean): this;
}
