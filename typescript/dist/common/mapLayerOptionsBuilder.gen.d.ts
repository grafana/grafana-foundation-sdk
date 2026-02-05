import * as cog from '../cog';
import * as common from '../common';
export declare class MapLayerOptionsBuilder implements cog.Builder<common.MapLayerOptions> {
    protected readonly internal: common.MapLayerOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): common.MapLayerOptions;
    type(type: string): this;
    name(name: string): this;
    config(config: any): this;
    location(location: cog.Builder<common.FrameGeometrySource>): this;
    filterData(filterData: any): this;
    opacity(opacity: number): this;
    tooltip(tooltip: boolean): this;
}
