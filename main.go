package main

import (
	"github.com/SwanHtetAungPhyo/spi_cli/models"
	"github.com/SwanHtetAungPhyo/spi_cli/services"
	"github.com/fatih/color"
	"time"
)

func main() {

	_, _ = models.Blue.Println("Welcome to the SPI Gateway Configuration CLI Tool")

	config := models.APIConfig{}
	services.GatherAPIConfigInput(&config)
	services.GetServiceInput(&config)

	gatewayTag := "gateway"
	gateway := models.Gateway{
		GatewayTag: gatewayTag,
		APIConfig:  config,
	}

	yamlFile, err := services.GenerateYaml(&gateway)
	if err != nil {
		services.ColorErrorHandle(color.New(color.FgRed), "Error generating YAML: "+err.Error())
		return
	}

	err = services.SaveToFile("api_gateway_config.yaml", yamlFile)
	_, err = models.Blue.Println("Saving into the file......")
	if err != nil {
		return
	}
	time.Sleep(1 * time.Second)
	if err != nil {
		services.ColorErrorHandle(color.New(color.FgRed), "Error saving to file: "+err.Error())
		return
	}
	time.Sleep(1 * time.Second)
	services.ColorErrorHandle(color.New(color.FgGreen), "Configuration generated successfully and saved to api_gateway_config.yaml")
	services.ColorErrorHandle(color.New(color.FgHiMagenta), "All of the default configuration are in the api_gateway_config.yaml file. You can add the services that you want like the default")
}
