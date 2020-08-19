package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/petoc/gbfs"
	"github.com/petoc/gbfs/validator"
)

func getFeedHandlers(db *sql.DB) []*gbfs.FeedHandler {
	return []*gbfs.FeedHandler{
		&gbfs.FeedHandler{
			// TTL: 60,
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feedSK := &gbfs.FeedSystemInformation{
					Data: &gbfs.FeedSystemInformationData{
						SystemID:    gbfs.NewID(s.Options.SystemID),
						Language:    gbfs.NewString("sk"),
						Name:        gbfs.NewString("Bike Sharing"),
						Operator:    gbfs.NewString("Bike Sharing, Street 123, 12345 City"),
						URL:         gbfs.NewString("http://localhost/bikesharing/sk"),
						PhoneNumber: gbfs.NewString("00421987654321"),
						Email:       gbfs.NewString("bikesharing@example.com"),
						Timezone:    gbfs.NewString("Europe/Bratislava"),
						LicenseID:   gbfs.NewString("MIT"),
					},
				}
				feedSK.Language = gbfs.NewString("sk")
				// feedSK.TTL = 60
				feedEN := &gbfs.FeedSystemInformation{
					Data: &gbfs.FeedSystemInformationData{
						SystemID:    gbfs.NewID(s.Options.SystemID),
						Language:    gbfs.NewString("en"),
						Name:        gbfs.NewString("Bike Sharing"),
						Operator:    gbfs.NewString("Bike Sharing, Street 123, 12345 City"),
						URL:         gbfs.NewString("http://localhost/bikesharing/en"),
						PhoneNumber: gbfs.NewString("00421987654321"),
						Email:       gbfs.NewString("bikesharing@example.com"),
						Timezone:    gbfs.NewString("Europe/Bratislava"),
						LicenseID:   gbfs.NewString("MIT"),
					},
				}
				feedEN.Language = gbfs.NewString("en")
				// feedEN.TTL = 60
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feedSK := &gbfs.FeedVehicleTypes{
					Data: &gbfs.FeedVehicleTypesData{
						VehicleTypes: []*gbfs.FeedVehicleTypesVehicleType{
							&gbfs.FeedVehicleTypesVehicleType{
								VehicleTypeID:  gbfs.NewID("vehicleType1"),
								FormFactor:     gbfs.NewString(gbfs.FormFactorBicycle),
								PropulsionType: gbfs.NewString(gbfs.PropulsionTypeHuman),
								Name:           gbfs.NewString("Bicykel"),
							},
							&gbfs.FeedVehicleTypesVehicleType{
								VehicleTypeID:  gbfs.NewID("vehicleType2"),
								FormFactor:     gbfs.NewString(gbfs.FormFactorMoped),
								PropulsionType: gbfs.NewString(gbfs.PropulsionTypeElectric),
								Name:           gbfs.NewString("Skúter"),
							},
						},
					},
				}
				feedSK.Language = gbfs.NewString("sk")
				feedEN := &gbfs.FeedVehicleTypes{
					Data: &gbfs.FeedVehicleTypesData{
						VehicleTypes: []*gbfs.FeedVehicleTypesVehicleType{
							&gbfs.FeedVehicleTypesVehicleType{
								VehicleTypeID:  gbfs.NewID("vehicleType1"),
								FormFactor:     gbfs.NewString(gbfs.FormFactorBicycle),
								PropulsionType: gbfs.NewString(gbfs.PropulsionTypeHuman),
								Name:           gbfs.NewString("Bicycle"),
							},
							&gbfs.FeedVehicleTypesVehicleType{
								VehicleTypeID:  gbfs.NewID("vehicleType2"),
								FormFactor:     gbfs.NewString(gbfs.FormFactorMoped),
								PropulsionType: gbfs.NewString(gbfs.PropulsionTypeElectric),
								Name:           gbfs.NewString("Moped"),
							},
						},
					},
				}
				feedEN.Language = gbfs.NewString("en")
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				stationArea := gbfs.NewGeoJSONGeometryMultiPolygon(
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
				feedSK := &gbfs.FeedStationInformation{
					Data: &gbfs.FeedStationInformationData{
						Stations: []*gbfs.FeedStationInformationStation{
							&gbfs.FeedStationInformationStation{
								StationID:   gbfs.NewID("station1"),
								Name:        gbfs.NewString("Stanica"),
								Lat:         gbfs.NewCoordinate(48.1234),
								Lon:         gbfs.NewCoordinate(21.1234),
								Address:     gbfs.NewString("Ulica 123"),
								StationArea: stationArea,
							},
						},
					},
				}
				feedSK.Language = gbfs.NewString("sk")
				feedEN := &gbfs.FeedStationInformation{
					Data: &gbfs.FeedStationInformationData{
						Stations: []*gbfs.FeedStationInformationStation{
							&gbfs.FeedStationInformationStation{
								StationID:   gbfs.NewID("station1"),
								Name:        gbfs.NewString("Station"),
								Lat:         gbfs.NewCoordinate(48.1234),
								Lon:         gbfs.NewCoordinate(21.1234),
								Address:     gbfs.NewString("Street 123"),
								StationArea: stationArea,
							},
						},
					},
				}
				feedEN.Language = gbfs.NewString("en")
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				stations := []*gbfs.FeedStationStatusStation{
					&gbfs.FeedStationStatusStation{
						StationID:         gbfs.NewID("station1"),
						NumBikesAvailable: gbfs.NewInt64(2),
						NumBikesDisabled:  gbfs.NewInt64(0),
						NumDocksAvailable: gbfs.NewInt64(0),
						NumDocksDisabled:  gbfs.NewInt64(0),
						IsInstalled:       gbfs.NewBoolean(true),
						IsRenting:         gbfs.NewBoolean(true),
						IsReturning:       gbfs.NewBoolean(true),
						LastReported:      gbfs.NewTimestamp(1577836800),
						VehicleDocksAvailable: []*gbfs.FeedStationStatusVehicleDock{
							&gbfs.FeedStationStatusVehicleDock{
								VehicleTypeIDs: []*gbfs.ID{gbfs.NewID("vehicleType1")},
								Count:          gbfs.NewInt64(0),
							},
						},
						Vehicles: []*gbfs.FeedStationStatusVehicle{
							&gbfs.FeedStationStatusVehicle{
								BikeID:             gbfs.NewID("scooter2"),
								IsDisabled:         gbfs.NewBoolean(false),
								IsReserved:         gbfs.NewBoolean(false),
								VehicleTypeID:      gbfs.NewID("vehicleType2"),
								CurrentRangeMeters: gbfs.NewFloat64(100000),
							},
						},
					},
				}
				feedSK := &gbfs.FeedStationStatus{
					Data: &gbfs.FeedStationStatusData{
						Stations: stations,
					},
				}
				feedSK.Language = gbfs.NewString("sk")
				feedEN := &gbfs.FeedStationStatus{
					Data: &gbfs.FeedStationStatusData{
						Stations: stations,
					},
				}
				feedEN.Language = gbfs.NewString("en")
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				bicycle1 := &gbfs.FeedFreeBikeStatusBike{
					BikeID:        gbfs.NewID("bicyle1"),
					SystemID:      gbfs.NewID(s.Options.SystemID),
					Lat:           gbfs.NewCoordinate(0),
					Lon:           gbfs.NewCoordinate(0),
					VehicleTypeID: gbfs.NewID("vehicleType1"),
					IsReserved:    gbfs.NewBoolean(false),
					IsDisabled:    gbfs.NewBoolean(false),
					LastReported:  gbfs.NewTimestamp(1577836800),
				}
				moped1 := &gbfs.FeedFreeBikeStatusBike{
					BikeID:             gbfs.NewID("moped1"),
					SystemID:           gbfs.NewID(s.Options.SystemID),
					Lat:                gbfs.NewCoordinate(48.7162),
					Lon:                gbfs.NewCoordinate(21.2613),
					VehicleTypeID:      gbfs.NewID("vehicleType2"),
					IsReserved:         gbfs.NewBoolean(false),
					IsDisabled:         gbfs.NewBoolean(false),
					LastReported:       gbfs.NewTimestamp(1577836800),
					CurrentRangeMeters: gbfs.NewFloat64(12345.67),
				}
				feedSK := &gbfs.FeedFreeBikeStatus{
					Data: &gbfs.FeedFreeBikeStatusData{
						Bikes: []*gbfs.FeedFreeBikeStatusBike{
							bicycle1,
							moped1,
						},
					},
				}
				feedSK.Language = gbfs.NewString("sk")
				feedEN := &gbfs.FeedFreeBikeStatus{
					Data: &gbfs.FeedFreeBikeStatusData{
						Bikes: []*gbfs.FeedFreeBikeStatusBike{
							bicycle1,
							moped1,
						},
					},
				}
				feedEN.Language = gbfs.NewString("en")
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
							StartTime: gbfs.NewString("00:00:00"),
							EndTime:   gbfs.NewString("23:59:59"),
						},
					},
				}
				feedSK := &gbfs.FeedSystemHours{
					Data: data,
				}
				feedSK.Language = gbfs.NewString("sk")
				feedEN := &gbfs.FeedSystemHours{
					Data: data,
				}
				feedEN.Language = gbfs.NewString("en")
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				data := &gbfs.FeedSystemCalendarData{
					Calendars: []*gbfs.FeedSystemCalendarCalendar{
						&gbfs.FeedSystemCalendarCalendar{
							StartMonth: gbfs.NewInt64(1),
							StartDay:   gbfs.NewInt64(1),
							StartYear:  gbfs.NewInt64(2020),
							EndMonth:   gbfs.NewInt64(12),
							EndDay:     gbfs.NewInt64(31),
							EndYear:    gbfs.NewInt64(2020),
						},
					},
				}
				feedSK := &gbfs.FeedSystemCalendar{
					Data: data,
				}
				feedSK.Language = gbfs.NewString("sk")
				feedEN := &gbfs.FeedSystemCalendar{
					Data: data,
				}
				feedEN.Language = gbfs.NewString("en")
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				regions := []*gbfs.FeedSystemRegionsRegion{
					&gbfs.FeedSystemRegionsRegion{
						RegionID: gbfs.NewID("region1"),
						Name:     gbfs.NewString("Region Name 1"),
					},
					&gbfs.FeedSystemRegionsRegion{
						RegionID: gbfs.NewID("region2"),
						Name:     gbfs.NewString("Region Name 2"),
					},
				}
				feedSK := &gbfs.FeedSystemRegions{
					Data: &gbfs.FeedSystemRegionsData{
						Regions: regions,
					},
				}
				feedSK.Language = gbfs.NewString("sk")
				feedEN := &gbfs.FeedSystemRegions{
					Data: &gbfs.FeedSystemRegionsData{
						Regions: regions,
					},
				}
				feedEN.Language = gbfs.NewString("en")
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feedSK := &gbfs.FeedSystemPricingPlans{
					Data: &gbfs.FeedSystemPricingPlansData{
						Plans: []*gbfs.FeedSystemPricingPlansPricingPlan{
							&gbfs.FeedSystemPricingPlansPricingPlan{
								PlanID:      gbfs.NewID("plan1"),
								Name:        gbfs.NewString("Cenový plán"),
								Currency:    gbfs.NewString("EUR"),
								Price:       gbfs.NewPrice(12.34),
								IsTaxable:   gbfs.NewBoolean(false),
								Description: gbfs.NewString("Popis cenového plánu"),
							},
						},
					},
				}
				feedSK.Language = gbfs.NewString("sk")
				feedEN := &gbfs.FeedSystemPricingPlans{
					Data: &gbfs.FeedSystemPricingPlansData{
						Plans: []*gbfs.FeedSystemPricingPlansPricingPlan{
							&gbfs.FeedSystemPricingPlansPricingPlan{
								PlanID:      gbfs.NewID("plan1"),
								Name:        gbfs.NewString("Pricing Plan"),
								Currency:    gbfs.NewString("EUR"),
								Price:       gbfs.NewPrice(12.34),
								IsTaxable:   gbfs.NewBoolean(false),
								Description: gbfs.NewString("Pricing plan description"),
							},
						},
					},
				}
				feedEN.Language = gbfs.NewString("en")
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
		&gbfs.FeedHandler{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feedSK := &gbfs.FeedSystemAlerts{
					Data: &gbfs.FeedSystemAlertsData{
						Alerts: []*gbfs.FeedSystemAlertsAlert{
							&gbfs.FeedSystemAlertsAlert{
								AlertID: gbfs.NewID("alert1"),
								Type:    gbfs.NewString(gbfs.AlertTypeSystemClosure),
								Times: []*gbfs.FeedSystemAlertsAlertTime{
									&gbfs.FeedSystemAlertsAlertTime{
										Start: gbfs.NewTimestamp(1577865600),
										End:   gbfs.NewTimestamp(1577908800),
									},
								},
								StationIDs:  []*gbfs.ID{gbfs.NewID("station1")},
								RegionIDs:   []*gbfs.ID{gbfs.NewID("region1")},
								URL:         gbfs.NewString("http://localhost/sk/alerts/alert1"),
								Summary:     gbfs.NewString("Zhrnutie upozornenia"),
								Description: gbfs.NewString("Popis upozornenia"),
								LastUpdated: gbfs.NewTimestamp(1577836800),
							},
						},
					},
				}
				feedSK.Language = gbfs.NewString("sk")
				feedEN := &gbfs.FeedSystemAlerts{
					Data: &gbfs.FeedSystemAlertsData{
						Alerts: []*gbfs.FeedSystemAlertsAlert{
							&gbfs.FeedSystemAlertsAlert{
								AlertID: gbfs.NewID("alert1"),
								Type:    gbfs.NewString(gbfs.AlertTypeSystemClosure),
								Times: []*gbfs.FeedSystemAlertsAlertTime{
									&gbfs.FeedSystemAlertsAlertTime{
										Start: gbfs.NewTimestamp(1577865600),
										End:   gbfs.NewTimestamp(1577908800),
									},
								},
								StationIDs:  []*gbfs.ID{gbfs.NewID("station1")},
								RegionIDs:   []*gbfs.ID{gbfs.NewID("region1")},
								URL:         gbfs.NewString("http://localhost/en/alerts/alert1"),
								Summary:     gbfs.NewString("Alert summary"),
								Description: gbfs.NewString("Alert description"),
								LastUpdated: gbfs.NewTimestamp(1577836800),
							},
						},
					},
				}
				feedEN.Language = gbfs.NewString("en")
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
						VehicleTypeIDs: []*gbfs.ID{
							gbfs.NewID("vehicleType1"),
							gbfs.NewID("vehicleType2"),
						},
						RideAllowed:        gbfs.NewBoolean(true),
						RideThroughAllowed: gbfs.NewBoolean(true),
						MaximumSpeedKph:    gbfs.NewInt64(15),
					},
				}
				feedSK := &gbfs.FeedGeofencingZones{
					Data: &gbfs.FeedGeofencingZonesData{
						GeofencingZones: gbfs.NewFeedGeofencingZonesGeoJSONFeatureCollection(
							[]*gbfs.FeedGeofencingZonesGeoJSONFeature{
								gbfs.NewFeedGeofencingZonesGeoJSONFeature(
									geometry,
									&gbfs.FeedGeofencingZonesGeoJSONFeatureProperties{
										Name:  gbfs.NewString("Slovensko"),
										Rules: rules,
									},
								),
							},
						),
					},
				}
				feedSK.Language = gbfs.NewString("sk")
				feedEN := &gbfs.FeedGeofencingZones{
					Data: &gbfs.FeedGeofencingZonesData{
						GeofencingZones: gbfs.NewFeedGeofencingZonesGeoJSONFeatureCollection(
							[]*gbfs.FeedGeofencingZonesGeoJSONFeature{
								gbfs.NewFeedGeofencingZonesGeoJSONFeature(
									geometry,
									&gbfs.FeedGeofencingZonesGeoJSONFeatureProperties{
										Name:  gbfs.NewString("Slovakia"),
										Rules: rules,
									},
								),
							},
						),
					},
				}
				feedEN.Language = gbfs.NewString("en")
				return []gbfs.Feed{feedSK, feedEN}, nil
			},
		},
	}
}

