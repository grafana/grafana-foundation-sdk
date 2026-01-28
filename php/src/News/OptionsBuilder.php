<?php

namespace Grafana\Foundation\News;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\News\Options>
 */
class OptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\News\Options $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\News\Options();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\News\Options
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * empty/missing will default to grafana blog
     */
    public function feedUrl(string $feedUrl): static
    {
        $this->internal->feedUrl = $feedUrl;
    
        return $this;
    }

    public function showImage(bool $showImage): static
    {
        $this->internal->showImage = $showImage;
    
        return $this;
    }

}
