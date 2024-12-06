<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardSpecialValueMapOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptionsBuilder())',
        ];
            
    
        {
    $buffer = 'match(';
        $arg0 ='\Grafana\Foundation\Dashboard\SpecialValueMatch::fromValue("'.$input->match.'")';
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

