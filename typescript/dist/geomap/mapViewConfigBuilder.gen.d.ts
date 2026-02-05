import * as cog from '../cog';
import * as geomap from '../geomap';
export declare class MapViewConfigBuilder implements cog.Builder<geomap.MapViewConfig> {
    protected readonly internal: geomap.MapViewConfig;
    constructor();
    /**
     * Builds the object.
     */
    build(): geomap.MapViewConfig;
    id(id: string): this;
    lat(lat: number): this;
    lon(lon: number): this;
    zoom(zoom: number): this;
    minZoom(minZoom: number): this;
    maxZoom(maxZoom: number): this;
    padding(padding: number): this;
    allLayers(allLayers: boolean): this;
    lastOnly(lastOnly: boolean): this;
    layer(layer: string): this;
    shared(shared: boolean): this;
}
