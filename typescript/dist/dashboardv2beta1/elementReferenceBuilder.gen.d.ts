import * as cog from '../cog';
import * as dashboardv2beta1 from '../dashboardv2beta1';
export declare class ElementReferenceBuilder implements cog.Builder<dashboardv2beta1.ElementReference> {
    protected readonly internal: dashboardv2beta1.ElementReference;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboardv2beta1.ElementReference;
    name(name: string): this;
}
