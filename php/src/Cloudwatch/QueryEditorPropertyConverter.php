<?php

namespace Grafana\Foundation\Cloudwatch;

final class QueryEditorPropertyConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\QueryEditorProperty $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\QueryEditorPropertyBuilder())',
        ];
            if ($input->name !== null && $input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

