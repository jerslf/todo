package view

import (
	"fmt"
	"io"
	"text/tabwriter"

	"github.com/jerslf/todo/internal/task"
	"github.com/mergestat/timediff"
)

func PrintTasks(out io.Writer, ts task.Tasks) {
	if len(ts.Items) == 0 {
		fmt.Fprintln(out, "No tasks found.")
		return
	}

	// Create tabwriter
	w := tabwriter.NewWriter(out, 0, 4, 2, ' ', 0)
	defer w.Flush()

	// Print header
	fmt.Fprintln(w, "ID\tTitle\tStatus\tCreated\tCompleted")

	// Print each task
	for _, t := range ts.Items {
		created := timediff.TimeDiff(t.TimeCreated)

		var doneAt string
		if t.TimeDone != nil {
			doneAt = timediff.TimeDiff(*t.TimeDone)
		} else {
			doneAt = "-"
		}

		status := "❌"
		if t.Done {
			status = "✅"
		}

		fmt.Fprintf(
			w,
			"%d\t%s\t%s\t%s\t%s\n",
			t.ID,
			t.Title,
			status,
			created,
			doneAt,
		)
	}
}
