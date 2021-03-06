package prof

import (
	gq "github.com/graphql-go/graphql"
	"github.com/midorigreen/gprof/prof/cpu"
	"github.com/midorigreen/gprof/prof/disk"
	"github.com/midorigreen/gprof/prof/file"
)

var rootQuery = gq.NewObject(
	gq.ObjectConfig{
		Name: "RootQuery",
		Fields: gq.Fields{
			"cpu": &gq.Field{
				Type:        cpu.Type,
				Description: "List of cpus",
				Resolve:     cpu.Resolve,
			},
			"disk": &gq.Field{
				Type:        disk.Type,
				Description: "prof disk",
				Resolve:     disk.Resolve,
				Args: gq.FieldConfigArgument{
					"path": &gq.ArgumentConfig{
						Type: gq.String,
					},
				},
			},
			"file": &gq.Field{
				Type:        file.Type,
				Description: "tail file content",
				Resolve:     file.Resolve,
				Args: gq.FieldConfigArgument{
					"path": &gq.ArgumentConfig{
						Type: gq.String,
					},
					"num": &gq.ArgumentConfig{
						Type: gq.Int,
					},
				},
			},
		},
		Description: "Root",
	})

var Schema, _ = gq.NewSchema(
	gq.SchemaConfig{
		Query: rootQuery,
	},
)
