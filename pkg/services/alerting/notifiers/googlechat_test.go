package notifiers

import (
	"testing"
	"time"

	"github.com/grafana/grafana/pkg/components/simplejson"
	"github.com/grafana/grafana/pkg/infra/localcache"
	"github.com/grafana/grafana/pkg/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGoogleChatNotifier(t *testing.T) {
	cacheService := localcache.New(time.Second, time.Second)

	Convey("Google Hangouts Chat notifier tests", t, func() {

		Convey("Parsing alert notification from settings", func() {
			Convey("empty settings should return error", func() {
				json := `{ }`

				settingsJSON, _ := simplejson.NewJson([]byte(json))
				model := &models.AlertNotification{
					Name:     "ops",
					Type:     "googlechat",
					Settings: settingsJSON,
				}

				_, err := newGoogleChatNotifier(model, cacheService)
				So(err, ShouldNotBeNil)
			})

			Convey("from settings", func() {
				json := `
				{
          			"url": "http://google.com"
				}`

				settingsJSON, _ := simplejson.NewJson([]byte(json))
				model := &models.AlertNotification{
					Name:     "ops",
					Type:     "googlechat",
					Settings: settingsJSON,
				}

				not, err := newGoogleChatNotifier(model, cacheService)
				webhookNotifier := not.(*GoogleChatNotifier)

				So(err, ShouldBeNil)
				So(webhookNotifier.Name, ShouldEqual, "ops")
				So(webhookNotifier.Type, ShouldEqual, "googlechat")
				So(webhookNotifier.URL, ShouldEqual, "http://google.com")
			})
		})
	})
}
