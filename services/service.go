package services

import (
	"github.com/SwanHtetAungPhyo/spi_cli/models"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"os"
)

func GatherAPIConfigInput(config *models.APIConfig) {

	config.Name = "MyAPIGateway"
	config.Version = "v1.0"
	config.Description = "This is a default API Gateway description"
	config.DefaultRoute = "/default"
	config.GatewayInfo = "http://myapi.gateway/info"
	config.Username = "admin"
	config.Password = "password123"
	config.LoadBalancing = "RoundRobin"
	config.MainApplication = "MainApp"
}

func GetServiceInput(config *models.APIConfig) {

	service := &models.Service{}

	_, err := models.Yellow.Println("Processing the service configuration ....")
	if err != nil {
		return
	}

	service.Name = "MyService"
	service.URL = "http://localhost:8081"
	service.Leader = "John Doe"

	service.Instance = []int{8081, 8082, 8083}
	service.Routes = []models.Route{
		{
			Path:        []string{"/route1", "/route2"},
			Method:      []string{"GET", "POST"},
			Description: "Sample route 1 description",
			Timeout:     "30s",
			Retries:     3,
			GeneratedRoute: []string{
				"http://localhost:8081/gate/route1/",
				"http://localhost:8081/gate/route2/",
			},
			RealUrl: []string{
				"http://localhost:3000/route1/",
				"http://localhost:3000/route2/",
			},
		},
	}

	config.Services = append(config.Services, *service)
}

func GetInstances() []int {

	return []int{8081, 8082, 8083}
}

func GetRoutes() []models.Route {

	return []models.Route{
		{
			Path:        []string{"/route1", "/route2"},
			Method:      []string{"GET", "POST"},
			Description: "Sample route 1 description",
			Timeout:     "30s",
			Retries:     3,
			GeneratedRoute: []string{
				"http://localhost:8081/gate/route1/",
				"http://localhost:8081/gate/route2/",
			},
			RealUrl: []string{
				"http://localhost:3000/route1/",
				"http://localhost:3000/route2/",
			},
		},
	}
}
func GenerateYaml(config any) ([]byte, error) {
	data, err := yaml.Marshal(&config)
	if err != nil {
		return nil, err
	}
	return data, err
}

func SaveToFile(filepath string, data []byte) error {
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func ColorErrorHandle(c *color.Color, s string) {
	_, err := c.Println(s)
	if err != nil {
		return
	}
}
