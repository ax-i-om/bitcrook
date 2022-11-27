$('#ipsearch').on('click', () => {
    ip = $('#ip').val()
    $.getJSON('http://localhost:6174/ip/' + ip, (res) => {
        let iptable = document.getElementById('ipresult');
        iptable.innerHTML = "";
        iptable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Status:&nbsp;</strong>" + res.status + "</span>";
        iptable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Continent:&nbsp;</strong>" + res.continent + "</span>";
        iptable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Country:&nbsp;</strong>" + res.country + "</span>";
        iptable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Region Name:&nbsp;</strong>" + res.regionname + "</span>";
        iptable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>City:&nbsp;</strong>" + res.city + "</span>";
        iptable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>District:&nbsp;</strong>" + res.district + "</span>";
        iptable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Zip:&nbsp;</strong>" + res.zip + "</span>";
        iptable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Latitude:&nbsp;</strong>" + res.lat + "</span>";
        iptable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Longitude:&nbsp;</strong>" + res.lon + "</span>";
        iptable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Time Zone:&nbsp;</strong>" + res.timezone + "</span>";
        iptable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Currency:&nbsp;</strong>" + res.currency + "</span>";
        iptable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>ISP:&nbsp;</strong>" + res.isp + "</span>";
        iptable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Org:&nbsp;</strong>" + res.org + "</span>";
        iptable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>As:&nbsp;</strong>" + res.as + "</span>";
        iptable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Asname:&nbsp;</strong>" + res.asname + "</span>";
        iptable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Reverse:&nbsp;</strong>" + res.reverse + "</span>";
        iptable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Mobile:&nbsp;</strong>" + res.mobile + "</span>";
        iptable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Proxy:&nbsp;</strong>" + res.proxy + "</span>";
        iptable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Hosting:&nbsp;</strong>" + res.hosting + "</span>";
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
        let domaintable = document.getElementById('domainresult'); // need API key to continue testing
        domaintable.innerHTML = "";
        domaintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Domain:&nbsp;</strong>" + res.domain + "</span>";
        domaintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Domain ID:&nbsp;</strong>" + res.domain_id + "</span>";
        domaintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Status:&nbsp;</strong>" + res.status + "</span>";
        domaintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Domain Age:&nbsp;</strong>" + res.domain_age + "</span>";
        domaintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>WHOIS Server:&nbsp;</strong>" + res.whois_server + "</span><br><br>";
        
        domaintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Registrar IANA ID:&nbsp;</strong>" + res.registrar.iana_id + "</span>";
        domaintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Registrar Name:&nbsp;</strong>" + res.registrar.name + "</span>";
        domaintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Registrar URL:&nbsp;</strong>" + res.registrar.url + "</span><br><br>";
        /*
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

        "Nameservers: " + res.nameservers
        */
    })
    $('#domainresult').val('') 
})