### The Challenge

Collect has a variety of long-running tasks which require time and resources on the servers. As it stands now, once we have triggered off a long-running task, there is no way to tap into it and pause/stop/terminate the task, upon realizing that an erroneous request went through from one of the clients (mostly web or pipeline).

### Examples:

* Example 1: A support user starts a baseline upload for an organization. The baseline upload had 100000 rows in the CSV file uploaded, and the support user - right after clicking the “Start” button - realizes that he seemed to have wrongly exchanged data between two adjacent columns. Subsequently, the support user would need to wait until the current request completes, which can be in the magnitude of hours, and since we don't allow two concurrent baseline uploads in Labs right now, this implies he would not be able to upload the correct file till this task completes. Once done, the user would need to remove those 1 lac entries, and finally upload the correct file.

* Example 2: The pipeline sends an export request to Collect to update the data on dashboards. In one of the incidents triggered manually, the “from date” in the export request was (by mistake) set as “2015-08-01”, while it was meant to be set as “2017-08-01”. As a result, the export started unnecessarily for 20 million rows, whereas setting the right dates would have resulted in only 2 million rows. There is no way for me to stop the previous export, so that time and resources are concentrated upon this request to be completed.

* Example 3: A user on the web dashboard wanted to bulk create teams. There were 10000 teams to be created which would take an hour or more. They observed that phone number mappings to teams had been mistakenly shifted down ten rows, due to which inappropriate managers would be added to teams. Now the user wants to stop the request, correct the mapping, and restart the upload.

You can try out the product at https://collect.socialcops.com to get understanding of above problems.

Create a Rest API endpoints where above three examples can be served and package the solution in a docker image which can be deployed on Kubernetes directly.

**Please use any file with dummy data.




Design the architecture using:
Task Management(Worker queue)
Monitoring and Reporting
Resource utilization
Fail-safe Mechanism

You can use following technology stack for the implementation:
Python/Javascript
RabbitMQ(Celery)
Flask/Pyramid/Node
Redis
MongoDB/SQL/Elasticsearch

P.S. Share your ideas around how we can scale Long-running jobs and what are the potential issues which can come up.
