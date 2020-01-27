package notifiers

import (
	"testing"
	"time"

	"github.com/grafana/grafana/pkg/components/simplejson"
	"github.com/grafana/grafana/pkg/infra/localcache"
	"github.com/grafana/grafana/pkg/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLineNotifier(t *testing.T) {
	cacheService := localcache.New(time.Second, time.Second)

	Convey("Line notifier tests", t, func() {
		Convey("empty settings should return error", func() {
			json := `{ }`

			settingsJSON, _ := simplejson.NewJson([]byte(json))
			model := &models.AlertNotification{
				Name:     "line_testing",
				Type:     "line",
				Settings: settingsJSON,
			}

			_, err := NewLINENotifier(model, cacheService)
			So(err, ShouldNotBeNil)

		})
		Convey("settings should trigger incident", func() {
			json := `
			{
  "token": "abcdefgh0123456789"
			}`
			settingsJSON, _ := simplejson.NewJson([]byte(json))
			model := &models.AlertNotification{
				Name:     "line_testing",
				Type:     "line",
				Settings: settingsJSON,
			}

			not, err := NewLINENotifier(model, cacheService)
			lineNotifier := not.(*LineNotifier)

			So(err, ShouldBeNil)
			So(lineNotifier.Name, ShouldEqual, "line_testing")
			So(lineNotifier.Type, ShouldEqual, "line")
			So(lineNotifier.Token, ShouldEqual, "abcdefgh0123456789")
		})
	})
}
