<?php

namespace Grafana\Foundation\Text;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Text\CodeOptions>
 */
class CodeOptionsBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Text\CodeOptions $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Text\CodeOptions();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Text\CodeOptions
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * The language passed to monaco code editor
     */
    public function language(\Grafana\Foundation\Text\CodeLanguage $language): static
    {
        $this->internal->language = $language;
    
        return $this;
    }
    public function showLineNumbers(bool $showLineNumbers): static
    {
        $this->internal->showLineNumbers = $showLineNumbers;
    
        return $this;
    }
    public function showMiniMap(bool $showMiniMap): static
    {
        $this->internal->showMiniMap = $showMiniMap;
    
        return $this;
    }

}
