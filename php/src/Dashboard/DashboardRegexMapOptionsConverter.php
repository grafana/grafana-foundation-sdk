<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardRegexMapOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DashboardRegexMapOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardRegexMapOptionsBuilder())',
        ];
            if ($input->pattern !== "") {
    
        
    $buffer = 'pattern(';
        $arg0 =\var_export($input->pattern, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'result(';
        $arg0 = \Grafana\Foundation\Dashboard\ValueMappingResultConverter::convert($input->result);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

