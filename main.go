package GoQ

import "net/http"

func AddTaskHandler(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := PushTask("task_queue", task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task added successfully"})
}
