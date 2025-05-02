<?php

namespace App\Custom;

use Grafana\Foundation\Cog;
use Grafana\Foundation\Dashboard;

/**
 * @implements Cog\Builder<Dashboard\Panel>
 */
class PanelBuilder extends Dashboard\PanelBuilder implements Cog\Builder
{
    public function __construct()
    {
        parent::__construct();
        $this->internal->type = 'custom-panel';
    }

    public function makeBeautiful(): static
    {
        if ($this->internal->options === null) {
            $this->internal->options = new Options();
        }
        assert($this->internal->options instanceof Options);
        $this->internal->options->makeBeautiful = true;
    
        return $this;
    }
}
