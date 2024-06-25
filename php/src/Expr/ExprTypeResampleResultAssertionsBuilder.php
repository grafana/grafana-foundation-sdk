<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeResampleResultAssertions>
 */
class ExprTypeResampleResultAssertionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeResampleResultAssertions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeResampleResultAssertions();
    }

    /**
     * @return \Grafana\Foundation\Expr\ExprTypeResampleResultAssertions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Maximum frame count
     */
    public function maxFrames(int $maxFrames): static
    {
        $this->internal->maxFrames = $maxFrames;
    
        return $this;
    }
    /**
     * Type asserts that the frame matches a known type structure.
     * Possible enum values:
     *  - `""` 
     *  - `"timeseries-wide"` 
     *  - `"timeseries-long"` 
     *  - `"timeseries-many"` 
     *  - `"timeseries-multi"` 
     *  - `"directory-listing"` 
     *  - `"table"` 
     *  - `"numeric-wide"` 
     *  - `"numeric-multi"` 
     *  - `"numeric-long"` 
     *  - `"log-lines"` 
     */
    public function type(\Grafana\Foundation\Expr\TypeResampleType $type): static
    {
        $this->internal->type = $type;
    
        return $this;
    }
    /**
     * TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
     * contract documentation https://grafana.github.io/dataplane/contract/.
     * @param array<int> $typeVersion
     */
    public function typeVersion(array $typeVersion): static
    {
        $this->internal->typeVersion = $typeVersion;
    
        return $this;
    }

}
