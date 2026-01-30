package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
	"github.com/paula-dot/knbs-open-data-api/backend/internal/database"
)

// Helper to convert float to pgtype.Numeric
func floatToNumeric(val float64) pgtype.Numeric {
	var n pgtype.Numeric
	// Use Scan which accepts a string and populates Numeric
	if err := n.Scan(fmt.Sprintf("%f", val)); err != nil {
		// If conversion fails, return an invalid Numeric
		return pgtype.Numeric{}
	}
	return n
}

func main() {
	// 1. Load Env
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("Warning: ../../.env file not found, assuming env vars are set")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	// 2. Connect
	ctx := context.Background()
	pool, err := database.NewConnection(dbURL)
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer pool.Close()

	db := database.New(pool)

	// 3. The Data (Canonical List)
	counties := []struct {
		ID             int32
		Name           string
		Code           string
		FormerProvince string
		Area           float64
	}{
		{1, "Mombasa", "001", "Coast", 212.5},
		{2, "Kwale", "002", "Coast", 8270.3},
		{3, "Kilifi", "003", "Coast", 12245.9},
		{4, "Tana River", "004", "Coast", 35375.8},
		{5, "Lamu", "005", "Coast", 6497.7},
		{6, "Taita Taveta", "006", "Coast", 17083.9},
		{7, "Garissa", "007", "North Eastern", 45720.2},
		{8, "Wajir", "008", "North Eastern", 55840.6},
		{9, "Mandera", "009", "North Eastern", 25797.7},
		{10, "Marsabit", "010", "Eastern", 66923.1},
		{11, "Isiolo", "011", "Eastern", 25336.1},
		{12, "Meru", "012", "Eastern", 7003.1},
		{13, "Tharaka-Nithi", "013", "Eastern", 2609.5},
		{14, "Embu", "014", "Eastern", 2818.0},
		{15, "Kitui", "015", "Eastern", 24385.1},
		{16, "Machakos", "016", "Eastern", 5952.9},
		{17, "Makueni", "017", "Eastern", 8008.9},
		{18, "Nyandarua", "018", "Central", 3107.7},
		{19, "Nyeri", "019", "Central", 2361.0},
		{20, "Kirinyaga", "020", "Central", 1205.4},
		{21, "Murang'a", "021", "Central", 2325.8},
		{22, "Kiambu", "022", "Central", 2449.2},
		{23, "Turkana", "023", "Rift Valley", 71597.8},
		{24, "West Pokot", "024", "Rift Valley", 8418.2},
		{25, "Samburu", "025", "Rift Valley", 20182.5},
		{26, "Trans Nzoia", "026", "Rift Valley", 2469.9},
		{27, "Uasin Gishu", "027", "Rift Valley", 2955.3},
		{28, "Elgeyo-Marakwet", "028", "Rift Valley", 3049.7},
		{29, "Nandi", "029", "Rift Valley", 2884.5},
		{30, "Baringo", "030", "Rift Valley", 11075.3},
		{31, "Laikipia", "031", "Rift Valley", 8696.1},
		{32, "Nakuru", "032", "Rift Valley", 7509.5},
		{33, "Narok", "033", "Rift Valley", 17921.2},
		{34, "Kajiado", "034", "Rift Valley", 21292.7},
		{35, "Kericho", "035", "Rift Valley", 2454.5},
		{36, "Bomet", "036", "Rift Valley", 1997.9},
		{37, "Kakamega", "037", "Western", 3033.8},
		{38, "Vihiga", "038", "Western", 531.3},
		{39, "Bungoma", "039", "Western", 2206.9},
		{40, "Busia", "040", "Western", 1628.4},
		{41, "Siaya", "041", "Nyanza", 2496.1},
		{42, "Kisumu", "042", "Nyanza", 2009.5},
		{43, "Homa Bay", "043", "Nyanza", 3154.7},
		{44, "Migori", "044", "Nyanza", 2586.4},
		{45, "Kisii", "045", "Nyanza", 1317.9},
		{46, "Nyamira", "046", "Nyanza", 912.5},
		{47, "Nairobi City", "047", "Nairobi", 694.9},
	}
	fmt.Println("Seeding Counties...")

	for _, c := range counties {
		// Prepare the args for sqlc
		args := database.CreateCountyParams{
			ID:   c.ID,
			Name: c.Name,
			Code: c.Code,
			FormerProvince: pgtype.Text{
				String: c.FormerProvince,
				Valid:  true,
			},
			AreaSqKm: floatToNumeric(c.Area),
		}
		// Insert the county
		_, err := db.CreateCounty(ctx, args)
		if err != nil {
			// If duplicate, just log and continue
			fmt.Printf("⚠️  Skipping %s (might already exist): %v\n", c.Name, err)
			continue
		}
		fmt.Printf("✅ Created %s\n", c.Name)
	}

	// 1. Create the Indicator
	fmt.Println("📊 Seeding Indicators...")
	popIndicator, err := db.CreateIndicator(ctx, database.CreateIndicatorParams{
		Code:        "POP_TOTAL",
		Name:        "Total Population",
		Description: pgtype.Text{String: "Total enumeration from 2019 Census", Valid: true},
		Unit:        "People",
		Source:      "KNBS 2019 Census",
	})
	if err != nil {
		// If it exists, we need to fetch it to get the ID (omitted for brevity in this snippet)
		fmt.Println("⚠️  Indicator likely exists, skipping...")
	} else {
		fmt.Printf("✅ Created Indicator: %s\n", popIndicator.Name)
	}

	// 2. We need the ID of the indicator we just created/fetched
	// Ideally, fetch it by code to be safe:
	ind, _ := db.GetIndicatorByCode(ctx, "POP_TOTAL")

	// 3. Seed Observations (Sample Data)
	// Map County Code ("001") -> Population Value
	popData := map[string]float64{
		"001": 1208333, // Mombasa
		"047": 4397073, // Nairobi
		"032": 2162202, // Nakuru
		"022": 2417735, // Kiambu
		"027": 1163186, // Uasin Gishu
		// ... add more if you like
	}

	fmt.Println("📈 Seeding Observations...")
	for code, value := range popData {
		// Simpler approach for this snippet: parse code to int since our IDs match the codes
		var countyID int32
		fmt.Sscanf(code, "%d", &countyID)

		_, err := db.CreateObservation(ctx, database.CreateObservationParams{
			CountyID:       countyID,
			IndicatorID:    ind.ID,
			Year:           2019,
			Value:          floatToNumeric(value),
			SourceDocument: pgtype.Text{String: "Volume I", Valid: true},
		})
		if err != nil {
			// If insert fails (duplicate, FK, etc.) just continue
			fmt.Printf("⚠️  Skipping observation for county %s: %v\n", code, err)
			continue
		}
		fmt.Printf("   -> Added data for County %s: %.0f\n", code, value)
	}

	fmt.Println("🏁 Seeding complete!")
}
