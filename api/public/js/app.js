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
        let usernametable = document.getElementById('usernameresult');
        let recentswap = false;
        usernametable.innerHTML = "";
        
        for(let i = 0; i < res.length; i++) {
            let obj = res[i];
            if (obj.Valid) {
                if (recentswap) {
                    usernametable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>" + obj.Title + ":&nbsp;</strong>" + obj.Domain + "</span>";
                    recentswap = false;
                } else {
                    usernametable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>" + obj.Title + ":&nbsp;</strong>" + obj.Domain + "</span>";
                    recentswap = true;
                }
            }
        }
    })
    $('#usernameresult').val('') 
})

$('#vinsearch').on('click', () => {
    vin = $('#vin').val()
    $.getJSON('http://localhost:6174/vin/' + vin, (res) => {
        let vintable = document.getElementById('vinresult');
        vintable.innerHTML = "";
        vintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>VIN:&nbsp;</strong>" + res.Vin + "</span>";
        vintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Make:&nbsp;</strong>" + res.Make + "</span>";
        vintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Model:&nbsp;</strong>" + res.Model + "</span>";
        vintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Year:&nbsp;</strong>" + res.Year + "</span>";
        vintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Trim:&nbsp;</strong>" + res.Trim + "</span>";
        vintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Body:&nbsp;</strong>" + res.Body + "</span>";
        vintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Engine:&nbsp;</strong>" + res.Engine + "</span>";
        vintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Manufactured In:&nbsp;</strong>" + res.ManufacturedIn + "</span>";
        vintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Trim Level:&nbsp;</strong>" + res.TrimLevel + "</span>";
        vintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Steering Type:&nbsp;</strong>" + res.SteeringType + "</span>";
        vintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>ABS:&nbsp;</strong>" + res.Abs + "</span>";
        vintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Tank Size:&nbsp;</strong>" + res.TankSize + "</span>";
        vintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Overall Height:&nbsp;</strong>" + res.OverallHeight + "</span>";
        vintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Overall Length:&nbsp;</strong>" + res.OverallLength + "</span>";
        vintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Overall Width:&nbsp;</strong>" + res.OverallWidth + "</span>";
        vintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Standard Seating:&nbsp;</strong>" + res.StandardSeating + "</span>";
        vintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Optional Seating:&nbsp;</strong>" + res.OptionalSeating + "</span>";
        vintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Highway Mileage:&nbsp;</strong>" + res.HighwayMileage + "</span>";
        vintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>City Mileage:&nbsp;</strong>" + res.CityMileage + "</span>";
        vintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Fuel Type:&nbsp;</strong>" + res.FuelType + "</span>";
        vintable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Drive Type:&nbsp;</strong>" + res.DriveType + "</span>";
        vintable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Transmission:&nbsp;</strong>" + res.Transmission + "</span>";
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

$('#discordsearch').on('click', () => {
    discord = $('#discord').val()
    $.getJSON('http://localhost:6174/discord/' + discord, (res) => {
        let discordtable = document.getElementById('discordresult');
        discordtable.innerHTML = "";
        discordtable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>ID:&nbsp;</strong>" + res.id + "</span>";
        discordtable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Username:&nbsp;</strong>" + res.username + "</span>";
        discordtable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Avatar:&nbsp;</strong>" + res.avatar + "</span>";
        discordtable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Discriminator:&nbsp;</strong>" + res.discriminator + "</span>";
        discordtable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Public Flags:&nbsp;</strong>" + res.public_flags + "</span>";
        discordtable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Flags:&nbsp;</strong>" + res.flags + "</span>";
        discordtable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Purchased Flags:&nbsp;</strong>" + res.purchased_flags + "</span>";
        discordtable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Locale:&nbsp;</strong>" + res.locale + "</span>";
        discordtable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>NSFW Allowed:&nbsp;</strong>" + res.nsfw_allowed + "</span>";
        discordtable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>MFA Enabled:&nbsp;</strong>" + res.mfa_enabled + "</span>";
        discordtable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Email:&nbsp;</strong>" + res.email + "</span>";
        discordtable.innerHTML += "<span style='display:flex; padding: 5px; justify-content: center'>" + "<strong>Verified:&nbsp;</strong>" + res.verified + "</span>";
        discordtable.innerHTML += "<span style='background-color: #353b48; display:flex; padding: 5px; justify-content: center'>" + "<strong>Phone:&nbsp;</strong>" + res.phone + "</span>";
     })
    $('#ipresult').val('')
})