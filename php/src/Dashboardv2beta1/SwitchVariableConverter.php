<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class SwitchVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\SwitchVariableBuilder('.\var_export($input->spec->name, true).'))',
        ];
            if ($input->spec->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->spec->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->current !== "" && $input->spec->current !== "false") {
    
        
    $buffer = 'current(';
        $arg0 =\var_export($input->spec->current, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->enabledValue !== "" && $input->spec->enabledValue !== "true") {
    
        
    $buffer = 'enabledValue(';
        $arg0 =\var_export($input->spec->enabledValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->disabledValue !== "" && $input->spec->disabledValue !== "false") {
    
        
    $buffer = 'disabledValue(';
        $arg0 =\var_export($input->spec->disabledValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->label !== null && $input->spec->label !== "") {
    
        
    $buffer = 'label(';
        $arg0 =\var_export($input->spec->label, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'hide(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\VariableHide::fromValue("'.$input->spec->hide.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->skipUrlSync !== false) {
    
        
    $buffer = 'skipUrlSync(';
        $arg0 =\var_export($input->spec->skipUrlSync, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->description !== null && $input->spec->description !== "") {
    
        
    $buffer = 'description(';
        $arg0 =\var_export($input->spec->description, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

