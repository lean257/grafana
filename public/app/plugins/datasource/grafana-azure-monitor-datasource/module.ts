import { DataSourcePlugin } from '@grafana/data';
import Datasource from './datasource';
import { ConfigEditor } from './components/ConfigEditor';
import { AzureMonitorAnnotationsQueryCtrl } from './annotations_query_ctrl';
import { AzureMonitorQuery, AzureDataSourceJsonData } from './types';
import AzureMonitorQueryEditor from './components/QueryEditor';

export const plugin = new DataSourcePlugin<Datasource, AzureMonitorQuery, AzureDataSourceJsonData>(Datasource)
  .setConfigEditor(ConfigEditor)
  .setQueryEditor(AzureMonitorQueryEditor)
  .setAnnotationQueryCtrl(AzureMonitorAnnotationsQueryCtrl);
