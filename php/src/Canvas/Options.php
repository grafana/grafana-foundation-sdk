<?php

namespace Grafana\Foundation\Canvas;

class Options implements \JsonSerializable
{
    /**
     * Enable inline editing
     */
    public bool $inlineEditing;

    /**
     * Show all available element types
     */
    public bool $showAdvancedTypes;

    /**
     * Enable pan and zoom
     */
    public bool $panZoom;

    /**
     * Enable infinite pan
     */
    public bool $infinitePan;

    /**
     * The root element of canvas (frame), where all canvas elements are nested
     * TODO: Figure out how to define a default value for this
     */
    public \Grafana\Foundation\Canvas\CanvasOptionsRoot $root;

    /**
     * @param bool|null $inlineEditing
     * @param bool|null $showAdvancedTypes
     * @param bool|null $panZoom
     * @param bool|null $infinitePan
     * @param \Grafana\Foundation\Canvas\CanvasOptionsRoot|null $root
     */
    public function __construct(?bool $inlineEditing = null, ?bool $showAdvancedTypes = null, ?bool $panZoom = null, ?bool $infinitePan = null, ?\Grafana\Foundation\Canvas\CanvasOptionsRoot $root = null)
    {
        $this->inlineEditing = $inlineEditing ?: true;
        $this->showAdvancedTypes = $showAdvancedTypes ?: true;
        $this->panZoom = $panZoom ?: true;
        $this->infinitePan = $infinitePan ?: true;
        $this->root = $root ?: new \Grafana\Foundation\Canvas\CanvasOptionsRoot();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{inlineEditing?: bool, showAdvancedTypes?: bool, panZoom?: bool, infinitePan?: bool, root?: mixed} $inputData */
        $data = $inputData;
        return new self(
            inlineEditing: $data["inlineEditing"] ?? null,
            showAdvancedTypes: $data["showAdvancedTypes"] ?? null,
            panZoom: $data["panZoom"] ?? null,
            infinitePan: $data["infinitePan"] ?? null,
            root: isset($data["root"]) ? (function($input) {
    	/** @var array{name?: string, type?: string, elements?: array<mixed>} */
    $val = $input;
    	return \Grafana\Foundation\Canvas\CanvasOptionsRoot::fromArray($val);
    })($data["root"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->inlineEditing = $this->inlineEditing;
        $data->showAdvancedTypes = $this->showAdvancedTypes;
        $data->panZoom = $this->panZoom;
        $data->infinitePan = $this->infinitePan;
        $data->root = $this->root;
        return $data;
    }
}
