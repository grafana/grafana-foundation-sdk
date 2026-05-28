<?php

namespace Grafana\Foundation\Dashboardv2;

final class Dashboardv2FieldConfigSourceOverridesConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\Dashboardv2FieldConfigSourceOverrides $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\Dashboardv2FieldConfigSourceOverridesBuilder())',
        ];
            if ($input->systemRef !== null && $input->systemRef !== "") {
    
        
    $buffer = 'systemRef(';
        $arg0 =\var_export($input->systemRef, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'matcher(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2\MatcherConfig(id: '.\var_export($input->matcher->id, true).','.(($input->matcher->scope !== null) ? 'scope: '.'\Grafana\Foundation\Dashboardv2\MatcherScope::fromValue("'.$input->matcher->scope.'")'.', ' : '').''.(($input->matcher->options !== null) ? 'options: '.\var_export($input->matcher->options, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->properties) >= 1) {
    
        
    $buffer = 'properties(';
        $tmparg0 = [];
        foreach ($input->properties as $arg1) {
        $tmppropertiesarg1 ='(new \Grafana\Foundation\Dashboardv2\DynamicConfigValue(id: '.\var_export($arg1->id, true).','.(($arg1->value !== null) ? 'value: '.\var_export($arg1->value, true).', ' : '').'))';
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

