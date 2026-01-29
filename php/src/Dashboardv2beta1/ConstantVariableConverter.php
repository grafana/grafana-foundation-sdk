<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class ConstantVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\ConstantVariableBuilder('.\var_export($input->spec->name, true).'))',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\ConstantVariableSpec(name: '.\var_export($input->spec->name, true).',query: '.\var_export($input->spec->query, true).',current: '.'(new \Grafana\Foundation\Dashboardv2beta1\VariableOption('.(($input->spec->current->selected !== null) ? 'selected: '.\var_export($input->spec->current->selected, true).', ' : '').'text: '.\var_export($input->spec->current->text, true).',value: '.\var_export($input->spec->current->value, true).',))'.','.(($input->spec->label !== null) ? 'label: '.\var_export($input->spec->label, true).', ' : '').'hide: '.'\Grafana\Foundation\Dashboardv2beta1\VariableHide::fromValue("'.$input->spec->hide.'")'.',skipUrlSync: '.\var_export($input->spec->skipUrlSync, true).','.(($input->spec->description !== null) ? 'description: '.\var_export($input->spec->description, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->spec->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->query !== "") {
    
        
    $buffer = 'query(';
        $arg0 =\var_export($input->spec->query, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'current(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\VariableOption('.(($input->spec->current->selected !== null) ? 'selected: '.\var_export($input->spec->current->selected, true).', ' : '').'text: '.\var_export($input->spec->current->text, true).',value: '.\var_export($input->spec->current->value, true).',))';
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

