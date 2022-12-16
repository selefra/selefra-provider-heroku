# Table: heroku_addon

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| addon_service | json | X | √ | Identity of add-on service. | 
| state | string | X | √ | State in the add-on's lifecycle. | 
| plan | json | X | √ | Identity of add-on plan. | 
| actions | json | X | √ | Provider actions for this specific add-on. | 
| created_at | timestamp | X | √ | When add-on was created. | 
| updated_at | timestamp | X | √ | When add-on was updated. | 
| web_url | string | X | √ | URL for logging into web interface of add-on (e.g. a dashboard). | 
| name | string | X | √ | Globally unique name of the add-on. | 
| id | string | X | √ | Unique identifier of add-on. | 
| provider_id | string | X | √ | Id of this add-on with its provider. | 
| app | json | X | √ | Billing application associated with this add-on. | 
| billed_price | json | X | √ | Billed price. | 
| billing_entity | json | X | √ | Billing entity associated with this add-on. | 
| config_vars | json | X | √ | Config vars exposed to the owning app by this add-on. | 


