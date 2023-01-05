# Table: heroku_add_ons

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| actions | json | X | √ |  | 
| id | string | √ | √ |  | 
| name | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| web_url | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| addon_service | json | X | √ |  | 
| app | json | X | √ |  | 
| plan | json | X | √ |  | 
| provider_id | string | X | √ |  | 
| state | string | X | √ |  | 
| billed_price | json | X | √ |  | 
| config_vars | string_array | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| billing_entity | json | X | √ |  | 


