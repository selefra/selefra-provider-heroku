# Table: heroku_review_apps

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| status | string | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| created_at | timestamp | X | √ |  | 
| error_status | string | X | √ |  | 
| pr_number | int | X | √ |  | 
| id | string | √ | √ |  | 
| pipeline | json | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| wait_for_ci | bool | X | √ |  | 
| app | json | X | √ |  | 
| branch | string | X | √ |  | 
| fork_repo | json | X | √ |  | 
| creator | json | X | √ |  | 
| app_setup | json | X | √ |  | 
| message | string | X | √ |  | 


