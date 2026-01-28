<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class SwitchVariableSpecConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\SwitchVariableSpecBuilder())',
        ];
            if ($input->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->current !== "" && $input->current !== "false") {
    
        
    $buffer = 'current(';
        $arg0 =\var_export($input->current, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->enabledValue !== "" && $input->enabledValue !== "true") {
    
        
    $buffer = 'enabledValue(';
        $arg0 =\var_export($input->enabledValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->disabledValue !== "" && $input->disabledValue !== "false") {
    
        
    $buffer = 'disabledValue(';
        $arg0 =\var_export($input->disabledValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->label !== null && $input->label !== "") {
    
        
    $buffer = 'label(';
        $arg0 =\var_export($input->label, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'hide(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\VariableHide::fromValue("'.$input->hide.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->skipUrlSync !== false) {
    
        
    $buffer = 'skipUrlSync(';
        $arg0 =\var_export($input->skipUrlSync, true);
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

        return \implode("\n\t->", $calls);
    }
}

