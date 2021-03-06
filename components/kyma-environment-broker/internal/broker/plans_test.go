package broker

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestSchemaGenerator(t *testing.T) {
	tests := []struct {
		name         string
		generator    func([]string) []byte
		machineTypes []string
		want         string
	}{
		{
			name:         "Azure schema is correct",
			generator:    AzureSchema,
			machineTypes: []string{"Standard_D8_v3"},
			want: `{
			"$schema": "http://json-schema.org/draft-04/schema#",
			"type": "object",
			"properties": {
			"components": {
			"type": "array",
			"items": [
		{
			"type": "string",
			"enum": ["Kiali", "Tracing"]
		}
		],
			"additionalItems": false,
			"uniqueItems": true
		},
			"name": {
			"type": "string"
		},
			"diskType": {
			"type": "string"
		},
			"volumeSizeGb": {
			"type": "integer",
			"minimum": 50
		},
			"machineType": {
			"type": "string",
			"enum": ["Standard_D8_v3"]
		},
			"region": {
			"type": "string",
			"enum": [ "centralus", "eastus", "westus2", "northeurope", "uksouth", "japaneast", "southeastasia", "westeurope" ]
		},
			"zones": {
			"type": "array",
			"items": [
			{
			  "type": "string"
			}
			]
		},
			"autoScalerMin": {
			"type": "integer"
		},
			"autoScalerMax": {
			"type": "integer"
		},
			"maxSurge": {
			"type": "integer"
		},
			"maxUnavailable": {
			"type": "integer"
		}
		},
			"required": [
			"name"
		]
		}`},
		{
			name:         "AzureLite schema is correct",
			generator:    AzureSchema,
			machineTypes: []string{"Standard_D4_v3"},
			want: `{
			"$schema": "http://json-schema.org/draft-04/schema#",
			"type": "object",
			"properties": {
			"components": {
			"type": "array",
			"items": [
		{
			"type": "string",
			"enum": ["Kiali", "Tracing"]
		}
		],
			"additionalItems": false,
			"uniqueItems": true
		},
			"name": {
			"type": "string"
		},
			"diskType": {
			"type": "string"
		},
			"volumeSizeGb": {
			"type": "integer",
			"minimum": 50
		},
			"machineType": {
			"type": "string",
			"enum": ["Standard_D4_v3"]
		},
			"region": {
			"type": "string",
			"enum": [ "centralus", "eastus", "westus2", "northeurope", "uksouth", "japaneast", "southeastasia", "westeurope" ]
		},
			"zones": {
			"type": "array",
			"items": [
			{
			  "type": "string"
			}
			]
		},
			"autoScalerMin": {
			"type": "integer"
		},
			"autoScalerMax": {
			"type": "integer"
		},
			"maxSurge": {
			"type": "integer"
		},
			"maxUnavailable": {
			"type": "integer"
		}
		},
			"required": [
			"name"
		]
		}`},
		{
			name:         "GCP schema is correct",
			generator:    GCPSchema,
			machineTypes: []string{"n1-standard-2", "n1-standard-4", "n1-standard-8", "n1-standard-16", "n1-standard-32", "n1-standard-64"},
			want: `{
			"$schema": "http://json-schema.org/draft-04/schema#",
			"type": "object",
			"properties": {
			"components": {
			"type": "array",
			"items": [
		{
			"type": "string",
			"enum": ["Kiali", "Tracing"]
		}
		],
			"additionalItems": false,
			"uniqueItems": true
		},
			"name": {
			"type": "string"
		},
			"diskType": {
			"type": "string"
		},
			"volumeSizeGb": {
			"type": "integer"
		},
			"machineType": {
			"type": "string",
			"enum": ["n1-standard-2", "n1-standard-4", "n1-standard-8", "n1-standard-16", "n1-standard-32", "n1-standard-64"]
		},
			"region": {
			"type": "string",
			"enum": ["asia-south1", "asia-southeast1",
					"asia-east2", "asia-east1",
					"asia-northeast1", "asia-northeast2", "asia-northeast-3",
					"australia-southeast1",
					"europe-west2", "europe-west4", "europe-west5", "europe-west6", "europe-west3",
					"europe-north1",
					"us-west1", "us-west2", "us-west3",
					"us-central1",
					"us-east4",
					"northamerica-northeast1", "southamerica-east1"]
		},
			"zones": {
			"type": "array",
			"items": [
			{
				"type": "string",
				"enum": ["asia-south1-a", "asia-south1-b", "asia-south1-c",
						"asia-southeast1-a", "asia-southeast1-b", "asia-southeast1-c",
						"asia-east2-a", "asia-east2-b", "asia-east2-c",
						"asia-east1-a", "asia-east1-b", "asia-east1-c",
						"asia-northeast1-a", "asia-northeast1-b", "asia-northeast1-c",
						"asia-northeast2-a", "asia-northeast2-b", "asia-northeast2-c",
						"asia-northeast-3-a", "asia-northeast-3-b", "asia-northeast-3-c",
						"australia-southeast1-a", "australia-southeast1-b", "australia-southeast1-c",
						"europe-west2-a", "europe-west2-b", "europe-west2-c",
						"europe-west4-a", "europe-west4-b", "europe-west4-c",
						"europe-west5-a", "europe-west5-b", "europe-west5-c",
						"europe-west6-a", "europe-west6-b", "europe-west6-c",
						"europe-west3-a", "europe-west3-b", "europe-west3-c",
						"europe-north1-a", "europe-north1-b", "europe-north1-c",
						"us-west1-a", "us-west1-b", "us-west1-c",
						"us-west2-a", "us-west2-b", "us-west2-c",
						"us-west3-a", "us-west3-b", "us-west3-c",
						"us-central1-a", "us-central1-b", "us-central1-c",
						"us-east4-a", "us-east4-b", "us-east4-c",
						"northamerica-northeast1-a", "northamerica-northeast1-b", "northamerica-northeast1-c",
						"southamerica-east1-a", "southamerica-east1-b", "southamerica-east1-c"]
				}
			]
		},
			"autoScalerMin": {
			"type": "integer"
		},
			"autoScalerMax": {
			"type": "integer"
		},
			"maxSurge": {
			"type": "integer"
		},
			"maxUnavailable": {
			"type": "integer"
		}
		},
			"required": [
			"name"
		]
		}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.generator(tt.machineTypes)
			validateSchema(t, got, tt.want)

		})
	}
}

func TestTrialSchemaGenerator(t *testing.T) {
	wantGcp := `{
          "$schema": "http://json-schema.org/draft-04/schema#",
          "type": "object",
          "properties": {
            "name": {
              "type": "string"
            },
            "region": {
              "type": "string",
              "enum": [
                "europe-west4",
                "us-east4"
              ]
            },
            "zones": {
              "type": "array",
              "items": [
                {
                  "type": "string",
                  "enum": [
                    "europe-west4-a",
                    "europe-west4-b",
                    "europe-west4-c",
                    "us-east4-a",
                    "us-east4-b",
                    "us-east4-c"
                  ]
                }
              ]
            }
          },
          "required": [
            "name"
          ]
        }`

	gotGcp := GcpTrialSchema()
	validateSchema(t, gotGcp, wantGcp)

	wantAzure := `{
	             "$schema": "http://json-schema.org/draft-04/schema#",
	             "type": "object",
	             "properties": {
	               "name": {
	                 "type": "string"
	               },
	               "region": {
	                 "type": "string",
	                 "enum": [
	   				"eastus",
	   				"westeurope"
	                 ]
	               },
	               "zones": {
	                 "type": "array",
	                 "items": [
	                   {
	                     "type": "string"
	                   }
	                 ]
	               }
	             },
	             "required": [
	               "name"
	             ]
	           }`
	gotAzure := AzureTrialSchema()
	validateSchema(t, gotAzure, wantAzure)
}

func validateSchema(t *testing.T, got []byte, want string) {
	var prettyWant bytes.Buffer
	err := json.Indent(&prettyWant, []byte(want), "", "  ")
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	var prettyGot bytes.Buffer
	err = json.Indent(&prettyGot, got, "", "  ")
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if !reflect.DeepEqual(string(prettyGot.String()), prettyWant.String()) {
		t.Errorf("Schema() = \n######### GOT ###########%v\n######### ENDGOT ########, want \n##### WANT #####%v\n##### ENDWANT #####", prettyGot.String(), prettyWant.String())
	}
}
