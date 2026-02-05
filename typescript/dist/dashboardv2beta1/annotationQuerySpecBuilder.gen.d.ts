import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class AnnotationQuerySpecBuilder implements cog.Builder<dashboardv2beta1.AnnotationQuerySpec> {
    protected readonly internal: dashboardv2beta1.AnnotationQuerySpec;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AnnotationQuerySpec;
    query(query: cog.Builder<dashboardv2beta1.DataQueryKind>): this;
    enable(enable: boolean): this;
    hide(hide: boolean): this;
    iconColor(iconColor: string): this;
    name(name: string): this;
    builtIn(builtIn: boolean): this;
    filter(filter: cog.Builder<dashboardv2beta1.AnnotationPanelFilter>): this;
    placement(placement: "inControlsMenu"): this;
    legacyOptions(legacyOptions: Record<string, any>): this;
}
