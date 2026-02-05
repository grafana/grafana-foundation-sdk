import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class FieldColorBuilder implements cog.Builder<dashboardv2beta1.FieldColor> {
    protected readonly internal: dashboardv2beta1.FieldColor;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.FieldColor;
    mode(mode: dashboardv2beta1.FieldColorModeId): this;
    fixedColor(fixedColor: string): this;
    seriesBy(seriesBy: dashboardv2beta1.FieldColorSeriesByMode): this;
}
