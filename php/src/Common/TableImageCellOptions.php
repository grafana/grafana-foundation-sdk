<?php

namespace Grafana\Foundation\Common;

/**
 * Json view cell options
 */
class TableImageCellOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Common\TableCellDisplayMode $type;

    public function __construct()
    {
        $this->type = \Grafana\Foundation\Common\TableCellDisplayMode::auto();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        return new self(
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->type = $this->type;
        return $data;
    }
}
