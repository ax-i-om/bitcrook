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

/**
 * Function to add a row of information with a light background to a results table
 * @param table The result table where the row will be added
 * @param key The key to be used
 * @param value The value to be used
 * @returns Nothing
 */
function addLight(table, key, value) {
    table.innerHTML += `<span style='background-color: #4b4b4b; display:flex; padding: 5px; justify-content: center'><strong>${key}:&nbsp;</strong>${value}</span>`;
}

/**
 * Function to add a row of information with a dark background to a results table
 * @param table The result table where the row will be added
 * @param key The key to be used
 * @param value The value to be used
 * @returns Nothing
 */
function addDark(table, key, value) {
    table.innerHTML += `<span style='display:flex; padding: 5px; justify-content: center'><strong>${key}:&nbsp;</strong>${value}</span>`;
}

/**
 * Function to a gap to a results table
 * @param table The result table where the row will be added
 * @returns Nothing
 */
function addGap(table) {
    table.innerHTML += "<br><br>"
} 

$('#ipsearch').on('click', () => {
    const ip = $('#ip').val()
    if (ip) {
        $('#iploadercircle').addClass('loader');
        $.getJSON(`/ip/${ip}`, (res) => {
            const iptable = document.getElementById('ipresult');
            iptable.innerHTML = "";
            $('#iploadercircle').removeClass('loader');
            addLight(iptable, "Status", res.status)
            addDark(iptable, "Continent", res.continent)
            addLight(iptable, "Country", res.country)
            addDark(iptable, "Region Name", res.regionname)
            addLight(iptable, "City", res.city)
            addDark(iptable, "District", res.district)
            addLight(iptable, "Zip", res.zip)
            addDark(iptable, "Latitude", res.lat)
            addLight(iptable, "Longitude", res.lon)
            addDark(iptable, "Timezone", res.timezone)
            addLight(iptable, "Currency", res.currency)
            addDark(iptable, "ISP", res.isp)
            addLight(iptable, "Org", res.org)
            addDark(iptable, "As", res.as)
            addLight(iptable, "Asname", res.asname)
            addDark(iptable, "Reverse", res.reverse)
            addLight(iptable, "Mobile", res.mobile)
            addDark(iptable, "Proxy", res.proxy)
            addLight(iptable, "Hosting", res.hosting)
        })
        $('#ipresult').val('')
    } else {
        const iptable = document.getElementById('ipresult');
        iptable.innerHTML = "";
        addLight(iptable, "Error", "IP address field was left empty")
    }
})

$('#usernamesearch').on('click', () => {
    const username = $('#username').val()
    if (username) {
        $('#userloadercircle').addClass('loader');
        const usernametable = document.getElementById('usernameresult');
        usernametable.innerHTML = "";
        $.getJSON(`/username/${username}`, (res) => {
            let recentswap = false;
            $('#userloadercircle').removeClass('loader');
            for(let i = 0; i < res.length; i++) {
                const obj = res[i];
                if (obj.Valid) {
                    if (recentswap) {
                        usernametable.innerHTML += `<span style='display:flex; padding: 5px; justify-content: center'><strong>${obj.Title}:&nbsp;</strong><a>${obj.Domain}</a></span>`;
                        $(function(){ $(`a:contains(${obj.Domain})`).attr("href", obj.Domain)});
                        $(function(){ $(`a:contains(${obj.Domain})`).attr("target", "_blank")});
                        recentswap = false;
                    } else {
                        usernametable.innerHTML += `<span style='background-color: #4b4b4b; display:flex; padding: 5px; justify-content: center'><strong>${obj.Title}:&nbsp;</strong><a>${obj.Domain}</a></span>`;
                        $(function(){ $(`a:contains(${obj.Domain})`).attr("href", obj.Domain)});
                        $(function(){ $(`a:contains(${obj.Domain})`).attr("target", "_blank")});
                        recentswap = true;
                    }
                }
            }
        })
        $('#usernameresult').val('') 
    } else {
        const usernametable = document.getElementById('usernameresult');
        usernametable.innerHTML = "";
        addLight(usernametable, "Error", "Username field was left empty")
    }
})

