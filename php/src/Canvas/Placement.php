<?php

namespace Grafana\Foundation\Canvas;

class Placement implements \JsonSerializable
{
    public ?float $top;

    public ?float $left;

    public ?float $right;

    public ?float $bottom;

    public ?float $width;

    public ?float $height;

    public ?float $rotation;

    /**
     * @param float|null $top
     * @param float|null $left
     * @param float|null $right
     * @param float|null $bottom
     * @param float|null $width
     * @param float|null $height
     * @param float|null $rotation
     */
    public function __construct(?float $top = null, ?float $left = null, ?float $right = null, ?float $bottom = null, ?float $width = null, ?float $height = null, ?float $rotation = null)
    {
        $this->top = $top;
        $this->left = $left;
        $this->right = $right;
        $this->bottom = $bottom;
        $this->width = $width;
        $this->height = $height;
        $this->rotation = $rotation;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{top?: float, left?: float, right?: float, bottom?: float, width?: float, height?: float, rotation?: float} $inputData */
        $data = $inputData;
        return new self(
            top: $data["top"] ?? null,
            left: $data["left"] ?? null,
            right: $data["right"] ?? null,
            bottom: $data["bottom"] ?? null,
            width: $data["width"] ?? null,
            height: $data["height"] ?? null,
            rotation: $data["rotation"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        if (isset($this->top)) {
            $data->top = $this->top;
        }
        if (isset($this->left)) {
            $data->left = $this->left;
        }
        if (isset($this->right)) {
            $data->right = $this->right;
        }
        if (isset($this->bottom)) {
            $data->bottom = $this->bottom;
        }
        if (isset($this->width)) {
            $data->width = $this->width;
        }
        if (isset($this->height)) {
            $data->height = $this->height;
        }
        if (isset($this->rotation)) {
            $data->rotation = $this->rotation;
        }
        return $data;
    }
}
