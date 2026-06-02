<?php

namespace Grafana\Foundation\Common;

final class ResourceDimensionMode implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, ResourceDimensionMode>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function fixed(): self
    {
        if (!isset(self::$instances["Fixed"])) {
            self::$instances["Fixed"] = new self("fixed");
        }

        return self::$instances["Fixed"];
    }

    public static function field(): self
    {
        if (!isset(self::$instances["Field"])) {
            self::$instances["Field"] = new self("field");
        }

        return self::$instances["Field"];
    }

    public static function mapping(): self
    {
        if (!isset(self::$instances["Mapping"])) {
            self::$instances["Mapping"] = new self("mapping");
        }

        return self::$instances["Mapping"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "fixed") {
            return self::fixed();
        }

        if ($value === "field") {
            return self::field();
        }

        if ($value === "mapping") {
            return self::mapping();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum ResourceDimensionMode");
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

