import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class TargetBuilder implements cog.Builder<dashboardv2beta1.PanelQueryKind> {
    protected readonly internal: dashboardv2beta1.PanelQueryKind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.PanelQueryKind;
    query(query: cog.Builder<dashboardv2beta1.DataQueryKind>): this;
    refId(refId: string): this;
    hidden(hidden: boolean): this;
}
