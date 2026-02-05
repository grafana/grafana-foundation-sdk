<?php

namespace Grafana\Foundation\Bigquery;

final class QueryConverter
{
    public static function convert(\Grafana\Foundation\Dashboardv2beta1\DataQueryKind $input): string
    {
        
        $calls = [
            '(new \Grafana\Foundation\Bigquery\QueryBuilder())',
        ];
            if ($input->version !== "" && $input->version !== "v0") {
    
        
    $buffer = 'version(';
        $arg0 =\var_export($input->version, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->datasource !== null) {
    
        
    $buffer = 'datasource(';
        $arg0 = \Grafana\Foundation\Dashboardv2beta1\Dashboardv2beta1DataQueryKindDatasourceConverter::convert($input->datasource);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->dataset !== null && $input->spec->dataset !== "") {
    
        
    $buffer = 'dataset(';
        $arg0 =\var_export($input->spec->dataset, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->table !== null && $input->spec->table !== "") {
    
        
    $buffer = 'table(';
        $arg0 =\var_export($input->spec->table, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->project !== null && $input->spec->project !== "") {
    
        
    $buffer = 'project(';
        $arg0 =\var_export($input->spec->project, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery) {
    
        
    $buffer = 'format(';
        $arg0 ='\Grafana\Foundation\Bigquery\QueryFormat::fromValue("'.$input->spec->format.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->rawQuery !== null) {
    
        
    $buffer = 'rawQuery(';
        $arg0 =\var_export($input->spec->rawQuery, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->rawSql !== "") {
    
        
    $buffer = 'rawSql(';
        $arg0 =\var_export($input->spec->rawSql, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->location !== null && $input->spec->location !== "") {
    
        
    $buffer = 'location(';
        $arg0 =\var_export($input->spec->location, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->partitioned !== null) {
    
        
    $buffer = 'partitioned(';
        $arg0 =\var_export($input->spec->partitioned, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->partitionedField !== null && $input->spec->partitionedField !== "") {
    
        
    $buffer = 'partitionedField(';
        $arg0 =\var_export($input->spec->partitionedField, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->convertToUTC !== null) {
    
        
    $buffer = 'convertToUTC(';
        $arg0 =\var_export($input->spec->convertToUTC, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->sharded !== null) {
    
        
    $buffer = 'sharded(';
        $arg0 =\var_export($input->spec->sharded, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->queryPriority !== null) {
    
        
    $buffer = 'queryPriority(';
        $arg0 ='\Grafana\Foundation\Bigquery\QueryPriority::fromValue("'.$input->spec->queryPriority.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->timeShift !== null && $input->spec->timeShift !== "") {
    
        
    $buffer = 'timeShift(';
        $arg0 =\var_export($input->spec->timeShift, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->editorMode !== null) {
    
        
    $buffer = 'editorMode(';
        $arg0 ='\Grafana\Foundation\Bigquery\EditorMode::fromValue("'.$input->spec->editorMode.'")';
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->sql !== null) {
    
        
    $buffer = 'sql(';
        $arg0 = \Grafana\Foundation\Bigquery\SQLExpressionConverter::convert($input->spec->sql);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->refId !== null && $input->spec->refId !== "") {
    
        
    $buffer = 'refId(';
        $arg0 =\var_export($input->spec->refId, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->hide !== null) {
    
        
    $buffer = 'hide(';
        $arg0 =\var_export($input->spec->hide, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }
            if ($input->spec !== null && $input->spec instanceof \Grafana\Foundation\Bigquery\Dataquery && $input->spec->queryType !== null && $input->spec->queryType !== "") {
    
        
    $buffer = 'queryType(';
        $arg0 =\var_export($input->spec->queryType, true);
        $buffer .= $arg0;
        
    $buffer .= ')';

    $calls[] = $buffer;
    
    
    }

        return \implode("\n\t->", $calls);
    }
}

