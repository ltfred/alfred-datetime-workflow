package datetime_workflow

import "testing"

func TestDatetimeWorkflow(t *testing.T) {
	w := DatetimeWorkflow{}
	w.handleDatetime(DatetimeFormat1)

	t.Logf("items: %v", w.items)
}
