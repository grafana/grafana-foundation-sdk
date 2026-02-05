import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class AnnotationQueryBuilder implements cog.Builder<dashboardv2beta1.AnnotationQueryKind> {
    protected readonly internal: dashboardv2beta1.AnnotationQueryKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.AnnotationQueryKind;
    spec(spec: cog.Builder<dashboardv2beta1.AnnotationQuerySpec>): this;
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
