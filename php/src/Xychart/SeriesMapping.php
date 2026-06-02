<?php

namespace Grafana\Foundation\Xychart;

final class SeriesMapping implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, SeriesMapping>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function auto(): self
    {
        if (!isset(self::$instances["Auto"])) {
            self::$instances["Auto"] = new self("auto");
        }

        return self::$instances["Auto"];
    }

    public static function manual(): self
    {
        if (!isset(self::$instances["Manual"])) {
            self::$instances["Manual"] = new self("manual");
        }

        return self::$instances["Manual"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "auto") {
            return self::auto();
        }

        if ($value === "manual") {
            return self::manual();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum SeriesMapping");
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

