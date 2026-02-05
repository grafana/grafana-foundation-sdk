<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ThresholdsConfigConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\ThresholdsConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\ThresholdsConfigBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\ThresholdsMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->steps) >= 1) {
    
        
    $buffer = 'steps(';
        $tmparg0 = [];
        foreach ($input->steps as $arg1) {
        $tmpstepsarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\Threshold(value: '.\var_export($arg1->value, true).',color: '.\var_export($arg1->color, true).',))';
        $tmparg0[] = $tmpstepsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

