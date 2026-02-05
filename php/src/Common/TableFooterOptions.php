<?php

namespace Grafana\Foundation\Common;

/**
 * Footer options
 */
class TableFooterOptions implements \JsonSerializable
{
    public bool $show;

    /**
     * actually 1 value
     * @var array<string>
     */
    public array $reducer;

    /**
     * @var array<string>|null
     */
    public ?array $fields;

    public ?bool $enablePagination;

    public ?bool $countRows;

    /**
     * @param bool|null $show
     * @param array<string>|null $reducer
     * @param array<string>|null $fields
     * @param bool|null $enablePagination
     * @param bool|null $countRows
     */
    public function __construct(?bool $show = null, ?array $reducer = null, ?array $fields = null, ?bool $enablePagination = null, ?bool $countRows = null)
    {
        $this->show = $show ?: false;
        $this->reducer = $reducer ?: [];
        $this->fields = $fields;
        $this->enablePagination = $enablePagination;
        $this->countRows = $countRows;
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{show?: bool, reducer?: array<string>, fields?: array<string>, enablePagination?: bool, countRows?: bool} $inputData */
        $data = $inputData;
        return new self(
            show: $data["show"] ?? null,
            reducer: $data["reducer"] ?? null,
            fields: $data["fields"] ?? null,
            enablePagination: $data["enablePagination"] ?? null,
            countRows: $data["countRows"] ?? null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->show = $this->show;
        $data->reducer = $this->reducer;
        if (isset($this->fields)) {
            $data->fields = $this->fields;
        }
        if (isset($this->enablePagination)) {
            $data->enablePagination = $this->enablePagination;
        }
        if (isset($this->countRows)) {
            $data->countRows = $this->countRows;
        }
        return $data;
    }
}
