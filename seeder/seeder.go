package main

import (
	_ "learn-echo/autoload"
	"learn-echo/seeder/seeders"
)

func main() {
	seeders.SeedPermission()
}
