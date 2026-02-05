import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class KindBuilder implements cog.Builder<dashboardv2beta1.Kind> {
    protected readonly internal: dashboardv2beta1.Kind;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.Kind;
    kind(kind: string): this;
    spec(spec: any): this;
    metadata(metadata: any): this;
}
