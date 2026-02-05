import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class AnnotationPanelFilterBuilder implements cog.Builder<dashboardv2beta1.AnnotationPanelFilter> {
    protected readonly internal: dashboardv2beta1.AnnotationPanelFilter;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AnnotationPanelFilter;
    exclude(exclude: boolean): this;
    ids(ids: number[]): this;
}
