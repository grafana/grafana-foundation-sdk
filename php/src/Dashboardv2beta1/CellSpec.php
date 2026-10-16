<?php

namespace Grafana\Foundation\Dashboardv2beta1;

class CellSpec implements \JsonSerializable
{
    /**
     * @var \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentKind|\Grafana\Foundation\Dashboardv2beta1\CodeCellContentKind
     */
    public $content;

    /**
     * @param \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentKind|\Grafana\Foundation\Dashboardv2beta1\CodeCellContentKind|null $content
     */
    public function __construct( $content = null)
    {
        $this->content = $content ?: new \Grafana\Foundation\Dashboardv2beta1\MarkdownCellContentKind();
    }

    /**
     * @param array<string, mixed> $inputData
     */
    public static function fromArray(array $inputData): self
    {
        /** @var array{content?: mixed|mixed} $inputData */
        $data = $inputData;
        return new self(
            content: isset($data["content"]) ? (function($input) {
        \assert(is_array($input), 'expected disjunction value to be an array');
        /** @var array<string, mixed> $input */
        switch ($input["kind"]) {
        case "Code":
            return CodeCellContentKind::fromArray($input);
        case "Markdown":
            return MarkdownCellContentKind::fromArray($input);
        default:
            throw new \ValueError('can not parse disjunction from array');
    }
    })($data["content"]) : null,
        );
    }

    /**
     * @return mixed
     */
    public function jsonSerialize(): mixed
    {
        $data = new \stdClass;
        $data->content = $this->content;
        return $data;
    }
}