$('#vinsearch').on('click', () => {
    $('#vinloadercircle').addClass('loader');
    const vin = $('#vin').val()
    $.getJSON(`/vin/${vin}`, (res) => {
        const vintable = document.getElementById('vinresult');
        vintable.innerHTML = "";
        $('#vinloadercircle').removeClass('loader');
        addLight(vintable, "VIN", res.Vin)
        addDark(vintable, "Make", res.Make)
        addLight(vintable, "Model", res.Model)
        addDark(vintable, "Year", res.Year)
        addLight(vintable, "Trim", res.Trim)
        addDark(vintable, "Body", res.Body)
        addLight(vintable, "Engine", res.Engine)
        addDark(vintable, "Manufactured In", res.ManufacturedIn)
        addLight(vintable, "Trim Level", res.TrimLevel)
        addDark(vintable, "Steering Type", res.SteeringType)
        addLight(vintable, "ABS", res.Abs)
        addDark(vintable, "Tank Size", res.TankSize)
        addLight(vintable, "Overall Height", res.OverallHeight)
        addDark(vintable, "Overall Length", res.OverallLength)
        addLight(vintable, "Overall Width", res.OverallWidth)
        addDark(vintable, "Standard Seating", res.StandardSeating)
        addLight(vintable, "Optional Seating", res.OptionalSeating)
        addDark(vintable, "Highway Mileage", res.HighwayMileage)
        addLight(vintable, "City Mileage", res.CityMileage)
        addDark(vintable, "Fuel Type", res.FuelType)
        addLight(vintable, "Drive Type", res.DriveType)
        addDark(vintable, "Transmission", res.Transmission)
    })
    $('#vinresult').val('') 
})

$('#domainsearch').on('click', () => {
    const domain = $('#domain').val()
    if (domain) {
        $('#domainloadercircle').addClass('loader');
        $.getJSON(`/domain/${domain}`, (res) => {
        $('#domainloadercircle').removeClass('loader');
        const domaintable = document.getElementById('domainresult'); 
        domaintable.innerHTML = "";
        addLight(domaintable, "Domain", res.domain)
        addDark(domaintable, "Domain ID", res.domain_id)
        addLight(domaintable, "Status", res.status)
        addDark(domaintable, "Domain Age", res.domain_age)
        addLight(domaintable, "WHOIS Server", res.whois_server)
        
        addGap(domaintable)

        addLight(domaintable, "Registrar IANA ID", res.registrar.iana_id)
        addDark(domaintable, "Registrar Name", res.registrar.name)
        addLight(domaintable, "Registrar URL", res.registrar.url)

        addGap(domaintable)
        
        addLight(domaintable, "Registrant Name", res.registrant.name)
        addDark(domaintable, "Registrant Organization", res.registrant.organization)
        addLight(domaintable, "Registrant Street Address", res.registrant.street_address)
        addDark(domaintable, "Registrant City", res.registrant.city)
        addLight(domaintable, "Registrant Region", res.registrant.region)
        addDark(domaintable, "Registrant Zip Code", res.registrant.zip_code)
        addLight(domaintable, "Registrant Country", res.registrant.country)
        addDark(domaintable, "Registrant Phone", res.registrant.phone)
        addLight(domaintable, "Registrant Fax", res.registrant.fax)
        addDark(domaintable, "Registrant Email", res.registrant.email)

        addGap(domaintable)

        addLight(domaintable, "Admin Name", res.admin.name)
        addDark(domaintable, "Admin Organization", res.admin.organization)
        addLight(domaintable, "Admin Street Address", res.admin.street_address)
        addDark(domaintable, "Admin City", res.admin.city)
        addLight(domaintable, "Admin Region", res.admin.region)
        addDark(domaintable, "Admin Zip Code", res.admin.zip_code)
        addLight(domaintable, "Admin Country", res.admin.country)
        addDark(domaintable, "Admin Phone", res.admin.phone)
        addLight(domaintable, "Admin Fax", res.admin.fax)
        addDark(domaintable, "Admin Email", res.admin.email)

        addGap(domaintable)

        addLight(domaintable, "Tech Name", res.tech.name)
        addDark(domaintable, "Tech Organization", res.tech.organization)
        addLight(domaintable, "Tech Street Address", res.tech.street_address)
        addDark(domaintable, "Tech City", res.tech.city)
        addLight(domaintable, "Tech Region", res.tech.region)
        addDark(domaintable, "Tech Zip Code", res.tech.zip_code)
        addLight(domaintable, "Tech Country", res.tech.country)
        addDark(domaintable, "Tech Phone", res.tech.phone)
        addLight(domaintable, "Tech Fax", res.tech.fax)
        addDark(domaintable, "Tech Email", res.tech.email)

        addGap(domaintable)

        addLight(domaintable, "Billing Name", res.billing.name)
        addDark(domaintable, "Billing Organization", res.billing.organization)
        addLight(domaintable, "Billing Street Address", res.billing.street_address)
        addDark(domaintable, "Billing City", res.billing.city)
        addLight(domaintable, "Billing Region", res.billing.region)
        addDark(domaintable, "Billing Zip Code", res.billing.zip_code)
        addLight(domaintable, "Billing Country", res.billing.country)
        addDark(domaintable, "Billing Phone", res.billing.phone)
        addLight(domaintable, "Billing Fax", res.billing.fax)
        addDark(domaintable, "Billing Email", res.billing.email)

        addGap(domaintable)

        addLight(domaintable, "Nameservers", res.nameservers)
    })
    $('#domainresult').val('') 
    } else {
        const domaintable = document.getElementById('domainresult');
        domaintable.innerHTML = "";
        addLight(domaintable, "Error", "Domain field was left empty")
    }
})

