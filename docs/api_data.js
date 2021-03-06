define({ "api": [
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./docs/main.js",
    "group": "_home_angad_go_src_github_com_angadsharma1016_socialcops_docs_main_js",
    "groupTitle": "_home_angad_go_src_github_com_angadsharma1016_socialcops_docs_main_js",
    "name": ""
  },
  {
    "type": "get",
    "url": "/api/v1/process/kill",
    "title": "kill the long running task",
    "name": "kill_the_long_running_task",
    "group": "all",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "uint",
            "optional": false,
            "field": "id",
            "description": "<p>id to kill the task by</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "request-example",
          "content": "\ncurl http://<domain:port>/api/v1/process/kill?id=1",
          "type": "string"
        },
        {
          "title": "response-example",
          "content": "\nSent kill signal to task 1",
          "type": "string"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controller/kill.go",
    "groupTitle": "all"
  },
  {
    "type": "get",
    "url": "/api/v1/process/start",
    "title": "start the long running task",
    "name": "start_the_long_running_task",
    "group": "all",
    "parameter": {
      "examples": [
        {
          "title": "request-example",
          "content": "\ncurl http://<domain:port>/api/v1/process/start",
          "type": "string"
        },
        {
          "title": "response-example",
          "content": "\n{\n\t\"taskID\": 1,\n\t\"timestamp\": \"string Unix timestamp\",\n\t\"taskName\": \"bigproc\",\n\t\"isCompleted\":false,\n\t\"wasInterrupted\": false\n}",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controller/start.go",
    "groupTitle": "all"
  },
  {
    "type": "get",
    "url": "/api/v1/process/view",
    "title": "view all the long running tasks",
    "name": "view_the_long_running_task",
    "group": "all",
    "parameter": {
      "examples": [
        {
          "title": "request-example",
          "content": "\ncurl http://<domain:port>/api/v1/process/view",
          "type": "string"
        },
        {
          "title": "response-example",
          "content": "\n[\n\t{\n\t\t\"taskID\": 1,\n\t\t\"timestamp\": \"string Unix timestamp\",\n\t\t\"taskName\": \"bigproc\",\n\t\t\"isCompleted\":false,\n\t\t\"wasInterrupted\": false\n\t},\n\t{\n\t\t\"taskID\": 2,\n\t\t\"timestamp\": \"string Unix timestamp\",\n\t\t\"taskName\": \"bigproc\",\n\t\t\"isCompleted\":false,\n\t\t\"wasInterrupted\": false\n\t},\n\t{\t\"taskID\": 3,\n\t\t\"timestamp\": \"string Unix timestamp\",\n\t\t\"taskName\": \"bigproc\",\n\t\t\"isCompleted\":false,\n\t\t\"wasInterrupted\": false\n\t}\n]",
          "type": "json"
        }
      ]
    },
    "version": "0.0.0",
    "filename": "./controller/view.go",
    "groupTitle": "all"
  }
] });
