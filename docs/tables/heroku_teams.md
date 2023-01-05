# Table: heroku_teams

## Primary Keys 

```
id
```


## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| created_at | timestamp | X | √ |  | 
| credit_card_collections | bool | X | √ |  | 
| id | string | √ | √ |  | 
| identity_provider | json | X | √ |  | 
| name | string | X | √ |  | 
| updated_at | timestamp | X | √ |  | 
| selefra_id | string | √ | √ | primary keys value md5 | 
| default | bool | X | √ |  | 
| enterprise_account | json | X | √ |  | 
| membership_limit | float | X | √ |  | 
| provisioned_licenses | bool | X | √ |  | 
| role | string | X | √ |  | 
| type | string | X | √ |  | 


