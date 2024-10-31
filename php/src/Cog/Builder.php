<?php

namespace Grafana\Foundation\Cog;

/**
 * @template T
 */
interface Builder
{
    /**
     * @return T
     */
    public function build();
}
