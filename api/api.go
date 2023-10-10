/*
Copyright © 2021 ax-i-om <addressaxiom@pm.me>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package api ...
package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/TwiN/go-color"
	"github.com/ax-i-om/bitcrook/pkg/authfree/ip2location"
	"github.com/ax-i-om/bitcrook/pkg/noauth/discord"
	"github.com/ax-i-om/bitcrook/pkg/noauth/ip"
	"github.com/ax-i-om/bitcrook/pkg/noauth/tin"
	"github.com/ax-i-om/bitcrook/pkg/noauth/userlookup"
	"github.com/ax-i-om/bitcrook/pkg/noauth/vin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func StartServer() {
	e := echo.New()

	e.HideBanner = true
	e.HidePort = true

	e.Use(middleware.Static("./api/static"))

	e.GET("/ip/:ip", func(c echo.Context) error {
		ipInfo, err := ip.IPLookup(c.Param("ip"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, ipInfo)
	})

	e.GET("/username/:username", func(c echo.Context) error {
		return c.JSON(http.StatusOK, userlookup.UserLookup(c.Param("username")))
	})

	e.GET("/vin/:vin", func(c echo.Context) error {
		vinInfo, err := vin.VinLookup(c.Param("vin"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, vinInfo)
	})

	e.GET("/discord/:discord", func(c echo.Context) error {
		discordInfo, err := discord.TokenLookup((c.Param("discord")))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, discordInfo)
	})

	e.GET("/domain/:domain", func(c echo.Context) error {
		domainInfo, err := ip2location.DomainLookup(os.Getenv("BITCROOK_IPTL"), c.Param("domain"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, domainInfo)
	})

	e.GET("/tin/:tin", func(c echo.Context) error {
		tinInfo, err := tin.TINLookup(c.Param("tin"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, tinInfo)
	})
	fmt.Println(color.Colorize(color.Blue, "[i]"), "HTTP server started on", color.Colorize(color.Green, "[::]:6174"))
	fmt.Println(color.Colorize(color.Blue, "[i]"), "Access Bitcrook via", color.Colorize(color.Green, "127.0.0.1:6174"))
	fmt.Println()

	e.Logger.Fatal(e.Start(":6174"))
}
