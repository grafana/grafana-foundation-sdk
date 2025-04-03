<?php

namespace Grafana\Foundation\Cloudwatch;

final class QueryEditorPropertyConverter
{
    public static function convert(\Grafana\Foundation\Cloudwatch\QueryEditorProperty $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Cloudwatch\QueryEditorPropertyBuilder())',
        ];
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Cloudwatch\QueryEditorPropertyType::fromValue("'.$input->type.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
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

