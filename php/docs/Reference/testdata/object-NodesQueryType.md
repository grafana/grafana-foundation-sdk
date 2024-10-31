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

    public static function response(): self
    {
        if (!isset(self::$instances["Response"])) {
            self::$instances["Response"] = new self("response");
        }

        return self::$instances["Response"];
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

        if ($value === "response") {
            return self::response();
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
