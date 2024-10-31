---
title: <span class="badge object-type-enum"></span> NodesQueryType
---
# <span class="badge object-type-enum"></span> NodesQueryType

## Definition

```php
final class NodesQueryType implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, NodesQueryType>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function random(): self
    {
        if (!isset(self::$instances["Random"])) {
            self::$instances["Random"] = new self("random");
        }

        return self::$instances["Random"];
    }

    public static function responseSmall(): self
    {
        if (!isset(self::$instances["ResponseSmall"])) {
            self::$instances["ResponseSmall"] = new self("response_small");
        }

        return self::$instances["ResponseSmall"];
    }

    public static function responseMedium(): self
    {
        if (!isset(self::$instances["ResponseMedium"])) {
            self::$instances["ResponseMedium"] = new self("response_medium");
        }

        return self::$instances["ResponseMedium"];
    }

    public static function randomEdges(): self
    {
        if (!isset(self::$instances["RandomEdges"])) {
            self::$instances["RandomEdges"] = new self("random edges");
        }

        return self::$instances["RandomEdges"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "random") {
            return self::random();
        }

        if ($value === "response_small") {
            return self::responseSmall();
        }

        if ($value === "response_medium") {
            return self::responseMedium();
        }

        if ($value === "random edges") {
            return self::randomEdges();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum NodesQueryType");
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

```
