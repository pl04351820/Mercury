# API

This documents describes the back-end API provided by Mercury.

### AddJob
This API adds new job to mercury state machine.

ParameterName | Description
:---: | :----:
TaskName | The name of task 
TaskInfo | The json file to describe the state of Task.

### DescribeJobState

This API checks the state of Job.

ParameterName | Description
:---: | :---:
TaskName | The name of task