<?php

namespace Grafana\Foundation\Dashboard;

final class DataSourceRefConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DataSourceRef $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DataSourceRefBuilder())',
        ];
            if ($input->type !== null && $input->type !== "") {
    
        
    $buffer = 'type(';
        $arg0 =\var_export($input->type, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->uid !== null && $input->uid !== "") {
    
        
    $buffer = 'uid(';
        $arg0 =\var_export($input->uid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

