# Table: heroku_domains

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| acm_status | string | X | √ |  | 
| acm_status_reason | string | X | √ |  | 
| sni_endpoint | json | X | √ |  | 
| status | string | X | √ |  | 
| id | string | √ | √ |  | 
| kind | string | X | √ |  | 
| app | json | X | √ |  | 
| cname | string | X | √ |  | 
| created_at | timestamp | X | √ |  | 
| hostname | string | X | √ |  | 


