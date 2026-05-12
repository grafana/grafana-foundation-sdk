---
title: <span class="badge object-type-enum"></span> AnnotationEventFieldSource
---
# <span class="badge object-type-enum"></span> AnnotationEventFieldSource

Annotation event field source. Defines how to obtain the value for an annotation event field.

- "field": Find the value with a matching key (default)

- "text": Write a constant string into the value

- "skip": Do not include the field

## Definition

```php
final class AnnotationEventFieldSource implements \JsonSerializable, \Stringable {
    /**
     * @var string
     */
    private $value;

    /**
     * @var array<string, AnnotationEventFieldSource>
     */
    private static $instances = [];

    private function __construct(string $value)
    {
        $this->value = $value;
    }

    public static function field(): self
    {
        if (!isset(self::$instances["field"])) {
            self::$instances["field"] = new self("field");
        }

        return self::$instances["field"];
    }

    public static function text(): self
    {
        if (!isset(self::$instances["text"])) {
            self::$instances["text"] = new self("text");
        }

        return self::$instances["text"];
    }

    public static function skip(): self
    {
        if (!isset(self::$instances["skip"])) {
            self::$instances["skip"] = new self("skip");
        }

        return self::$instances["skip"];
    }

    public static function fromValue(string $value): self
    {
        if ($value === "field") {
            return self::field();
        }

        if ($value === "text") {
            return self::text();
        }

        if ($value === "skip") {
            return self::skip();
        }

        throw new \UnexpectedValueException("Value '$value' is not part of the enum AnnotationEventFieldSource");
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
