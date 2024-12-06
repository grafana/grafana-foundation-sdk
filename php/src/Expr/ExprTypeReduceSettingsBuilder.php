<?php

namespace Grafana\Foundation\Expr;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Expr\ExprTypeReduceSettings>
 */
class ExprTypeReduceSettingsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Expr\ExprTypeReduceSettings $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Expr\ExprTypeReduceSettings();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Expr\ExprTypeReduceSettings
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * Non-number reduce behavior
     * Possible enum values:
     *  - `"dropNN"` Drop non-numbers
     *  - `"replaceNN"` Replace non-numbers
     */
    public function mode(\Grafana\Foundation\Expr\TypeReduceMode $mode): static
    {
        $this->internal->mode = $mode;
    
        return $this;
    }
    /**
     * Only valid when mode is replace
     */
    public function replaceWithValue(float $replaceWithValue): static
    {
        $this->internal->replaceWithValue = $replaceWithValue;
    
        return $this;
    }

}
