<?php

namespace Grafana\Foundation\Dashboardv2;

final class DatasourceVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2\DatasourceVariableKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2\DatasourceVariableBuilder('.\var_export($input->spec->name, true).'))',
        ];
            if ($input->spec->name !== "") {
    
        
    $buffer = 'name(';
        $arg0 =\var_export($input->spec->name, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->pluginId !== "") {
    
        
    $buffer = 'pluginId(';
        $arg0 =\var_export($input->spec->pluginId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'refresh(';
        $arg0 ='\Grafana\Foundation\Dashboardv2\VariableRefresh::fromValue("'.$input->spec->refresh.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->regex !== "") {
    
        
    $buffer = 'regex(';
        $arg0 =\var_export($input->spec->regex, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'current(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2\VariableOption('.(($input->spec->current->selected !== null) ? 'selected: '.\var_export($input->spec->current->selected, true).', ' : '').'text: '.\var_export($input->spec->current->text, true).',value: '.\var_export($input->spec->current->value, true).','.(($input->spec->current->properties !== null) ? 'properties: '.\var_export($input->spec->current->properties, true).', ' : '').'))';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if (count($input->spec->options) >= 1) {
    
        
    $buffer = 'options(';
        $tmparg0 = [];
        foreach ($input->spec->options as $arg1) {
        $tmpoptionsarg1 ='(new \Grafana\Foundation\Dashboardv2\VariableOption('.(($arg1->selected !== null) ? 'selected: '.\var_export($arg1->selected, true).', ' : '').'text: '.\var_export($arg1->text, true).',value: '.\var_export($arg1->value, true).','.(($arg1->properties !== null) ? 'properties: '.\var_export($arg1->properties, true).', ' : '').'))';
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
        $arg0 ='\Grafana\Foundation\Dashboardv2\VariableHide::fromValue("'.$input->spec->hide.'")';
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
            if ($input->spec->origin !== null) {
    
        
    $buffer = 'origin(';
        $arg0 = \Grafana\Foundation\Dashboardv2\ControlSourceRefConverter::convert($input->spec->origin);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

