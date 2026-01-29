<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class QueryVariableConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\QueryVariableKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\QueryVariableBuilder('.\var_export($input->spec->name, true).'))',
        ];
            
    
        {
    $buffer = 'spec(';
        $arg0 ='(new \Grafana\Foundation\Dashboardv2beta1\QueryVariableSpec(name: '.\var_export($input->spec->name, true).',current: '.'(new \Grafana\Foundation\Dashboardv2beta1\VariableOption('.(($input->spec->current->selected !== null) ? 'selected: '.\var_export($input->spec->current->selected, true).', ' : '').'text: '.\var_export($input->spec->current->text, true).',value: '.\var_export($input->spec->current->value, true).',))'.','.(($input->spec->label !== null) ? 'label: '.\var_export($input->spec->label, true).', ' : '').'hide: '.'\Grafana\Foundation\Dashboardv2beta1\VariableHide::fromValue("'.$input->spec->hide.'")'.',refresh: '.'\Grafana\Foundation\Dashboardv2beta1\VariableRefresh::fromValue("'.$input->spec->refresh.'")'.',skipUrlSync: '.\var_export($input->spec->skipUrlSync, true).','.(($input->spec->description !== null) ? 'description: '.\var_export($input->spec->description, true).', ' : '').'query: '.'(new \Grafana\Foundation\Dashboardv2beta1\DataQueryKind(kind: '.\var_export($input->spec->query->kind, true).',group: '.\var_export($input->spec->query->group, true).',version: '.\var_export($input->spec->query->version, true).','.(($input->spec->query->datasource !== null) ? 'datasource: '.'(new \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasource('.(($input->spec->query->datasource->name !== null) ? 'name: '.\var_export($input->spec->query->datasource->name, true).', ' : '').'))'.', ' : '').''.(($input->spec->query->spec !== null) ? 'spec: '.\var_export($input->spec->query->spec, true).', ' : '').'))'.',regex: '.\var_export($input->spec->regex, true).','.(($input->spec->regexApplyTo !== null) ? 'regexApplyTo: '.'\Grafana\Foundation\Dashboardv2beta1\VariableRegexApplyTo::fromValue("'.$input->spec->regexApplyTo.'")'.', ' : '').'sort: '.'\Grafana\Foundation\Dashboardv2beta1\VariableSort::fromValue("'.$input->spec->sort.'")'.','.(($input->spec->definition !== null) ? 'definition: '.\var_export($input->spec->definition, true).', ' : '').'options: '.\var_export($input->spec->options, true).',multi: '.\var_export($input->spec->multi, true).',includeAll: '.\var_export($input->spec->includeAll, true).','.(($input->spec->allValue !== null) ? 'allValue: '.\var_export($input->spec->allValue, true).', ' : '').''.(($input->spec->placeholder !== null) ? 'placeholder: '.\var_export($input->spec->placeholder, true).', ' : '').'allowCustomValue: '.\var_export($input->spec->allowCustomValue, true).','.(($input->spec->staticOptions !== null) ? 'staticOptions: '.\var_export($input->spec->staticOptions, true).', ' : '').''.(($input->spec->staticOptionsOrder !== null) ? 'staticOptionsOrder: '.'\Grafana\Foundation\Dashboardv2beta1\QueryVariableSpecStaticOptionsOrder::fromValue("'.$input->spec->staticOptionsOrder.'")'.', ' : '').'))';
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
    
    
            
    
        {
    $buffer = 'refresh(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\VariableRefresh::fromValue("'.$input->spec->refresh.'")';
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
            
    
        {
    $buffer = 'query(';
        $arg0 = \Grafana\Foundation\Cog\Runtime::get()->convertDataQueryKindToCode($input->spec->query, $input->spec->query->group, );
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
            if ($input->spec->regexApplyTo !== null) {
    
        
    $buffer = 'regexApplyTo(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\VariableRegexApplyTo::fromValue("'.$input->spec->regexApplyTo.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            
    
        {
    $buffer = 'sort(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\VariableSort::fromValue("'.$input->spec->sort.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    }
    
    
            if ($input->spec->definition !== null && $input->spec->definition !== "") {
    
        
    $buffer = 'definition(';
        $arg0 =\var_export($input->spec->definition, true);
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
            if ($input->spec->placeholder !== null && $input->spec->placeholder !== "") {
    
        
    $buffer = 'placeholder(';
        $arg0 =\var_export($input->spec->placeholder, true);
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
            if ($input->spec->staticOptions !== null && count($input->spec->staticOptions) >= 1) {
    
        
    $buffer = 'staticOptions(';
        $tmparg0 = [];
        foreach ($input->spec->staticOptions as $arg1) {
        $tmpstaticOptionsarg1 ='(new \Grafana\Foundation\Dashboardv2beta1\VariableOption('.(($arg1->selected !== null) ? 'selected: '.\var_export($arg1->selected, true).', ' : '').'text: '.\var_export($arg1->text, true).',value: '.\var_export($arg1->value, true).',))';
        $tmparg0[] = $tmpstaticOptionsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec->staticOptionsOrder !== null) {
    
        
    $buffer = 'staticOptionsOrder(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\QueryVariableSpecStaticOptionsOrder::fromValue("'.$input->spec->staticOptionsOrder.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

