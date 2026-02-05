<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class Dashboardv2beta1SpecialValueMapOptionsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1SpecialValueMapOptions $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1SpecialValueMapOptionsBuilder())',
        ];
            
    
        {
    $buffer = 'match(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\SpecialValueMatch::fromValue("'.$input->match.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            
    
        {
    $buffer = 'result(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\ValueMappingResultConverter::convert($input->result);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    

        return \implode("\n\t->", $calls);
    }
}

