import * as cog from '../cog';
import * as dashboard from '../dashboard';
export declare class TimeOptionBuilder implements cog.Builder<dashboard.TimeOption> {
    protected readonly internal: dashboard.TimeOption;
    constructor();
    /**
     * Builds the object.
     */
    build(): dashboard.TimeOption;
    display(display: string): this;
    from(from: string): this;
    to(to: string): this;
}
