<?php

namespace Grafana\Foundation\Testdata;

final class ScenarioConverter
{
    public static function convert(\Grafana\Foundation\Testdata\Scenario $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Testdata\ScenarioBuilder())',
        ];
            if ($input->id !== "") {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->stringInput !== "") {
    
        
    $buffer = 'stringInput(';
        $arg0 =\var_export($input->stringInput, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->description !== null && $input->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hideAliasField !== null) {
    
        
    $buffer = 'hideAliasField(';
        $arg0 =\var_export($input->hideAliasField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