func main() {
	systemID := "system_id"
	s, err := gbfs.NewServer(gbfs.ServerOptions{
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
			log.Printf("system=%s ttl=%d version=%s updated=%s", s.Options.SystemID, feed.GetTTL(), feed.GetVersion(), path)
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// err = os.RemoveAll(filepath.FromSlash(s.Options.RootDir + "/" + s.Options.BasePath))
	// if err != nil {
	// 	log.Println(err)
	// }
	go (func() {
		log.Fatal(s.Start())
	})()
	go (func() {
		fs, err := gbfs.NewFileServer("127.0.0.1:8080", "public")
		if err != nil {
			log.Fatal(err)
		}
		log.Fatal(fs.ListenAndServe())
	})()
	// delay to start server
	time.Sleep(100 * time.Millisecond)
	c, err := gbfs.NewClient(gbfs.ClientOptions{
		AutoDiscoveryURL: "http://127.0.0.1:8080/v3/" + systemID + "/gbfs.json",
		DefaultLanguage:  "en",
	})
	if err != nil {
		log.Fatal(err)
	}
	f := &gbfs.FeedSystemInformation{}
	err = c.Get(f)
	if err != nil {
		log.Fatal(err)
	}
	v := validator.New()
	r := v.Validate(f, gbfs.V10)
	log.Printf("infos=%v", r.Infos)
	log.Printf("warnings=%v", r.Warnings)
	log.Printf("errors=%v", r.Errors)
	err = c.Subscribe(gbfs.ClientSubscribeOptions{
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
