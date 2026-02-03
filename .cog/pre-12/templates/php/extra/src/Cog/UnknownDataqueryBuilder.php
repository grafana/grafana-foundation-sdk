<?php

namespace {{ .Data.NamespaceRoot }}\Cog;

/**
 * @implements Builder<UnknownDataquery>
 */
final class UnknownDataqueryBuilder implements Builder
{
    protected UnknownDataquery $internal;

	public function __construct(?UnknownDataquery $object = null)
	{
    	$this->internal = $object ?: new UnknownDataquery([]);
	}

    /**
     * @return UnknownDataquery
     */
    public function build()
    {
        return $this->internal;
    }
}
