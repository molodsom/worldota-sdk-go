package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Float64Str float64

func (v *Float64Str) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%f", *v)
	return json.Marshal(s)
}

func (v *Float64Str) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		value, _ := strconv.ParseFloat(s, 64)
		*v = Float64Str(value)
		return nil
	}
	return json.Unmarshal(b, (*float64)(v))
}

type Paragraphs struct {
	Paragraphs *[]string `json:"paragraphs"`
	Title      *string   `json:"title"`
}

type RgExt struct {
	Bathroom *int8 `json:"bathroom"`
	Bedding  *int8 `json:"bedding"`
	Capacity *int8 `json:"capacity"`
	Class    *int8 `json:"class"`
	Club     *int8 `json:"club"`
	Family   *int8 `json:"family"`
	Quality  *int8 `json:"quality"`
	Sex      *int8 `json:"sex"`
}

type Hotel struct {
	ID    *string `json:"id"`
	Rates *[]struct {
		MatchHash      *string       `json:"match_hash"`
		DailyPrices    *[]Float64Str `json:"daily_prices"`
		Meal           *string       `json:"meal"`
		PaymentOptions *struct {
			PaymentTypes *[]struct {
				Amount               *Float64Str `json:"amount"`
				ShowAmount           *Float64Str `json:"show_amount"`
				CurrencyCode         *string     `json:"currency_code"`
				ShowCurrencyCode     *string     `json:"show_currency_code"`
				By                   *string     `json:"by"`
				IsNeedCreditCardData *bool       `json:"is_need_credit_card_data"`
				IsNeedCVC            *bool       `json:"is_need_cvc"`
				Type                 *string     `json:"type"`
				TaxData              *struct {
					Taxes *[]struct {
						Name               *string  `json:"name"`
						IncludedBySupplier *bool    `json:"included_by_supplier"`
						Amount             *float64 `json:"amount,string"`
						CurrencyCode       *string  `json:"currency_code"`
					} `json:"taxes"`
				} `json:"tax_data"`
				VATData *struct {
					Included *bool       `json:"included"`
					Value    *Float64Str `json:"value"`
				} `json:"vat_data"`
				CancellationPenalties *struct {
					Policies *[]struct {
						StartAt      *string     `json:"start_at"`
						EndAt        *string     `json:"end_at"`
						AmountCharge *Float64Str `json:"amount_charge"`
						AmountShow   *Float64Str `json:"amount_show"`
					} `json:"policies"`
				} `json:"cancellation_penalties"`
				FreeCancellationBefore *string `json:"free_cancellation_before"`
			} `json:"payment_types"`
			RgExt         RgExt     `json:"rg_ext"`
			RoomName      *string   `json:"room_name"`
			SERPFilters   *[]string `json:"serp_filters"`
			Allotment     *int8     `json:"allotment"`
			AmenitiesData *[]string `json:"amenities_data"`
			AnyResidency  *bool     `json:"any_residency"`
			Deposit       *struct {
				Amount       *Float64Str `json:"amount"`
				CurrencyCode *string     `json:"currency_code"`
				IsRefundable *bool       `json:"is_refundable"`
			} `json:"deposit"`
			NoShow *struct {
				Amount       *Float64Str `json:"availability"`
				CurrencyCode *string     `json:"currency_code"`
				FromTime     *string     `json:"from_time"`
			} `json:"no_show"`
			RoomDataTrans *struct {
				MainRoomType *string `json:"main_room_type"`
				MainName     *string `json:"main_name"`
				Bathroom     *string `json:"bathroom"`
				BeddingType  *string `json:"bedding_type"`
				MiscRoomType *string `json:"misc_room_type"`
			}
		} `json:"payment_options"`
	} `json:"rates"`
}

