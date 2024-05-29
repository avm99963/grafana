package core

timeInterval: {
	kind:  "TimeInterval"
	group: "notifications"
	apiResource: {
		groupOverride: "notifications.alerting.grafana.app"
	}
	codegen: {
		frontend: true
		backend:  true
	}
	pluralName: "TimeIntervals"
	current:    "v0alpha1"
	versions: {
		"v0alpha1": {
			schema: {
				#TimeRange: {
					start_time: string
					end_time: string
				}
				#Interval: {
					times?: [...#TimeRange]
					weekdays?: [...string]
					days_of_month?: [...string]
					months?: [...string]
					years?: [...string]
					location?: string
				}
				spec: {
					time_intervals: [...#Interval]
				}
			}
		}
	}
}