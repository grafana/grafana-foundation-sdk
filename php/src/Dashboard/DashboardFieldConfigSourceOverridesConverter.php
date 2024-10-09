<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardFieldConfigSourceOverridesConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverrides $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardFieldConfigSourceOverridesBuilder())',
        ];
            
    
        {
    $buffer = 'matcher(';
        $arg0 ='(new \Grafana\Foundation\Dashboard\MatcherConfig(id: '.\var_export($input->matcher->id, true).','.(($input->matcher->options !== null) ? 'options: '.\var_export($input->matcher->options, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->properties) >= 1) {
    
        
    $buffer = 'properties(';
        $tmparg0 = [];
        foreach ($input->properties as $arg1) {
        $tmppropertiesarg1 ='(new \Grafana\Foundation\Dashboard\DynamicConfigValue(id: '.\var_export($arg1->id, true).','.(($arg1->value !== null) ? 'value: '.\var_export($arg1->value, true).', ' : '').'))';
        $tmparg0[] = $tmppropertiesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

