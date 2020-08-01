package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/petoc/gbfs"
)

func getFeedHandlers(db *sql.DB) []*gbfs.FeedHandler {
	return []*gbfs.FeedHandler{
		&gbfs.FeedHandler{
			// TTL: 60,
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feedSK := &gbfs.FeedSystemInformation{
					Data: &gbfs.FeedSystemInformationData{
						SystemID:    s.Options.SystemID,
						Language:    "en",
						Name:        "Bike Sharing",
						Operator:    "Bike Sharing, Street 123, 12345 City",
						URL:         "http://localhost/bikesharing/sk",
						PhoneNumber: "00421987654321",
						Email:       "bikesharing@localhost",
						Timezone:    "Europe/Bratislava",
					},
				}
				feedSK.Language = "sk"
				// feedSK.TTL = 60
				feedEN := &gbfs.FeedSystemInformation{
					Data: &gbfs.FeedSystemInformationData{
						SystemID:    s.Options.SystemID,
						Language:    "en",
						Name:        "Bike Sharing",
						Operator:    "Bike Sharing, Street 123, 12345 City",
						URL:         "http://localhost/bikesharing/en",
						PhoneNumber: "00421987654321",
						Email:       "bikesharing@localhost",
						Timezone:    "Europe/Bratislava",
					},
				}
				feedEN.Language = "en"
				// feedEN.TTL = 60
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feedSK := &gbfs.FeedVehicleTypes{
					Data: &gbfs.FeedVehicleTypesData{
						VehicleTypes: []*gbfs.FeedVehicleTypesType{
							&gbfs.FeedVehicleTypesType{
								VehicleTypeID:  "vehicleType1",
								FormFactor:     gbfs.FormFactorBicycle,
								PropulsionType: gbfs.PropulsionTypeHuman,
								Name:           "Bicykel",
							},
							&gbfs.FeedVehicleTypesType{
								VehicleTypeID:  "vehicleType2",
								FormFactor:     gbfs.FormFactorMoped,
								PropulsionType: gbfs.PropulsionTypeElectric,
								Name:           "Skúter",
							},
						},
					},
				}
				feedSK.Language = "sk"
				feedEN := &gbfs.FeedVehicleTypes{
					Data: &gbfs.FeedVehicleTypesData{
						VehicleTypes: []*gbfs.FeedVehicleTypesType{
							&gbfs.FeedVehicleTypesType{
								VehicleTypeID:  "vehicleType1",
								FormFactor:     gbfs.FormFactorBicycle,
								PropulsionType: gbfs.PropulsionTypeHuman,
								Name:           "Bicycle",
							},
							&gbfs.FeedVehicleTypesType{
								VehicleTypeID:  "vehicleType2",
								FormFactor:     gbfs.FormFactorMoped,
								PropulsionType: gbfs.PropulsionTypeElectric,
								Name:           "Moped",
							},
						},
					},
				}
				feedEN.Language = "en"
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feedSK := &gbfs.FeedStationInformation{
					Data: &gbfs.FeedStationInformationData{
						Stations: []*gbfs.FeedStationInformationStation{
							&gbfs.FeedStationInformationStation{
								StationID: "station1",
								Name:      "Stanica",
								Lat:       48.1234,
								Lon:       21.1234,
								Address:   "Ulica 123",
								StationArea: gbfs.NewGeoJSONGeometryMultiPolygon(
									[][][][]float64{
										[][][]float64{
											[][]float64{
												[]float64{16.8331891, 47.7314286},
												[]float64{22.56571, 47.7314286},
												[]float64{22.56571, 49.6138162},
												[]float64{16.8331891, 49.6138162},
												[]float64{16.8331891, 47.7314286},
											},
										},
									},
									nil,
								),
							},
						},
					},
				}
				feedSK.Language = "sk"
				feedEN := &gbfs.FeedStationInformation{
					Data: &gbfs.FeedStationInformationData{
						Stations: []*gbfs.FeedStationInformationStation{
							&gbfs.FeedStationInformationStation{
								StationID: "station1",
								Name:      "Station",
								Lat:       48.1234,
								Lon:       21.1234,
								Address:   "Street 123",
								StationArea: gbfs.NewGeoJSONGeometryMultiPolygon(
									[][][][]float64{
										[][][]float64{
											[][]float64{
												[]float64{16.8331891, 47.7314286},
												[]float64{22.56571, 47.7314286},
												[]float64{22.56571, 49.6138162},
												[]float64{16.8331891, 49.6138162},
												[]float64{16.8331891, 47.7314286},
											},
										},
									},
									nil,
								),
							},
						},
					},
				}
				feedEN.Language = "en"
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				stations := []*gbfs.FeedStationStatusStation{
					&gbfs.FeedStationStatusStation{
						StationID:             "station1",
						NumBikesAvailable:     2,
						NumBikesDisabled:      0,
						NumDocksAvailable:     0,
						NumDocksDisabled:      0,
						IsInstalled:           gbfs.Boolean(true),
						IsRenting:             gbfs.Boolean(true),
						IsReturning:           gbfs.Boolean(true),
						LastReported:          gbfs.Timestamp(1577836800),
						VehicleDocksAvailable: []*gbfs.FeedStationStatusVehicleDock{},
						Vehicles:              []*gbfs.FeedStationStatusVehicle{},
					},
				}
				feedSK := &gbfs.FeedStationStatus{
					Data: &gbfs.FeedStationStatusData{
						Stations: stations,
					},
				}
				feedSK.Language = "sk"
				feedEN := &gbfs.FeedStationStatus{
					Data: &gbfs.FeedStationStatusData{
						Stations: stations,
					},
				}
				feedEN.Language = "en"
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				bicycle1 := &gbfs.FeedFreeBikeStatusBike{
					BikeID:        "bicyle1",
					SystemID:      s.Options.SystemID,
					Lat:           48.7162,
					Lon:           21.2613,
					VehicleTypeID: "vehicleType1",
					IsReserved:    gbfs.Boolean(false),
					IsDisabled:    gbfs.Boolean(false),
					LastReported:  gbfs.Timestamp(1577836800),
				}
				moped1 := &gbfs.FeedFreeBikeStatusBike{
					BikeID:             "moped1",
					SystemID:           s.Options.SystemID,
					Lat:                48.7162,
					Lon:                21.2613,
					VehicleTypeID:      "vehicleType2",
					IsReserved:         gbfs.Boolean(false),
					IsDisabled:         gbfs.Boolean(false),
					LastReported:       gbfs.Timestamp(1577836800),
					CurrentRangeMeters: 12345.67,
				}
				feedSK := &gbfs.FeedFreeBikeStatus{
					Data: &gbfs.FeedFreeBikeStatusData{
						Bikes: []*gbfs.FeedFreeBikeStatusBike{
							bicycle1,
							moped1,
						},
					},
				}
				feedSK.Language = "sk"
				feedEN := &gbfs.FeedFreeBikeStatus{
					Data: &gbfs.FeedFreeBikeStatusData{
						Bikes: []*gbfs.FeedFreeBikeStatusBike{
							bicycle1,
							moped1,
						},
					},
				}
				feedEN.Language = "en"
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				data := &gbfs.FeedSystemHoursData{
					RentalHours: []*gbfs.FeedSystemHoursRentalHour{
						&gbfs.FeedSystemHoursRentalHour{
							UserTypes: gbfs.UserTypeAll(),
							Days:      gbfs.DayAll(),
							StartTime: gbfs.Time("00:00:00"),
							EndTime:   gbfs.Time("23:59:59"),
						},
					},
				}
				feedSK := &gbfs.FeedSystemHours{
					Data: data,
				}
				feedSK.Language = "sk"
				feedEN := &gbfs.FeedSystemHours{
					Data: data,
				}
				feedEN.Language = "en"
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				data := &gbfs.FeedSystemCalendarData{
					Calendars: []*gbfs.FeedSystemCalendarCalendar{
						&gbfs.FeedSystemCalendarCalendar{
							StartMonth: 1,
							StartDay:   1,
							StartYear:  2020,
							EndMonth:   12,
							EndDay:     31,
							EndYear:    2020,
						},
					},
				}
				feedSK := &gbfs.FeedSystemCalendar{
					Data: data,
				}
				feedSK.Language = "sk"
				feedEN := &gbfs.FeedSystemCalendar{
					Data: data,
				}
				feedEN.Language = "en"
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				regions := []*gbfs.FeedSystemRegionsRegion{
					&gbfs.FeedSystemRegionsRegion{
						RegionID: "region1",
						Name:     "Region Name",
					},
				}
				feedSK := &gbfs.FeedSystemRegions{
					Data: &gbfs.FeedSystemRegionsData{
						Regions: regions,
					},
				}
				feedSK.Language = "sk"
				feedEN := &gbfs.FeedSystemRegions{
					Data: &gbfs.FeedSystemRegionsData{
						Regions: regions,
					},
				}
				feedEN.Language = "en"
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feedSK := &gbfs.FeedSystemPricingPlans{
					Data: &gbfs.FeedSystemPricingPlansData{
						Plans: []*gbfs.FeedSystemPricingPlansPricingPlan{
							&gbfs.FeedSystemPricingPlansPricingPlan{
								PlanID:      "plan1",
								Name:        "Cenový plán",
								Currency:    "EUR",
								Price:       12.34,
								IsTaxable:   gbfs.Boolean(false),
								Description: "Popis cenového plánu",
							},
						},
					},
				}
				feedSK.Language = "sk"
				feedEN := &gbfs.FeedSystemPricingPlans{
					Data: &gbfs.FeedSystemPricingPlansData{
						Plans: []*gbfs.FeedSystemPricingPlansPricingPlan{
							&gbfs.FeedSystemPricingPlansPricingPlan{
								PlanID:      "plan1",
								Name:        "Pricing Plan",
								Currency:    "EUR",
								Price:       12.34,
								IsTaxable:   gbfs.Boolean(false),
								Description: "Pricing plan description",
							},
						},
					},
				}
				feedEN.Language = "en"
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feedSK := &gbfs.FeedSystemAlerts{
					Data: &gbfs.FeedSystemAlertsData{
						Alerts: []*gbfs.FeedSystemAlertsAlert{
							&gbfs.FeedSystemAlertsAlert{
								AlertID: "alert1",
								Type:    gbfs.AlertTypeSystemClosure,
								Times: []*gbfs.FeedSystemAlertsAlertTime{
									&gbfs.FeedSystemAlertsAlertTime{
										Start: gbfs.Timestamp(1577865600),
										End:   gbfs.Timestamp(1577908800),
									},
								},
								StationIDs:  []string{"station1"},
								RegionIDs:   []string{"region1"},
								URL:         "http://localhost/sk/alerts/alert1",
								Summary:     "Zhrnutie upozornenia",
								Description: "Popis upozornenia",
								LastUpdated: gbfs.Timestamp(1577836800),
							},
						},
					},
				}
				feedSK.Language = "sk"
				feedEN := &gbfs.FeedSystemAlerts{
					Data: &gbfs.FeedSystemAlertsData{
						Alerts: []*gbfs.FeedSystemAlertsAlert{
							&gbfs.FeedSystemAlertsAlert{
								AlertID: "alert1",
								Type:    gbfs.AlertTypeSystemClosure,
								Times: []*gbfs.FeedSystemAlertsAlertTime{
									&gbfs.FeedSystemAlertsAlertTime{
										Start: gbfs.Timestamp(1577865600),
										End:   gbfs.Timestamp(1577908800),
									},
								},
								StationIDs:  []string{"station1"},
								RegionIDs:   []string{"region1"},
								URL:         "http://localhost/en/alerts/alert1",
								Summary:     "Alert summary",
								Description: "Alert description",
								LastUpdated: gbfs.Timestamp(1577836800),
							},
						},
					},
				}
				feedEN.Language = "en"
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				geometry := gbfs.NewGeoJSONGeometryMultiPolygon(
					[][][][]float64{
						[][][]float64{
							[][]float64{
								[]float64{16.8331891, 47.7314286},
								[]float64{22.56571, 47.7314286},
								[]float64{22.56571, 49.6138162},
								[]float64{16.8331891, 49.6138162},
								[]float64{16.8331891, 47.7314286},
							},
						},
					},
					nil,
				)
				rules := []*gbfs.FeedGeofencingZonesGeoJSONFeaturePropertiesRule{
					&gbfs.FeedGeofencingZonesGeoJSONFeaturePropertiesRule{
						VehicleTypeIDs:     []string{"vehicleType1", "vehicleType2"},
						RideAllowed:        gbfs.Boolean(true),
						RideThroughAllowed: gbfs.Boolean(true),
						MaximumSpeedKph:    15,
					},
				}
				feedSK := &gbfs.FeedGeofencingZones{
					Data: &gbfs.FeedGeofencingZonesData{
						GeofencingZones: gbfs.NewFeedGeofencingZonesGeoJSONFeatureCollection(
							[]*gbfs.FeedGeofencingZonesGeoJSONFeature{
								gbfs.NewFeedGeofencingZonesGeoJSONFeature(
									geometry,
									&gbfs.FeedGeofencingZonesGeoJSONFeatureProperties{
										Name:  "Slovensko",
										Rules: rules,
									},
								),
							},
						),
					},
				}
				feedSK.Language = "sk"
				feedEN := &gbfs.FeedGeofencingZones{
					Data: &gbfs.FeedGeofencingZonesData{
						GeofencingZones: gbfs.NewFeedGeofencingZonesGeoJSONFeatureCollection(
							[]*gbfs.FeedGeofencingZonesGeoJSONFeature{
								gbfs.NewFeedGeofencingZonesGeoJSONFeature(
									geometry,
									&gbfs.FeedGeofencingZonesGeoJSONFeatureProperties{
										Name:  "Slovakia",
										Rules: rules,
									},
								),
							},
						),
					},
				}
				feedEN.Language = "en"
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
	}
}

func main() {
	systemID := "system_id"
	s, err := gbfs.NewServer(&gbfs.ServerOptions{
		SystemID:     systemID,
		RootDir:      "public",
		BaseURL:      "http://127.0.0.1:8080",
		BasePath:     "v3/" + systemID,
		Version:      gbfs.V30,
		DefaultTTL:   60,
		FeedHandlers: getFeedHandlers(nil),
		UpdateHandler: func(s *gbfs.Server, feed gbfs.Feed, path string, err error) {
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("system=%s ttl=%d updated=%s", s.Options.SystemID, feed.GetTTL(), path)
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// err = os.RemoveAll(filepath.FromSlash(s.Options.RootDir + "/" + s.Options.BasePath))
	// if err != nil {
	// 	log.Println(err)
	// }
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
	go (func() {
		fs, err := gbfs.NewFileServer(&gbfs.FileServerOptions{
			Addr:    "127.0.0.1:8080",
			RootDir: "public",
		})
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(fs.ListenAndServe())
	})()
	// delay to start server
	time.Sleep(100 * time.Millisecond)
	c, err := gbfs.NewClient(&gbfs.ClientOptions{
		AutoDiscoveryURL: "http://127.0.0.1:8080/v3/" + systemID + "/gbfs.json",
		DefaultLanguage:  "en",
	})
	if err != nil {
		log.Fatal(err)
	}
	err = c.Subscribe(&gbfs.ClientSubscribeOptions{
		// Languages: []string{"en"},
		// FeedNames: []string{gbfs.FeedNameStationInformation, gbfs.FeedNameFreeBikeStatus},
		Handler: func(c *gbfs.Client, feed gbfs.Feed, err error) {
			if err != nil {
				log.Println(err)
				return
			}
			j, _ := json.Marshal(feed)
			log.Printf("feed=%s language=%s data=%s", feed.Name(), feed.GetLanguage(), j)
		},
	})
	if err != nil {
		log.Println(err)
	}
}
