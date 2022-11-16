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

$('#domainsearch').on('click', () => {
    domain = $('#domain').val()
    $.getJSON('http://localhost:6174/domain/' + domain, (res) => {
        $('#domainresult').val("Domain: " + res.domain + "\n" +
        "Domain ID: " + res.domain_id + "\n" +
        "Status: " + res.status + "\n" + 
        "Domain Age: " + res.domain_age + "\n" + 
        "WHOIS Server: " + res.whois_server + "\n\n" +
        
        "Registrar IANA ID: " + res.registrar.iana_id + "\n" + 
        "Registrar Name: " + res.registrar.name + "\n" + 
        "Registrar URL: " + res.registrar.url + "\n\n" +

        "Registrant Name: " + res.registrant.name + "\n" +
        "Registrant Organization: " + res.registrant.organization + "\n" +
        "Registrant Street Address: " + res.registrant.street_address + "\n" +
        "Registrant City: " + res.registrant.city + "\n" +
        "Registrant Region: " + res.registrant.region + "\n" +
        "Registrant Zip Code: " + res.registrant.zip_code + "\n" +
        "Registrant Country: " + res.registrant.country + "\n" +
        "Registrant Phone: " + res.registrant.phone + "\n" +
        "Registrant Fax: " + res.registrant.fax + "\n" +
        "Registrant Email: " + res.registrant.email + "\n\n" +

        "Admin Name: " + res.admin.name + "\n" +
        "Admin Organization: " + res.admin.organization + "\n" +
        "Admin Street Address: " + res.admin.street_address + "\n" +
        "Admin City: " + res.admin.city + "\n" +
        "Admin Region: " + res.admin.region + "\n" +
        "Admin Zip Code: " + res.admin.zip_code + "\n" +
        "Admin Country: " + res.admin.country + "\n" +
        "Admin Phone: " + res.admin.phone + "\n" +
        "Admin Fax: " + res.admin.fax + "\n" +
        "Admin Email: " + res.admin.email + "\n\n" +

        "Tech Name: " + res.tech.name + "\n" +
        "Tech Organization: " + res.tech.organization + "\n" +
        "Tech Street Address: " + res.tech.street_address + "\n" +
        "Tech City: " + res.tech.city + "\n" +
        "Tech Region: " + res.tech.region + "\n" +
        "Tech Zip Code: " + res.tech.zip_code + "\n" +
        "Tech Country: " + res.tech.country + "\n" +
        "Tech Phone: " + res.tech.phone + "\n" +
        "Tech Fax: " + res.tech.fax + "\n" +
        "Tech Email: " + res.tech.email + "\n\n" +

        "Billing Name: " + res.billing.name + "\n" +
        "Billing Organization: " + res.billing.organization + "\n" +
        "Billing Street Address: " + res.billing.street_address + "\n" +
        "Billing City: " + res.billing.city + "\n" +
        "Billing Region: " + res.billing.region + "\n" +
        "Billing Zip Code: " + res.billing.zip_code + "\n" +
        "Billing Country: " + res.billing.country + "\n" +
        "Billing Phone: " + res.billing.phone + "\n" +
        "Billing Fax: " + res.billing.fax + "\n" +
        "Billing Email: " + res.billing.email + "\n\n" +

        "Nameservers: " + res.nameservers)
    })
    $('#domainresult').val('') 
})