<?php

namespace Grafana\Foundation\Canvas;

/**
 * @implements \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\CanvasConnection>
 */
class CanvasConnectionBuilder implements \Grafana\Foundation\Cog\Builder
{
    protected \Grafana\Foundation\Canvas\CanvasConnection $internal;

    public function __construct()
    {
    	$this->internal = new \Grafana\Foundation\Canvas\CanvasConnection();
    }

    /**
     * Builds the object.
     * @return \Grafana\Foundation\Canvas\CanvasConnection
     */
    public function build()
    {
        return $this->internal;
    }

    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\ConnectionCoordinates> $source
     */
    public function source(\Grafana\Foundation\Cog\Builder $source): static
    {
        $sourceResource = $source->build();
        $this->internal->source = $sourceResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\ConnectionCoordinates> $target
     */
    public function target(\Grafana\Foundation\Cog\Builder $target): static
    {
        $targetResource = $target->build();
        $this->internal->target = $targetResource;
    
        return $this;
    }
    public function targetName(string $targetName): static
    {
        $this->internal->targetName = $targetName;
    
        return $this;
    }
    public function path(\Grafana\Foundation\Canvas\ConnectionPath $path): static
    {
        $this->internal->path = $path;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ColorDimensionConfig> $color
     */
    public function color(\Grafana\Foundation\Cog\Builder $color): static
    {
        $colorResource = $color->build();
        $this->internal->color = $colorResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Common\ScaleDimensionConfig> $size
     */
    public function size(\Grafana\Foundation\Cog\Builder $size): static
    {
        $sizeResource = $size->build();
        $this->internal->size = $sizeResource;
    
        return $this;
    }
    /**
     * @param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\ConnectionCoordinates>> $vertices
     */
    public function vertices(array $vertices): static
    {
            $verticesResources = [];
            foreach ($vertices as $r1) {
                    $verticesResources[] = $r1->build();
            }
        $this->internal->vertices = $verticesResources;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\ConnectionCoordinates> $sourceOriginal
     */
    public function sourceOriginal(\Grafana\Foundation\Cog\Builder $sourceOriginal): static
    {
        $sourceOriginalResource = $sourceOriginal->build();
        $this->internal->sourceOriginal = $sourceOriginalResource;
    
        return $this;
    }
    /**
     * @param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Canvas\ConnectionCoordinates> $targetOriginal
     */
    public function targetOriginal(\Grafana\Foundation\Cog\Builder $targetOriginal): static
    {
        $targetOriginalResource = $targetOriginal->build();
        $this->internal->targetOriginal = $targetOriginalResource;
    
        return $this;
    }

}
