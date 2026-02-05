<?php

namespace Grafana\Foundation\Dashboardv2beta1;

final class TimeSettingsConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboardv2beta1\TimeSettingsBuilder())',
        ];
            if ($input->timezone !== null && $input->timezone !== "" && $input->timezone !== "browser") {
    
        
    $buffer = 'timezone(';
        $arg0 =\var_export($input->timezone, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->from !== "" && $input->from !== "now-6h") {
    
        
    $buffer = 'from(';
        $arg0 =\var_export($input->from, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->to !== "" && $input->to !== "now") {
    
        
    $buffer = 'to(';
        $arg0 =\var_export($input->to, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->autoRefresh !== "") {
    
        
    $buffer = 'autoRefresh(';
        $arg0 =\var_export($input->autoRefresh, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if (count($input->autoRefreshIntervals) >= 1) {
    
        
    $buffer = 'autoRefreshIntervals(';
        $tmparg0 = [];
        foreach ($input->autoRefreshIntervals as $arg1) {
        $tmpautoRefreshIntervalsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmpautoRefreshIntervalsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->quickRanges !== null && count($input->quickRanges) >= 1) {
    
        
    $buffer = 'quickRanges(';
        $tmparg0 = [];
        foreach ($input->quickRanges as $arg1) {
        $tmpquickRangesarg1 = \Grafana\Foundation\Dashboardv2beta1\TimeRangeOptionConverter::convert($arg1);
        $tmparg0[] = $tmpquickRangesarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->hideTimepicker !== false) {
    
        
    $buffer = 'hideTimepicker(';
        $arg0 =\var_export($input->hideTimepicker, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->weekStart !== null) {
    
        
    $buffer = 'weekStart(';
        $arg0 ='\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpecWeekStart::fromValue("'.$input->weekStart.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fiscalYearStartMonth !== 0) {
    
        
    $buffer = 'fiscalYearStartMonth(';
        $arg0 =\var_export($input->fiscalYearStartMonth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->nowDelay !== null && $input->nowDelay !== "") {
    
        
    $buffer = 'nowDelay(';
        $arg0 =\var_export($input->nowDelay, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

