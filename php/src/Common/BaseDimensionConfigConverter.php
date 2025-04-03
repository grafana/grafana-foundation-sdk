<?php

namespace Grafana\Foundation\Common;

final class BaseDimensionConfigConverter
{
    public static function convert(\Grafana\Foundation\Common\BaseDimensionConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Common\BaseDimensionConfigBuilder())',
        ];
            if ($input->field !== null && $input->field !== "") {
    
        
    $buffer = 'field(';
        $arg0 =\var_export($input->field, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

