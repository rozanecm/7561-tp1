package main

import (
	"cloud.google.com/go/cloudtasks/apiv2"
	"context"
	"fmt"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"io/ioutil"
	"log"
	"net/http"
)

func createTask(queueID string) (*tasks.Task, error) {
	/* src: https://github.com/GoogleCloudPlatform/golang-samples/blob/master/appengine/go11x/tasks/create_task */
	// Create a new Cloud Tasks client instance.
	// See https://godoc.org/cloud.google.com/go/cloudtasks/apiv2
	ctx := context.Background()
	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewClient: %v", err)
	}
	defer client.Close()

	// Build the Task queue path.
	queuePath := fmt.Sprintf("projects/%s/locations/%s/queues/%s", "taller3-tp1-rozanecm", "us-east1", queueID)

	// Build the Task payload.
	// https://godoc.org/google.golang.org/genproto/googleapis/cloud/tasks/v2#CreateTaskRequest
	req := &tasks.CreateTaskRequest{
		Parent: queuePath,
		Task: &tasks.Task{
			// https://godoc.org/google.golang.org/genproto/googleapis/cloud/tasks/v2#AppEngineHttpRequest
			MessageType: &tasks.Task_AppEngineHttpRequest{
				AppEngineHttpRequest: &tasks.AppEngineHttpRequest{
					HttpMethod:  tasks.HttpMethod_POST,
					RelativeUri: "/task_handler",
				},
			},
		},
	}

	// Add a payload message if one is present.
	req.Task.GetAppEngineHttpRequest().Body = []byte("")

	createdTask, err := client.CreateTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cloudtasks.CreateTask: %v", err)
	}

	return createdTask, nil
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	taskName := r.Header.Get("X-Appengine-Taskname")
	if taskName == "" {
		// You may use the presence of the X-Appengine-Taskname header to validate
		// the request comes from Cloud Tasks.
		log.Println("Invalid Task: No X-Appengine-Taskname request header found")
		http.Error(w, "Bad Request - Invalid Task", http.StatusBadRequest)
		return
	}

	// Pull useful headers from Task request.
	queueName := r.Header.Get("X-Appengine-Queuename")

	// Extract the request body for further task details.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("ReadAll: %v", err)
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	updateCounter(queueName)

	// Log & output details of the task.
	output := fmt.Sprintf("Completed task: task queue(%s), task name(%s), payload(%s)",
		queueName,
		taskName,
		string(body),
	)
	log.Println(output)

	// Set a non-2xx status code to indicate a failure in task processing that should be retried.
	// For example, http.Error(w, "Internal Server Error: Task Processing", http.StatusInternalServerError)
	fmt.Fprintln(w, output)
}
