<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class AdhocVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\AdhocVariableBuilder('.\var_export($input->spec->name, true).'))',
        ];
            if ($input->group !== "") {
    
        
    $buffer = 'group(';
        $arg0 =\var_export($input->group, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1AdhocVariableKindDatasourceConverter::convert($input->datasource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'spec(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\AdhocVariableSpec(name: '.\var_export($input->spec->name, true).',baseFilters: '.\var_export($input->spec->baseFilters, true).',filters: '.\var_export($input->spec->filters, true).',defaultKeys: '.\var_export($input->spec->defaultKeys, true).','.(($input->spec->label !== null) ? 'label: '.\var_export($input->spec->label, true).', ' : '').'hide: '.'\Grafana\Foundation\Dashboardv2beta1\VariableHide::fromValue("'.$input->spec->hide.'")'.',skipUrlSync: '.\var_export($input->spec->skipUrlSync, true).','.(($input->spec->description !== null) ? 'description: '.\var_export($input->spec->description, true).', ' : '').'allowCustomValue: '.\var_export($input->spec->allowCustomValue, true).',))';
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
            if (count($input->spec->baseFilters) >= 1) {
    
        
    $buffer = 'baseFilters(';
        $tmparg0 = [];
        foreach ($input->spec->baseFilters as $arg1) {
        $tmpbaseFiltersarg1 = \Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabelsConverter::convert($arg1);
        $tmparg0[] = $tmpbaseFiltersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->spec->filters) >= 1) {
    
        
    $buffer = 'filters(';
        $tmparg0 = [];
        foreach ($input->spec->filters as $arg1) {
        $tmpfiltersarg1 = \Grafana\Foundation\Dashboardv2beta1\AdHocFilterWithLabelsConverter::convert($arg1);
        $tmparg0[] = $tmpfiltersarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->spec->defaultKeys) >= 1) {
    
        
    $buffer = 'defaultKeys(';
        $tmparg0 = [];
        foreach ($input->spec->defaultKeys as $arg1) {
        $tmpdefaultKeysarg1 = \Grafana\Foundation\Dashboardv2beta1\MetricFindValueConverter::convert($arg1);
        $tmparg0[] = $tmpdefaultKeysarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
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

        return \implode("\n\t->", $calls);
    }
}

