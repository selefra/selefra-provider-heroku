# Table: heroku_app_webhook_deliveries

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| next_attempt_at | timestamp | X | √ |  | 
| status | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| webhook | json | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| event | json | X | √ |  | 
| id | string | √ | √ |  | 
| last_attempt | json | X | √ |  | 
| num_attempts | int | X | √ |  | 
| created_at | timestamp | X | √ |  | 


