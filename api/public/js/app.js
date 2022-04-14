$('#ipsearch').on('click', () => {
    ip = $('#ip').val()
    $.getJSON('http://localhost:6174/ip/' + ip, (res) => {
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
    $.getJSON('http://localhost:6174/username/' + username, (res) => {
        $('#usernameresult').val(JSON.stringify(res, null, 4))
    })
    $('#usernameresult').val('') 
})

$('#vinsearch').on('click', () => {
    vin = $('#vin').val()
    $.getJSON('http://localhost:6174/vin/' + vin, (res) => {
        $('#vinresult').val("Vin: " + res.Vin + "\n" +
        "Make: " + res.Make + "\n" +
        "Model: " + res.Model + "\n" +
        "Year: " + res.Year + "\n" +
        "Trim: " + res.Trim + "\n" +
        "Body: " + res.Body + "\n" +
        "Engine: " + res.Engine + "\n" +
        "Manufactured In: " + res.ManufacturedIn + "\n" +
        "Trim Level: " + res.TrimLevel + "\n" +
        "Steering Type: " + res.SteeringType + "\n" +
        "ABS: " + res.Abs + "\n" +
        "Tank Size: " + res.TankSize + "\n" +
        "Overall Height: " + res.OverallHeight + "\n" +
        "Overall Length: " + res.OverallLength + "\n" +
        "Overall Width: " + res.OverallWidth + "\n" +
        "Standard Seating: " + res.StandardSeating + "\n" +
        "Optional Seating: " + res.OptionalSeating + "\n" +
        "Highway Mileage: " + res.HighwayMileage + "\n" +
        "City Mileage: " + res.CityMileage + "\n" +
        "Fuel Type: " + res.FuelType + "\n" +
        "Drive Type: " + res.DriveType + "\n" +
        "Transmission: " + res.Transmission)
    })
    $('#vinresult').val('') 
})