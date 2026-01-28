<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class CustomVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\CustomVariableKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\CustomVariableBuilder('.\var_export($input->spec->name, true).'))',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\CustomVariableSpec(name: '.\var_export($input->spec->name, true).',query: '.\var_export($input->spec->query, true).',current: '.'(new \Grafana\Foundation\Dashboardv2beta1\VariableOption('.(($input->spec->current->selected !== null) ? 'selected: '.\var_export($input->spec->current->selected, true).', ' : '').'text: '.\var_export($input->spec->current->text, true).',value: '.\var_export($input->spec->current->value, true).',))'.',options: '.\var_export($input->spec->options, true).',multi: '.\var_export($input->spec->multi, true).',includeAll: '.\var_export($input->spec->includeAll, true).','.(($input->spec->allValue !== null) ? 'allValue: '.\var_export($input->spec->allValue, true).', ' : '').''.(($input->spec->label !== null) ? 'label: '.\var_export($input->spec->label, true).', ' : '').'hide: '.'\Grafana\Foundation\Dashboardv2beta1\VariableHide::fromValue("'.$input->spec->hide.'")'.',skipUrlSync: '.\var_export($input->spec->skipUrlSync, true).','.(($input->spec->description !== null) ? 'description: '.\var_export($input->spec->description, true).', ' : '').'allowCustomValue: '.\var_export($input->spec->allowCustomValue, true).','.(($input->spec->valuesFormat !== null) ? 'valuesFormat: '.'\Grafana\Foundation\Dashboardv2beta1\CustomVariableSpecValuesFormat::fromValue("'.$input->spec->valuesFormat.'")'.', ' : '').'))';
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
    
    
            if (count($input->spec->options) >= 1) {
    
        
    $buffer = 'options(';
        $tmparg0 = [];
        foreach ($input->spec->options as $arg1) {
        $tmpoptionsarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\VariableOption('.(($arg1->selected !== null) ? 'selected: '.\var_export($arg1->selected, true).', ' : '').'text: '.\var_export($arg1->text, true).',value: '.\var_export($arg1->value, true).',))';
        $tmparg0[] = $tmpoptionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->multi !== false) {
    
        
    $buffer = 'multi(';
        $arg0 =\var_export($input->spec->multi, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->includeAll !== false) {
    
        
    $buffer = 'includeAll(';
        $arg0 =\var_export($input->spec->includeAll, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->allValue !== null && $input->spec->allValue !== "") {
    
        
    $buffer = 'allValue(';
        $arg0 =\var_export($input->spec->allValue, true);
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
            if ($input->spec->allowCustomValue !== true) {
    
        
    $buffer = 'allowCustomValue(';
        $arg0 =\var_export($input->spec->allowCustomValue, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->valuesFormat !== null) {
    
        
    $buffer = 'valuesFormat(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\CustomVariableSpecValuesFormat::fromValue("'.$input->spec->valuesFormat.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

