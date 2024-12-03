<?php

namespace Grafana\Foundation\Bigquery;

final class QueryEditorPropertyConverter
{
    public static function convert(\Grafana\Foundation\Bigquery\QueryEditorProperty $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Bigquery\QueryEditorPropertyBuilder())',
        ];
            
    
        {
    $buffer = 'type(';
        $arg0 ='\Grafana\Foundation\Bigquery\QueryEditorPropertyType::fromValue("'.$input->type.'")';
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

