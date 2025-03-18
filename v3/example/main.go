package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/petoc/gbfs/v3"
)

func feedHandlers() []*gbfs.FeedHandler {
	return []*gbfs.FeedHandler{
		{
			// TTL: 60,
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feed := &gbfs.FeedSystemInformation{
					Data: &gbfs.FeedSystemInformationData{
						SystemID:  gbfs.NewID(s.Options.SystemID),
						Languages: []string{"en", "sk"},
						Name: []*gbfs.LocalizedString{
							gbfs.NewLocalizedString("Bike Sharing", "en"),
							gbfs.NewLocalizedString("Bike Sharing", "sk"),
						},
						Operator: []*gbfs.LocalizedString{
							gbfs.NewLocalizedString("Bike Sharing, Street 123, 12345 City", "en"),
							gbfs.NewLocalizedString("Bike Sharing, Street 123, 12345 City", "sk"),
						},
						URL:         gbfs.NewString("http://localhost/bikesharing/sk"),
						PhoneNumber: gbfs.NewString("+421987654321"),
						Email:       gbfs.NewString("bikesharing@example.com"),
						Timezone:    gbfs.NewString("Europe/Bratislava"),
						LicenseID:   gbfs.NewString("MIT"),
					},
				}
				// feed.TTL = 60
				return []gbfs.Feed{feed}, nil
			},
		},
		{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feed := &gbfs.FeedVehicleTypes{
					Data: &gbfs.FeedVehicleTypesData{
						VehicleTypes: []*gbfs.FeedVehicleTypesVehicleType{
							{
								VehicleTypeID:  gbfs.NewID("vehicleType1"),
								FormFactor:     gbfs.NewString(gbfs.FormFactorBicycle),
								PropulsionType: gbfs.NewString(gbfs.PropulsionTypeHuman),
								Name: []*gbfs.LocalizedString{
									gbfs.NewLocalizedString("Bicycle", "en"),
									gbfs.NewLocalizedString("Bicykel", "sk"),
								},
							},
							{
								VehicleTypeID:  gbfs.NewID("vehicleType2"),
								FormFactor:     gbfs.NewString(gbfs.FormFactorMoped),
								PropulsionType: gbfs.NewString(gbfs.PropulsionTypeElectric),
								Name: []*gbfs.LocalizedString{
									gbfs.NewLocalizedString("Moped", "en"),
									gbfs.NewLocalizedString("Skúter", "sk"),
								},
								MaxRangeMeters: gbfs.NewFloat64(100000),
							},
						},
					},
				}
				return []gbfs.Feed{feed}, nil
			},
		},
		{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feed := &gbfs.FeedStationInformation{
					Data: &gbfs.FeedStationInformationData{
						Stations: []*gbfs.FeedStationInformationStation{
							{
								StationID: gbfs.NewID("station1"),
								Name: []*gbfs.LocalizedString{
									gbfs.NewLocalizedString("Station", "en"),
									gbfs.NewLocalizedString("Stanica", "sk"),
								},
								Lat:     gbfs.NewCoordinate(48.1234),
								Lon:     gbfs.NewCoordinate(21.1234),
								Address: gbfs.NewString("Ulica 123"),
								StationArea: gbfs.NewGeoJSONGeometryMultiPolygon(
									[][][][]float64{
										{
											{
												{16.8331891, 47.7314286},
												{22.56571, 47.7314286},
												{22.56571, 49.6138162},
												{16.8331891, 49.6138162},
												{16.8331891, 47.7314286},
											},
										},
									},
									nil,
								),
								VehicleTypesCapacity: []*gbfs.VehicleTypesCapacity{
									gbfs.NewVehicleTypesCapacity([]string{"vehicleType1", "vehicleType2"}, 10),
								},
								VehicleDocksCapacity: []*gbfs.VehicleTypesCapacity{
									gbfs.NewVehicleTypesCapacity([]string{"vehicleType1", "vehicleType2"}, 10),
								},
							},
						},
					},
				}
				return []gbfs.Feed{feed}, nil
			},
		},
		{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feed := &gbfs.FeedStationStatus{
					Data: &gbfs.FeedStationStatusData{
						Stations: []*gbfs.FeedStationStatusStation{
							{
								StationID:            gbfs.NewID("station1"),
								NumVehiclesAvailable: gbfs.NewInt64(2),
								NumVehiclesDisabled:  gbfs.NewInt64(0),
								NumDocksAvailable:    gbfs.NewInt64(0),
								NumDocksDisabled:     gbfs.NewInt64(0),
								IsInstalled:          gbfs.NewBoolean(true),
								IsRenting:            gbfs.NewBoolean(true),
								IsReturning:          gbfs.NewBoolean(true),
								LastReported:         gbfs.NewTimestamp("2020-01-01T00:00:00+00:00"),
								VehicleTypesAvailable: []*gbfs.VehicleTypeCapacity{
									gbfs.NewVehicleTypeCapacity("vehicleType1", 0),
								},
								VehicleDocksAvailable: []*gbfs.VehicleTypesCapacity{
									gbfs.NewVehicleTypesCapacity([]string{"vehicleType1"}, 0),
								},
							},
						},
					},
				}
				return []gbfs.Feed{feed}, nil
			},
		},
		{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feed := &gbfs.FeedVehicleStatus{
					Data: &gbfs.FeedVehicleStatusData{
						Vehicles: []*gbfs.FeedVehicleStatusVehicle{
							{
								VehicleID:     gbfs.NewID("bicyle1"),
								Lat:           gbfs.NewCoordinate(0),
								Lon:           gbfs.NewCoordinate(0),
								VehicleTypeID: gbfs.NewID("vehicleType1"),
								IsReserved:    gbfs.NewBoolean(false),
								IsDisabled:    gbfs.NewBoolean(false),
								LastReported:  gbfs.NewTimestamp("2020-01-01T00:00:00+00:00"),
							},
							{
								VehicleID:          gbfs.NewID("moped1"),
								Lat:                gbfs.NewCoordinate(48.7162),
								Lon:                gbfs.NewCoordinate(21.2613),
								VehicleTypeID:      gbfs.NewID("vehicleType2"),
								IsReserved:         gbfs.NewBoolean(false),
								IsDisabled:         gbfs.NewBoolean(false),
								LastReported:       gbfs.NewTimestamp("2020-01-01T00:00:00+00:00"),
								CurrentRangeMeters: gbfs.NewFloat64(12345.67),
							},
						},
					},
				}
				return []gbfs.Feed{feed}, nil
			},
		},
		{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feed := &gbfs.FeedSystemRegions{
					Data: &gbfs.FeedSystemRegionsData{
						Regions: []*gbfs.FeedSystemRegionsRegion{
							{
								RegionID: gbfs.NewID("region1"),
								Name: []*gbfs.LocalizedString{
									gbfs.NewLocalizedString("Region Name 1", "en"),
									gbfs.NewLocalizedString("Názov regiónu 1", "sk"),
								},
							},
							{
								RegionID: gbfs.NewID("region2"),
								Name: []*gbfs.LocalizedString{
									gbfs.NewLocalizedString("Region Name 2", "en"),
									gbfs.NewLocalizedString("Názov regiónu 2", "sk"),
								},
							},
						},
					},
				}
				return []gbfs.Feed{feed}, nil
			},
		},
		{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feed := &gbfs.FeedSystemPricingPlans{
					Data: &gbfs.FeedSystemPricingPlansData{
						Plans: []*gbfs.FeedSystemPricingPlansPricingPlan{
							{
								PlanID: gbfs.NewID("plan1"),
								Name: []*gbfs.LocalizedString{
									gbfs.NewLocalizedString("Price plan", "en"),
									gbfs.NewLocalizedString("Cenový plán", "sk"),
								},
								Currency:  gbfs.NewString("EUR"),
								Price:     gbfs.NewPrice(12.34),
								IsTaxable: gbfs.NewBoolean(false),
								Description: []*gbfs.LocalizedString{
									gbfs.NewLocalizedString("Price plan description", "en"),
									gbfs.NewLocalizedString("Popis cenového plánu", "sk"),
								},
							},
						},
					},
				}
				return []gbfs.Feed{feed}, nil
			},
		},
		{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feed := &gbfs.FeedSystemAlerts{
					Data: &gbfs.FeedSystemAlertsData{
						Alerts: []*gbfs.FeedSystemAlertsAlert{
							{
								AlertID: gbfs.NewID("alert1"),
								Type:    gbfs.NewString(gbfs.AlertTypeSystemClosure),
								Times: []*gbfs.FeedSystemAlertsAlertTime{
									{
										Start: gbfs.NewTimestamp("2020-01-01T08:00:00+00:00"),
										End:   gbfs.NewTimestamp("2020-01-01T20:00:00+00:00"),
									},
								},
								StationIDs: []*gbfs.ID{gbfs.NewID("station1")},
								RegionIDs:  []*gbfs.ID{gbfs.NewID("region1")},
								URL: []*gbfs.LocalizedString{
									gbfs.NewLocalizedString("http://localhost/en/alerts/alert1", "en"),
									gbfs.NewLocalizedString("http://localhost/sk/alerts/alert1", "sk"),
								},
								Summary: []*gbfs.LocalizedString{
									gbfs.NewLocalizedString("Alert summary", "en"),
									gbfs.NewLocalizedString("Zhrnutie upozornenia", "sk"),
								},
								Description: []*gbfs.LocalizedString{
									gbfs.NewLocalizedString("Alert description", "en"),
									gbfs.NewLocalizedString("Popis upozornenia", "sk"),
								},
								LastUpdated: gbfs.NewTimestamp("2020-01-01T00:00:00+00:00"),
							},
						},
					},
				}
				return []gbfs.Feed{feed}, nil
			},
		},
		{
			Handler: func(s *gbfs.Server) ([]gbfs.Feed, error) {
				feed := &gbfs.FeedGeofencingZones{
					Data: &gbfs.FeedGeofencingZonesData{
						GeofencingZones: gbfs.NewFeedGeofencingZonesGeoJSONFeatureCollection(
							[]*gbfs.FeedGeofencingZonesGeoJSONFeature{
								gbfs.NewFeedGeofencingZonesGeoJSONFeature(
									gbfs.NewGeoJSONGeometryMultiPolygon(
										[][][][]float64{
											{
												{
													{16.8331891, 47.7314286},
													{22.56571, 47.7314286},
													{22.56571, 49.6138162},
													{16.8331891, 49.6138162},
													{16.8331891, 47.7314286},
												},
											},
										},
										nil,
									),
									&gbfs.FeedGeofencingZonesGeoJSONFeatureProperties{
										Name: []*gbfs.LocalizedString{
											gbfs.NewLocalizedString("Slovakia", "en"),
											gbfs.NewLocalizedString("Slovensko", "sk"),
										},
										Rules: []*gbfs.FeedGeofencingZonesRule{
											{
												VehicleTypeIDs: []*gbfs.ID{
													gbfs.NewID("vehicleType1"),
													gbfs.NewID("vehicleType2"),
												},
												RideStartAllowed:   gbfs.NewBoolean(true),
												RideEndAllowed:     gbfs.NewBoolean(true),
												RideThroughAllowed: gbfs.NewBoolean(true),
												MaximumSpeedKph:    gbfs.NewInt64(15),
											},
										},
									},
								),
							},
						),
					},
				}
				return []gbfs.Feed{feed}, nil
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
		FeedHandlers: feedHandlers(),
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
	})
	if err != nil {
		log.Fatal(err)
	}
	f := &gbfs.FeedSystemInformation{}
	err = c.Get(f)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Subscribe(gbfs.ClientSubscribeOptions{
		// FeedNames: []string{gbfs.FeedNameStationInformation, gbfs.FeedNameVehicleStatus},
		Handler: func(c *gbfs.Client, feed gbfs.Feed, err error) {
			if err != nil {
				log.Println(err)
				return
			}
			j, _ := json.Marshal(feed)
			log.Printf("feed=%s data=%s", feed.Name(), j)
		},
	})
	if err != nil {
		log.Println(err)
	}
}
