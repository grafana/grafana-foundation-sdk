<?php

namespace Grafana\Foundation\Nodegraph;

final class ZoomMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ZoomMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function cooperative(): self
    {
        if (!isset(self::$instances["cooperative"])) {
            self::$instances["cooperative"] = new self("cooperative");
        }

        return self::$instances["cooperative"];
    }

    public static function greedy(): self
    {
        if (!isset(self::$instances["greedy"])) {
            self::$instances["greedy"] = new self("greedy");
        }

        return self::$instances["greedy"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "cooperative") {
            return self::cooperative();
        }

        if ($value === "greedy") {
            return self::greedy();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ZoomMode");
    }

    public function jsonSerialize(): string
    {
        return $this->value;
    }

    public function __toString(): string
    {
        return $this->value;
    }
}

