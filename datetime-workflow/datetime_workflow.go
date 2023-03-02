package datetime_workflow

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	aw "github.com/deanishe/awgo"
)

const (
	DatetimeFormat1 = "2006-01-02 15:04:05"
	DatetimeFormat2 = "20060102150405"
)

var (
	datetimeLayouts = []string{
		"2006-01-02",
		"2006-01-02 15:04",
		DatetimeFormat1,
		DatetimeFormat2,
	}
)

type DatetimeWorkflow struct {
	Workflow *aw.Workflow
	items    []WorkflowItem
}

type WorkflowItem struct {
	Title    string
	Arg      string
	Subtitle string
}

func (w *DatetimeWorkflow) Run() {
	args := w.Workflow.Args()

	if len(args) == 0 {
		return
	}

	arg := strings.ToLower(strings.Join(args, ""))
	if arg == "n" || arg == "no" || arg == "now" {
		w.handleNow()
	} else {
		err := w.handleDatetime(arg)
		if err != nil {
			stamp, err := strconv.ParseUint(arg, 10, 64)
			if err != nil {
				return
			}
			w.handleTimestamp(int64(stamp))
		}
	}

	w.sendItems()
}

func (w *DatetimeWorkflow) handleTimestamp(stamp int64) {
	t := time.Unix(stamp, 0)
	w.items = append(w.items, WorkflowItem{
		Title: "目标时间：" + t.Format(DatetimeFormat1),
		Arg:   t.Format(DatetimeFormat1),
	}, WorkflowItem{
		Title: "目标时间：" + t.Format(DatetimeFormat2),
		Arg:   t.Format(DatetimeFormat2),
	})
}

func (w *DatetimeWorkflow) handleDatetime(arg string) error {
	var t time.Time
	var err error
	location, _ := time.LoadLocation("Asia/Shanghai")
	for _, layout := range datetimeLayouts {
		t, err = time.ParseInLocation(layout, arg, location)
		if err != nil {
			continue
		} else {
			break
		}
	}
	if err != nil {
		return err
	}

	w.items = append(w.items, WorkflowItem{
		Title: "目标时间戳：" + fmt.Sprint(t.Unix()),
		Arg:   fmt.Sprint(t.Unix()),
	}, WorkflowItem{
		Title: "目标毫秒时间戳：" + fmt.Sprint(t.UnixMilli()),
		Arg:   fmt.Sprint(t.UnixMilli()),
	})

	return nil
}

func (w *DatetimeWorkflow) handleNow() {
	now := time.Now()
	w.items = append(w.items, WorkflowItem{
		Title: "当前时间戳：" + fmt.Sprint(now.Unix()),
		Arg:   fmt.Sprint(now.Unix()),
	}, WorkflowItem{
		Title: "当前毫秒时间戳：" + fmt.Sprint(now.UnixMilli()),
		Arg:   fmt.Sprint(now.UnixMilli()),
	}, WorkflowItem{
		Title: "当前时间：" + now.Format(DatetimeFormat1),
		Arg:   now.Format(DatetimeFormat1),
	}, WorkflowItem{
		Title: "当前时间：" + now.Format(DatetimeFormat2),
		Arg:   now.Format(DatetimeFormat2),
	})
}

func (w *DatetimeWorkflow) sendItems() {
	for _, item := range w.items {
		w.Workflow.NewItem(item.Title).Arg(item.Arg).Valid(true).Subtitle(item.Subtitle)
	}
	w.Workflow.SendFeedback()
}
