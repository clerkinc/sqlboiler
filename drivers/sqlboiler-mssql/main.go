package main

import (
	"github.com/clerkinc/sqlboiler/v4/drivers"
	"github.com/clerkinc/sqlboiler/v4/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
