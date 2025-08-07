package configs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"gopkg.in/yaml.v3"
)

type Settings struct {
	Simulation struct {
		TickIntervalMs int `yaml:"tick_interval_ms"`
		MaxPopulation  int `yaml:"max_population"`
		StartYear      int `yaml:"start_year"`
		CurrentYear    int `yaml:"current_year"` // Simülasyonun şu anki yılı, başlangıçta start_year ile aynı
	} `yaml:"simulation"`

	World struct {
		Disasters struct {
			EarthquakeChance float64 `yaml:"earthquake_chance"`
			EpidemicChance   float64 `yaml:"epidemic_chance"`
			WarChance        float64 `yaml:"war_chance"`
		} `yaml:"disasters"`

		Temperature struct {
			Min int `yaml:"min"`
			Max int `yaml:"max"`
		} `yaml:"temperature"`
	} `yaml:"world"`

	Human struct {
		MaxAge            int     `yaml:"max_age"`
		FertilityAgeStart int     `yaml:"fertility_age_start"`
		FertilityAgeEnd   int     `yaml:"fertility_age_end"`
		AgingRate         float64 `yaml:"aging_rate"`
	} `yaml:"human"`

	Attributes struct {
		Intelligence struct {
			Min int `yaml:"min"`
			Max int `yaml:"max"`
		} `yaml:"intelligence"`

		Strength struct {
			Min int `yaml:"min"`
			Max int `yaml:"max"`
		} `yaml:"strength"`

		Empathy struct {
			Min int `yaml:"min"`
			Max int `yaml:"max"`
		} `yaml:"empathy"`

		Creativity struct {
			Min int `yaml:"min"`
			Max int `yaml:"max"`
		} `yaml:"creativity"`

		Resilience struct {
			Min int `yaml:"min"`
			Max int `yaml:"max"`
		} `yaml:"resilience"`
	} `yaml:"attributes"`

	Needs struct {
		HungerMax int `yaml:"hunger_max"`
		SleepMax  int `yaml:"sleep_max"`
		SocialMax int `yaml:"social_max"`
		HygieneMax int `yaml:"hygiene_max"`
	} `yaml:"needs"`

	Skills struct {
		CommunicationMax int `yaml:"communication_max"`
		CraftingMax      int `yaml:"crafting_max"`
		FarmingMax       int `yaml:"farming_max"`
	} `yaml:"skills"`

	Mutation struct {
		Chance float64 `yaml:"chance"`
		Impact float64 `yaml:"impact"`
	} `yaml:"mutation"`

	Naming struct {
		MaleNames   []string
		FemaleNames []string
		Surnames    []string 
	}
}

var LoadedSettings Settings

func LoadSettings(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read settings file: %w", err)
	}

	err = yaml.Unmarshal(data, &LoadedSettings)
	if err != nil {
		return fmt.Errorf("failed to parse settings yaml: %w", err)
	}

	// İsim dosyalarını yükle
	err = loadNameFiles()
	if err != nil {
		return fmt.Errorf("failed to load name files: %w", err)
	}

	return nil
}

// Dosyadan satır satır okur ve isim listesini döner
func loadNamesFromFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var names []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			names = append(names, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return names, nil
}

// İsim dosyalarını assets klasöründen okur ve LoadedSettings.Naming içine koyar
func loadNameFiles() error {
	maleNames, err := loadNamesFromFile("../assets/male_names.txt")
	if err != nil {
		return err
	}
	femaleNames, err := loadNamesFromFile("../assets/female_names.txt")
	if err != nil {
		return err
	}

	LoadedSettings.Naming.MaleNames = maleNames
	LoadedSettings.Naming.FemaleNames = femaleNames

	return nil
}
