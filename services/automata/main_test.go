package automata

import (
	"testing"
	"time"

	"github.com/barnybug/gofsm"
	"github.com/barnybug/gohome/config"
	"github.com/barnybug/gohome/pubsub"
	"github.com/barnybug/gohome/services"
	"github.com/stretchr/testify/assert"
)

var (
	evOn      = NewEventContext(pubsub.NewEvent("ack", pubsub.Fields{"device": "light.porch", "command": "on", "timestamp": "2017-09-26 19:24:22.069"}))
	evState   = NewEventContext(pubsub.NewEvent("state", pubsub.Fields{"device": "light.porch", "state": "On", "timestamp": "2017-09-26 19:24:22.069"}))
	evTime    = NewEventContext(pubsub.NewEvent("time", pubsub.Fields{"device": "time", "hhmm": "2230", "timestamp": "2017-09-26 22:30:00.000"}))
	evMissing = NewEventContext(pubsub.NewEvent("ack", pubsub.Fields{"timestamp": "2017-09-26 19:24:22.069"}))
)

func ExampleInterfaces() {
	var _ services.Service = (*Service)(nil)
	var _ services.Queryable = (*Service)(nil)
	// Output:
}

func TestEventSimple(t *testing.T) {
	assert.True(t, evOn.Match("device=='light.porch' && command=='on'"))
	assert.False(t, evOn.Match("device=='light.porch' && command=='off'"))
}

func TestEventType(t *testing.T) {
	assert.True(t, evOn.Match("type=='light' && command=='on'"))
	assert.True(t, evOn.Match("type=='light'"))
}

func TestEventOr(t *testing.T) {
	assert.True(t, evOn.Match("device=='door.front' && command=='off' || device=='light.porch' && command=='on'"))
	assert.True(t, evOn.Match("device=='light.porch' && command=='on' || device=='door.front' && command=='off'"))
}

func TestEventNotABoolean(t *testing.T) {
	assert.False(t, evOn.Match("'abc'"))
}

func TestBadExpression(t *testing.T) {
	assert.False(t, evOn.Match("blah()"))
}

var SimpleAutomata = `
simple:
  start: Start
  states:
    Start: {}
  transitions:
    Start: []
`

func TestStateFunction(t *testing.T) {
	assert.False(t, evOn.Match("State()"))
	automata, _ = gofsm.Load([]byte(SimpleAutomata))
	assert.True(t, evOn.Match("State('simple')=='Start'"))
	assert.False(t, evOn.Match("State('simple')=='Cobblers'"))
	assert.False(t, evOn.Match("State('blah')=='Cobblers'"))
}

func BenchmarkEventTrue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		evOn.Match("device=='door.front' && command=='off' || device=='light.porch' && command=='on'")
	}
}

func BenchmarkEventSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		evOn.Match("device=='light.porch' && command=='on'")
	}
}

func BenchmarkEventFalse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		evMissing.Match("device=='door.front' && command=='off' || device=='light.porch' && command=='on'")
	}
}

func TestEventMissing(t *testing.T) {
	assert.False(t, evMissing.Match("device=='light.porch' && command=='on'"))
}

func TestEventTime(t *testing.T) {
	assert.False(t, evTime.Match("device=='time' && hhmm=='2229'"))
	assert.True(t, evTime.Match("device=='time' && hhmm=='2230'"))
}

func TestEventContextString(t *testing.T) {
	assert.Equal(t, "light.porch command=on", evOn.String())
}

func TestFormat(t *testing.T) {
	services.Config = config.ExampleConfig
	ev := pubsub.NewEvent("state", pubsub.Fields{"device": "light.kitchen", "state": "On", "timestamp": "2017-09-26 19:24:22.069", "number": 2.5})
	now := time.Now()
	change := gofsm.Change{"", "", "", now, time.Minute, nil}
	context := ChangeContext{ev, change}

	assert.Equal(t, "test", context.Format("test"))
	assert.Equal(t, "$missing", context.Format("$missing"))
	assert.Equal(t, "light.kitchen", context.Format("$id"))
	assert.Equal(t, "light", context.Format("$type"))
	assert.Equal(t, "Kitchen", context.Format("$name"))
	assert.Equal(t, "1 minute", context.Format("$duration"))
	assert.Equal(t, "On", context.Format("$state"))
	assert.Equal(t, "2.5", context.Format("$number"))
}