type Overview struct {
	Data []struct {
		Endpoint       *string `json:"endpoint"`
		IsActive       *bool   `json:"is_active"`
		IsDebugMode    *bool   `json:"is_debug_mode"`
		IsLimited      *bool   `json:"is_limited"`
		RequestsNumber *int    `json:"requests_number"`
		SecondsNumber  *int    `json:"seconds_number"`
	} `json:"data"`
	Error  string `json:"error"`
	Status string `json:"status"`
}

type HotelDump struct {
	Data struct {
		LastUpdate *string `json:"last_update"`
		URL        *string `json:"url"`
	} `json:"data"`
	Error  string `json:"error"`
	Status string `json:"status"`
}

type HotelInfo struct {
	Data struct {
		Address       *string `json:"address"`
		AmenityGroups *[]struct {
			Amenities *[]string `json:"amenities"`
			GroupName *string   `json:"group_name"`
		} `json:"amenity_groups"`
		CheckInTime       *string      `json:"check_in_time"`
		CheckOutTime      *string      `json:"check_out_time"`
		DescriptionStruct []Paragraphs `json:"description_struct"`
		Email             *string      `json:"email"`
		Facts             *struct {
			Electricity *struct {
				Frequency []int16   `json:"frequency"`
				Sockets   *[]string `json:"sockets"`
				Voltage   []int16
			} `json:"electricity"`
			FloorsNumber  *int16 `json:"floors_number"`
			RoomsNumber   *int16 `json:"rooms_number"`
			YearBuilt     *int16 `json:"year_built"`
			YearRenovated *int16 `json:"year_renovated"`
		} `json:"facts"`
		HotelChain          *string   `json:"hotel_chain"`
		ID                  *string   `json:"id"`
		Images              *[]string `json:"images"`
		IsClosed            *bool     `json:"isClosed"`
		Kind                *string   `json:"kind"`
		Latitude            float32   `json:"latitude"`
		Longitude           float32   `json:"longitude"`
		MetapolicyExtraInfo *string   `json:"metapolicy_extra_info"`
		MetapolicyStruct    *struct {
			AddFee *[]struct {
				Currency  *string     `json:"currency"`
				FeeType   *string     `json:"fee_type"`
				Price     *Float64Str `json:"price"`
				PriceUnit *string     `json:"price_unit"`
			} `json:"add_fee"`
			Children     *[]string `json:"children"`
			ChildrenMeal *[]struct {
				AgeEnd    *int8       `json:"age_end"`
				AgeStart  *int8       `json:"age_start"`
				Currency  *string     `json:"currency"`
				Inclusion *string     `json:"inclusion"`
				MealType  *string     `json:"meal_type"`
				Price     *Float64Str `json:"price"`
			} `json:"children_meal"`
			CheckInCheckOut *[]struct {
				CheckInCheckOutType *string     `json:"check_in_check_out_type"`
				Currency            *string     `json:"currency"`
				Inclusion           *string     `json:"inclusion"`
				Price               *Float64Str `json:"price"`
			} `json:"check_in_check_out"`
			COT *[]struct {
				Amount    *int8       `json:"amount"`
				Currency  *string     `json:"currency"`
				Inclusion *string     `json:"inclusion"`
				Price     *Float64Str `json:"price"`
				PriceUnit *string     `json:"price_unit"`
			} `json:"cot"`
			Deposit *[]struct {
				Availability  *string     `json:"availability"`
				Currency      *string     `json:"currency"`
				DepositType   *string     `json:"deposit_type"`
				PaymentType   *string     `json:"payment_type"`
				Price         *Float64Str `json:"price"`
				PriceUnit     *string     `json:"price_unit"`
				PricingMethod *string     `json:"pricing_method"`
			} `json:"deposit"`
			ExtraBed *[]struct {
				Amount    *int8       `json:"amount"`
				Currency  *string     `json:"currency"`
				Inclusion *string     `json:"inclusion"`
				Price     *Float64Str `json:"price"`
				PriceUnit *string     `json:"price_unit"`
			} `json:"extra_bed"`
			Internet *[]struct {
				Currency     *string     `json:"currency"`
				Inclusion    *string     `json:"inclusion"`
				InternetType *string     `json:"internet_type"`
				Price        *Float64Str `json:"price"`
				PriceUnit    *string     `json:"price_unit"`
				WorkArea     *string     `json:"work_area"`
			} `json:"internet"`
			Meal *[]struct {
				Currency  *string     `json:"currency"`
				Inclusion *string     `json:"inclusion"`
				MealType  *string     `json:"meal_type"`
				Price     *Float64Str `json:"price"`
			} `json:"meal"`
			NoShow *struct {
				Availability *string `json:"availability"`
				DayPeriod    *string `json:"day_period"`
				Time         *string `json:"time"`
			} `json:"no_show"`
			Parking *[]struct {
				Currency      *string     `json:"currency"`
				Inclusion     *string     `json:"inclusion"`
				Price         *Float64Str `json:"price"`
				PriceUnit     *string     `json:"price_unit"`
				TerritoryType *string     `json:"territory_type"`
			} `json:"parking"`
			Pets *[]struct {
				Currency  *string     `json:"currency"`
				Inclusion *string     `json:"inclusion"`
				PetsType  *string     `json:"pets_type"`
				Price     *Float64Str `json:"price"`
				PriceUnit *string     `json:"price_unit"`
			} `json:"pets"`
			Shuttle *[]struct {
				Currency        *string     `json:"currency"`
				DestinationType *string     `json:"destination_type"`
				Inclusion       *string     `json:"inclusion"`
				Price           *Float64Str `json:"price"`
				ShuttleType     *string     `json:"shuttle_type"`
			} `json:"shuttle"`
			Visa *[]struct {
				VisaSupport *string `json:"visa_support"`
			} `json:"visa"`
		} `json:"metapolicy_struct"`
		Name           *string       `json:"name"`
		PaymentMethods *string       `json:"payment_methods"`
		Phone          *string       `json:"phone"`
		PolicyStruct   *[]Paragraphs `json:"policy_struct"`
		PostalCode     *string       `json:"postal_code"`
		Region         *struct {
			CountryCode *string `json:"country_code"`
			IATA        *string `json:"iata"`
			ID          *int    `json:"id"`
			Name        *string `json:"name"`
			Type        *string `json:"type"`
		} `json:"region"`
		RoomGroups *[]struct {
			Images     *[]string `json:"images"`
			Name       *string   `json:"name"`
			NameStruct *struct {
				Bathroom    *string `json:"bathroom"`
				BeddingType *string `json:"bedding_type"`
				MainName    *string `json:"main_name"`
			} `json:"name_struct"`
			RgExt         RgExt     `json:"rg_ext"`
			RoomAmenities *[]string `json:"room_amenities"`
			RoomGroupId   *int16    `json:"room_group_id"`
		} `json:"room_groups"`
		SERPFilters     *[]string `json:"serp_filters"`
		StarCertificate *struct {
			ValidTo       *string `json:"valid_to"`
			CertificateID *string `json:"certificate_id"`
		} `json:"star_certificate"`
		StarRating *int8
	} `json:"data"`
	Error  string `json:"error"`
	Status string `json:"status"`
}

type SearchMulticomplete *struct {
	Data struct {
		Hotels *[]struct {
			ID       *string `json:"id"`
			Name     *string `json:"name"`
			RegionId *string `json:"region_id"`
		} `json:"hotels"`
		Regions *[]struct {
			ID          *string `json:"id"`
			Name        *string `json:"name"`
			Type        *string `json:"type"`
			CountryCode *string `json:"country_code"`
		} `json:"regions"`
	}
	Error  string `json:"error"`
	Status string `json:"status"`
}

type SearchResult struct {
	Data struct {
		Hotels      *[]Hotel `json:"hotels"`
		TotalHotels *int     `json:"total_hotels"`
	} `json:"data"`
	Error  string `json:"error"`
	Status string `json:"status"`
}
