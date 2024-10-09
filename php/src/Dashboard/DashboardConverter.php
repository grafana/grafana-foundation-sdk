<?php

namespace Grafana\Foundation\Dashboard;

final class DashboardConverter
{
    public static function convert(\Grafana\Foundation\Dashboard\Dashboard $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Dashboard\DashboardBuilder('.\var_export($input->title, true).'))',
        ];
            if ($input->id !== null) {
    
        
    $buffer = 'id(';
        $arg0 =\var_export($input->id, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->uid !== null && $input->uid !== "") {
    
        
    $buffer = 'uid(';
        $arg0 =\var_export($input->uid, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->title !== null && $input->title !== "") {
    
        
    $buffer = 'title(';
        $arg0 =\var_export($input->title, true);
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
            if ($input->revision !== null) {
    
        
    $buffer = 'revision(';
        $arg0 =\var_export($input->revision, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->gnetId !== null && $input->gnetId !== "") {
    
        
    $buffer = 'gnetId(';
        $arg0 =\var_export($input->gnetId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->tags !== null && count($input->tags) >= 1) {
    
        
    $buffer = 'tags(';
        $tmparg0 = [];
        foreach ($input->tags as $arg1) {
        $tmptagsarg1 =\var_export($arg1, true);
        $tmparg0[] = $tmptagsarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timezone !== null && $input->timezone !== "" && $input->timezone !== "browser") {
    
        
    $buffer = 'timezone(';
        $arg0 =\var_export($input->timezone, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->editable !== null && $input->editable === true) {
    
        
    $buffer = 'editable(';
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->editable !== null && $input->editable === false) {
    
        
    $buffer = 'readonly(';
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->graphTooltip !== null) {
    
        
    $buffer = 'tooltip(';
        $arg0 ='\Grafana\Foundation\Dashboard\DashboardCursorSync::fromValue("'.$input->graphTooltip.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->time !== null && $input->time->from !== "" && $input->time->from !== "now-6h" && $input->time->to !== "" && $input->time->to !== "now") {
    
        
    $buffer = 'time(';
        $arg0 =\var_export($input->time->from, true);
        $buffer .= $arg0;
        $buffer .= ', ';
        $arg1 =\var_export($input->time->to, true);
        $buffer .= $arg1;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->timepicker !== null) {
    
        
    $buffer = 'timepicker(';
        $arg0 = \Grafana\Foundation\Dashboard\TimePickerConverter::convert($input->timepicker);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->fiscalYearStartMonth !== null && $input->fiscalYearStartMonth !== 0) {
    
        
    $buffer = 'fiscalYearStartMonth(';
        $arg0 =\var_export($input->fiscalYearStartMonth, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->liveNow !== null) {
    
        
    $buffer = 'liveNow(';
        $arg0 =\var_export($input->liveNow, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->weekStart !== null && $input->weekStart !== "") {
    
        
    $buffer = 'weekStart(';
        $arg0 =\var_export($input->weekStart, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->refresh !== null) {
    
        
    $buffer = 'refresh(';
        switch (true) {
            case is_string($input->refresh):
                $disjunctionrefresh =\var_export($input->refresh, true);
                $arg0 = $disjunctionrefresh;
                break;
            case is_bool($input->refresh):
                $disjunctionrefresh =\var_export($input->refresh, true);
                $arg0 = $disjunctionrefresh;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->version !== null) {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->panels !== null && count($input->panels) >= 1) {
    foreach ($input->panels as $item) {
        
    $buffer = 'withPanel(';
        switch (true) {
            case $item instanceof \Grafana\Foundation\Dashboard\Panel:
                $disjunctionitem = \Grafana\Foundation\Cog\Runtime::get()->convertPanelToCode($item, $item->type, );
                $arg0 = $disjunctionitem;
                break;
            case $item instanceof \Grafana\Foundation\Dashboard\RowPanel:
                $disjunctionitem = \Grafana\Foundation\Dashboard\RowConverter::convert($item);
                $arg0 = $disjunctionitem;
                break;
            default:
                throw new \ValueError('disjunction branch not handled');
        }
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    }
    }
            if ($input->templating->list !== null && count($input->templating->list) >= 1) {
    
        
    $buffer = 'variables(';
        $tmparg0 = [];
        foreach ($input->templating->list as $arg1) {
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
            if ($input->annotations->list !== null && count($input->annotations->list) >= 1) {
    
        
    $buffer = 'annotations(';
        $tmparg0 = [];
        foreach ($input->annotations->list as $arg1) {
        $tmplistarg1 = \Grafana\Foundation\Dashboard\AnnotationQueryConverter::convert($arg1);
        $tmparg0[] = $tmplistarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->links !== null && count($input->links) >= 1) {
    
        
    $buffer = 'links(';
        $tmparg0 = [];
        foreach ($input->links as $arg1) {
        $tmplinksarg1 = \Grafana\Foundation\Dashboard\DashboardLinkConverter::convert($arg1);
        $tmparg0[] = $tmplinksarg1;
        }
        $arg0 = "[" . implode(", \n", $tmparg0) . "]";
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->snapshot !== null) {
    
        
    $buffer = 'snapshot(';
        $arg0 = \Grafana\Foundation\Dashboard\SnapshotConverter::convert($input->snapshot);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

