import * as cog from '../cog';
import * as geomap from '../geomap';
export declare class ControlsOptionsBuilder implements cog.Builder<geomap.ControlsOptions> {
    protected readonly internal: geomap.ControlsOptions;
    constructor();
    /**
     * Builds the object.
     */
    build(): geomap.ControlsOptions;
    showZoom(showZoom: boolean): this;
    mouseWheelZoom(mouseWheelZoom: boolean): this;
    showAttribution(showAttribution: boolean): this;
    showScale(showScale: boolean): this;
    showDebug(showDebug: boolean): this;
    showMeasure(showMeasure: boolean): this;
}