$('#discordsearch').on('click', () => {
    const discord = $('#discord').val()
    if (discord) {
        $('#discordloadercircle').addClass('loader');
        $.getJSON(`/discord/${discord}`, (res) => {
            const discordtable = document.getElementById('discordresult');
            discordtable.innerHTML = "";
            $('#discordloadercircle').removeClass('loader');
            addLight(discordtable, "ID", res.id)
            addDark(discordtable, "Username", res.username)
            addLight(discordtable, "Avatar", res.avatar)
            addDark(discordtable, "Discriminator", res.discriminator)
            addLight(discordtable, "Public Flags", res.public_flags)
            addDark(discordtable, "Flags", res.flags)
            addLight(discordtable, "Banner", res.banner)
            addDark(discordtable, "Accent Color", res.accent_color)
            addLight(discordtable, "Global Name", res.global_name)
            addDark(discordtable, "Avatar Decoration", res.avatar_decoration_data)
            addLight(discordtable, "Banner Color", res.banner_color)
            addDark(discordtable, "MFA Enabled", res.mfa_enabled)
            addLight(discordtable, "Locale", res.locale)
            addDark(discordtable, "Premium Type", res.premium_type)
            addLight(discordtable, "Email", res.email)
            addDark(discordtable, "Verified", res.verified)
            addLight(discordtable, "Phone", res.phne)
            addDark(discordtable, "NSFW Allowed", res.nsfw_allowed)
            addLight(discordtable, "Linked Users", res.linked_users)
            addDark(discordtable, "Bought Flags", res.purchased_flags)
            addLight(discordtable, "Bio", res.bio)
            addDark(discordtable, "Auth Types", res.authenticator_types)
        })
        $('#ipresult').val('')
    } else {
        const discordtable = document.getElementById('discordresult');
        discordtable.innerHTML = "";
        addLight(discordtable, "Error", "Discord token field was left empty")
    }
})

$('#tinsearch').on('click', () => {
    const tin = $('#tin').val()
    if (tin) {
        const tintable = document.getElementById('tinresult');
        tintable.innerHTML = "";
        $('#tinloadercircle').addClass('loader');
        $.getJSON(`/tin/${tin}`, (res) => {
            $('#tinloadercircle').removeClass('loader');
            for(let i = 0; i < res.rows.length; i++) {
                const obj = res.rows[i];
                addLight(tintable, "Assignment Date", obj.r)
                addDark(tintable, "Termination Date", obj.e)
                addLight(tintable, "Page", obj.pg)
                addDark(tintable, "Total", obj.tot)
                addLight(tintable, "Count", obj.cnt)
                addDark(tintable, "INN", obj.i)
                addLight(tintable, "K", obj.k)
                addDark(tintable, "Name", obj.n)
                addLight(tintable, "OGRNIP", obj.o)
                addGap(tintable)
            }
        })
        $('#tinresult').val('')
    } else {
        const tintable = document.getElementById('tinresult');
        tintable.innerHTML = "";
        addLight(tintable, "Error", "TIN field was left empty")
    }
})