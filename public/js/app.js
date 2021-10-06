$('#ipsearch').on('click', () => {
    ip = $('#ip').val()
    $.getJSON('https://go-seek.herokuapp.com/ip/' + ip, (res) => {
        $('#ipresult').val("Status: " + res.status + "\n" +
        "Continent: " + res.continent + "\n" +
        "Country: " + res.country + "\n" +
        "Region Name: " + res.regionname + "\n" +
        "City: " + res.city + "\n" +
        "District: " + res.district + "\n" +
        "Zip: " + res.zip + "\n" +
        "Latitude: " + res.lat + "\n" +
        "Longitude: " + res.lon + "\n" +
        "Time Zone: " + res.timezone + "\n" +
        "Currency: " + res.currency + "\n" +
        "ISP: " + res.isp + "\n" +
        "Org: " + res.org + "\n" +
        "As: " + res.as + "\n" +
        "Asname: " + res.asname + "\n" +
        "Reverse: " + res.reverse + "\n" +
        "Mobile: " + res.mobile + "\n" +
        "Proxy: " + res.proxy + "\n" +
        "Hosting: " + res.hosting + "\n")
    })
    $('#ipresult').val('')
})

$('#usernamesearch').on('click', () => {
    username = $('#username').val()
    $.getJSON('https://go-seek.herokuapp.com/username/' + username, (res) => {
        $('#usernameresult').val(JSON.stringify(res, null, 4))
    })
    $('#usernameresult').val('') 
})

$('#vinsearch').on('click', () => {
    vin = $('#vin').val()
    $.getJSON('https://go-seek.herokuapp.com/vin/' + vin, (res) => {
        $('#vinresult').val("Vin: " + res.Vin + "\n" +
        "Year: " + res.Year + "\n" +
        "Make: " + res.Make + "\n" +
        "Model: " + res.Model + "\n" +
        "Trim: " + res.Trim + "\n" +
        "Short Trim: " + res.ShortTrim + "\n" +
        "Trim Variations: " + res.TrimVariations + "\n" +
        "Made In: " + res.MadeIn + "\n" +
        "Made In City: " + res.MadeInCity + "\n" +
        "Vehicle Style: " + res.VehicleStyle + "\n" +
        "Vehicle Type: " + res.VehicleType + "\n" +
        "Vehicle Size: " + res.VehicleSize + "\n" +
        "Vehicle Category: " + res.VehicleCategory + "\n" +
        "Doors: " + res.Doors + "\n" +
        "Fuel Type: " + res.FuelType + "\n" +
        "Fuel Capacity: " + res.FuelCapacity + "\n" +
        "City Mileage: " + res.CityMileage + "\n" +
        "Highway Mileage: " + res.HighwayMileage + "\n" +
        "Engine: " + res.Engine + "\n" +
        "Engine Size: " + res.EngineSize + "\n" +
        "Engine Cylinders: " + res.EngineCylinders + "\n" +
        "Transmission Type: " + res.TransmissionType + "\n" +
        "Transmission Gears: " + res.TransmissionGears + "\n" +
        "Driven Wheels: " + res.DrivenWheels + "\n" +
        "Anti-Brake System: " + res.AntiBrakeSystem + "\n" +
        "Steering Type: " + res.SteeringType + "\n" +
        "Curb Weight: " + res.CurbWeight + "\n" +
        "Gross Weight: " + res.GrossWeight + "\n" +
        "Overall Height: " + res.OverallHeight + "\n" +
        "Overall Length: " + res.OverallLength + "\n" +
        "Overall Width: " + res.OverallWidth + "\n" +
        "Standard Seating: " + res.StandardSeating + "\n" +
        "Optional Seating: " + res.OptionalSeating + "\n" +
        "Invoice Price: " + res.InvoicePrice + "\n" +
        "Delivery Charges: " + res.DeliveryCharges + "\n" +
        "MSRP: " + res.MSRP)
    })
    $('#vinresult').val('') 
})