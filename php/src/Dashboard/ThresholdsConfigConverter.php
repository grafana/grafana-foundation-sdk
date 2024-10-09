<?php

namespace Grafana\Foundation\Dashboard;

final class ThresholdsConfigConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\ThresholdsConfig $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\ThresholdsConfigBuilder())',
        ];
            
    
        {
    $buffer = 'mode(';
        $arg0 ='\Grafana\Foundation\Dashboard\ThresholdsMode::fromValue("'.$input->mode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->steps) >= 1) {
    
        
    $buffer = 'steps(';
        $tmparg0 = [];
        foreach ($input->steps as $arg1) {
        $tmpstepsarg1 ='(new \Grafana\Foundation\Dashboard\Threshold('.(($arg1->value !== null) ? 'value: '.\var_export($arg1->value, true).', ' : '').'color: '.\var_export($arg1->color, true).',))';
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

