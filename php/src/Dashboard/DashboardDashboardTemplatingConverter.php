<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardDashboardTemplatingConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\DashboardDashboardTemplating $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardDashboardTemplatingBuilder())',
        ];
            if ($input->list !== null && count($input->list) >= 1) {
    
        
    $buffer = 'list(';
        $tmparg0 = [];
        foreach ($input->list as $arg1) {
        $tmplistarg1 = '';
        
        if ($arg1->type === \Grafana\Foundation\Dashboard\VariableType::fromValue("query")) {
            $tmplistarg1 = \Grafana\Foundation\Dashboard\QueryVariableConverter::convert($arg1);
        }
        
        if ($arg1->type === \Grafana\Foundation\Dashboard\VariableType::fromValue("adhoc")) {
            $tmplistarg1 = \Grafana\Foundation\Dashboard\AdHocVariableConverter::convert($arg1);
        }
        
        if ($arg1->type === \Grafana\Foundation\Dashboard\VariableType::fromValue("constant")) {
            $tmplistarg1 = \Grafana\Foundation\Dashboard\ConstantVariableConverter::convert($arg1);
        }
        
        if ($arg1->type === \Grafana\Foundation\Dashboard\VariableType::fromValue("datasource")) {
            $tmplistarg1 = \Grafana\Foundation\Dashboard\DatasourceVariableConverter::convert($arg1);
        }
        
        if ($arg1->type === \Grafana\Foundation\Dashboard\VariableType::fromValue("interval")) {
            $tmplistarg1 = \Grafana\Foundation\Dashboard\IntervalVariableConverter::convert($arg1);
        }
        
        if ($arg1->type === \Grafana\Foundation\Dashboard\VariableType::fromValue("textbox")) {
            $tmplistarg1 = \Grafana\Foundation\Dashboard\TextBoxVariableConverter::convert($arg1);
        }
        
        if ($arg1->type === \Grafana\Foundation\Dashboard\VariableType::fromValue("custom")) {
            $tmplistarg1 = \Grafana\Foundation\Dashboard\CustomVariableConverter::convert($arg1);
        }
        
        $tmparg0[] = $tmplistarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

