# kundankumarjha-add
This activity provides your flogo application with rudementary adding facility.


## Installation

```bash
flogo add activity github.com/kundankumarjha/flogo/activity/add
```

## Schema
Inputs and Outputs:

```json
{
  "inputs":[
      {
        "name": "number1",
        "type": "integer"
      },
      {
        "name": "number2",
        "type": "integer"
      }
  ],
  "outputs": [
      {
        "name": "sum",
        "type": "integer"
      }
  ]
}
```
## Settings
| Setting   | Description    |
|:----------|:---------------|
| number1   | First input number |
| number2   | Second input number |
| sum       | Add inputs to the 'sum' output of the activity |


## Configuration Examples
### Simple
Configure a task to add 2 numbers:

```json
{
  "id": 3,
  "type": 1,
  "activityType": "kundankumarjha-add",
  "name": "Add Numbers",
  "attributes": [
    { "name": "number1", "value": 10 },
    { "name": "number2", "value": 20 }
  ]
}
```
