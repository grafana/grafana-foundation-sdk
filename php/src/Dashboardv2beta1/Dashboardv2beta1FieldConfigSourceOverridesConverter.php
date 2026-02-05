<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class Dashboardv2beta1FieldConfigSourceOverridesConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverrides $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1FieldConfigSourceOverridesBuilder())',
        ];
            
    
        {
    $buffer = 'matcher(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\MatcherConfig(id: '.\var_export($input->matcher->id, true).','.(($input->matcher->options !== null) ? 'options: '.\var_export($input->matcher->options, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->properties) >= 1) {
    
        
    $buffer = 'properties(';
        $tmparg0 = [];
        foreach ($input->properties as $arg1) {
        $tmppropertiesarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\DynamicConfigValue(id: '.\var_export($arg1->id, true).','.(($arg1->value !== null) ? 'value: '.\var_export($arg1->value, true).', ' : '').'))';
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

