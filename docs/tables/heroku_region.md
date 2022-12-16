# Table: heroku_region

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| locale | string | X | √ | Area in the country where the region exists. | 
| updated_at | timestamp | X | √ | When region was updated. | 
| name | string | X | √ | Unique name of region. | 
| id | string | X | √ | Unique identifier of region. | 
| description | string | X | √ | Description of region. | 
| private_capable | bool | X | √ | Whether or not region is available for creating a Private Space. | 
| provider | json | X | √ | Provider of underlying substrate. | 
| country | string | X | √ | Country where the region exists. | 
| created_at | timestamp | X | √ | When region was created. | 


