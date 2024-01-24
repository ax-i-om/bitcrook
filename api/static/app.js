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

const selectList = document.querySelector("#querytype")
// get initial width of select element. 
// we have to remember there is a dropdown arrow make it a little wider
const initialWidth = selectList.offsetWidth
// get text content length (not a value length) of widest option. 
const maxOptValLen = findMaxLengthOpt(selectList)
// calc width of single letter 
const letterWidth = initialWidth / maxOptValLen
const corCoef = 4.875; // Based on visual appearance
// console.log(initialWidth, maxOptValLen)

function reSizeDrop(selist) {
  const newOptValLen = getSelected(selist).textContent.length
  const correction = (maxOptValLen - newOptValLen) * corCoef
  const newValueWidth = (newOptValLen * letterWidth) + correction
  // console.log('new width', newValueWidth, 'new option len', newOptValLen)
  selist.style.width = `${newValueWidth}px`;
}

selectList.addEventListener("change", e => {
  reSizeDrop(e.target)
}, false);


function getSelected(selectEl) {
  return [...selectEl.options].find(o => o.selected)
}

function findMaxLengthOpt(selectEl) {
  return [...selectEl.options].reduce((result, o) => o.textContent.length > result ? o.textContent.length : result, 0)
}

particlesJS("particles-js", {
    particles: {
        number: { value: 80, density: { enable: true, value_area: 800 } },
        color: { value: "#ffffff" },
        shape: {
            type: "circle",
            stroke: { width: 0, color: "#000000" },
            polygon: { nb_sides: 5 },
            image: { src: "img/github.svg", width: 100, height: 100 }
        },
        opacity: {
            value: 0.5,
            random: false,
            anim: { enable: false, speed: 1, opacity_min: 0.1, sync: false }
        },
        size: {
            value: 3,
            random: true,
            anim: { enable: false, speed: 40, size_min: 0.1, sync: false }
        },
        line_linked: {
            enable: true,
            distance: 150,
            color: "#ffffff",
            opacity: 0.4,
            width: 1
        },
        move: {
            enable: true,
            speed: 4,
            direction: "none",
            random: false,
            straight: false,
            out_mode: "out",
            bounce: false,
            attract: { enable: false, rotateX: 600, rotateY: 1200 }
        }
    },
    interactivity: {
        detect_on: "canvas",
        events: {
            onhover: { enable: false, mode: "repulse" },
            onclick: { enable: false, mode: "push" },
            resize: true
        },
        modes: {
            grab: { distance: 400, line_linked: { opacity: 1 } },
            bubble: { distance: 400, size: 40, duration: 2, opacity: 8, speed: 3 },
            repulse: { distance: 100, duration: 0.4 },
            push: { particles_nb: 4 },
            remove: { particles_nb: 2 }
        }
    },
    retina_detect: true
});

/**
 * Function to add a row of information with a light background to a results table
 * @param table The result table where the row will be added
 * @param key The key to be used
 * @param value The value to be used
 * @returns Nothing
 */
function addLight(table, key, value) {
    table.innerHTML += `<span style='background: rgba(75, 75, 75, 1);  display:flex; padding: 5px; justify-content: center'><strong>${key}:&nbsp;</strong>${value}</span>`;
}

/**
 * Function to add a row of information with a dark background to a results table
 * @param table The result table where the row will be added
 * @param key The key to be used
 * @param value The value to be used
 * @returns Nothing
 */
function addDark(table, key, value) {
    table.innerHTML += `<span style='background: rgba(75, 75, 75, .25); display:flex; padding: 5px; justify-content: center'><strong>${key}:&nbsp;</strong>${value}</span>`;
}

/**
 * Function to a gap to a results table
 * @param table The result table where the row will be added
 * @returns Nothing
 */
function addGap(table) {
    table.innerHTML += "<br><br>"
} 

function launch() {
    $('#queryloadercircle').addClass('loader');
    const query = $('#query').val()
    const resulttable = document.getElementById('queryresult');
    resulttable.innerHTML = "";
    if (query) {
        const type = $('#querytype').val()
        if (type === "discord") {
            $.getJSON(`/discord/${query}`, (res) => {
                $('#queryloadercircle').removeClass('loader');
                addLight(resulttable, "ID", res.id)
                addDark(resulttable, "Username", res.username)
                addLight(resulttable, "Avatar", res.avatar)
                addDark(resulttable, "Discriminator", res.discriminator)
                addLight(resulttable, "Public Flags", res.public_flags)
                addDark(resulttable, "Flags", res.flags)
                addLight(resulttable, "Banner", res.banner)
                addDark(resulttable, "Accent Color", res.accent_color)
                addLight(resulttable, "Global Name", res.global_name)
                addDark(resulttable, "Avatar Decoration", res.avatar_decoration_data)
                addLight(resulttable, "Banner Color", res.banner_color)
                addDark(resulttable, "MFA Enabled", res.mfa_enabled)
                addLight(resulttable, "Locale", res.locale)
                addDark(resulttable, "Premium Type", res.premium_type)
                addLight(resulttable, "Email", res.email)
                addDark(resulttable, "Verified", res.verified)
                addLight(resulttable, "Phone", res.phne)
                addDark(resulttable, "NSFW Allowed", res.nsfw_allowed)
                addLight(resulttable, "Linked Users", res.linked_users)
                addDark(resulttable, "Bought Flags", res.purchased_flags)
                addLight(resulttable, "Bio", res.bio)
                addDark(resulttable, "Auth Types", res.authenticator_types)
            })
        } else if (type === "domain") {
            $.getJSON(`/domain/${query}`, (res) => {
                $('#queryloadercircle').removeClass('loader');
                addLight(resulttable, "Domain", res.domain)
                addDark(resulttable, "Domain ID", res.domain_id)
                addLight(resulttable, "Status", res.status)
                addDark(resulttable, "Domain Age", res.domain_age)
                addLight(resulttable, "WHOIS Server", res.whois_server)
                
                addGap(resulttable)
        
                addLight(resulttable, "Registrar IANA ID", res.registrar.iana_id)
                addDark(resulttable, "Registrar Name", res.registrar.name)
                addLight(resulttable, "Registrar URL", res.registrar.url)
        
                if (res.registrant.name) {
                    addGap(resulttable)
                
                    addLight(resulttable, "Registrant Name", res.registrant.name)
                    addDark(resulttable, "Registrant Organization", res.registrant.organization)
                    addLight(resulttable, "Registrant Street Address", res.registrant.street_address)
                    addDark(resulttable, "Registrant City", res.registrant.city)
                    addLight(resulttable, "Registrant Region", res.registrant.region)
                    addDark(resulttable, "Registrant Zip Code", res.registrant.zip_code)
                    addLight(resulttable, "Registrant Country", res.registrant.country)
                    addDark(resulttable, "Registrant Phone", res.registrant.phone)
                    addLight(resulttable, "Registrant Fax", res.registrant.fax)
                    addDark(resulttable, "Registrant Email", res.registrant.email)
                }
        
                if (res.admin.name) {
                    addGap(resulttable)
        
                    addLight(resulttable, "Admin Name", res.admin.name)
                    addDark(resulttable, "Admin Organization", res.admin.organization)
                    addLight(resulttable, "Admin Street Address", res.admin.street_address)
                    addDark(resulttable, "Admin City", res.admin.city)
                    addLight(resulttable, "Admin Region", res.admin.region)
                    addDark(resulttable, "Admin Zip Code", res.admin.zip_code)
                    addLight(resulttable, "Admin Country", res.admin.country)
                    addDark(resulttable, "Admin Phone", res.admin.phone)
                    addLight(resulttable, "Admin Fax", res.admin.fax)
                    addDark(resulttable, "Admin Email", res.admin.email)
                }
        
                if (res.tech.name) {
                    addGap(resulttable)
        
                    addLight(resulttable, "Tech Name", res.tech.name)
                    addDark(resulttable, "Tech Organization", res.tech.organization)
                    addLight(resulttable, "Tech Street Address", res.tech.street_address)
                    addDark(resulttable, "Tech City", res.tech.city)
                    addLight(resulttable, "Tech Region", res.tech.region)
                    addDark(resulttable, "Tech Zip Code", res.tech.zip_code)
                    addLight(resulttable, "Tech Country", res.tech.country)
                    addDark(resulttable, "Tech Phone", res.tech.phone)
                    addLight(resulttable, "Tech Fax", res.tech.fax)
                    addDark(resulttable, "Tech Email", res.tech.email)
                }
        
                if (res.billing.name) {
                    addGap(resulttable)
        
                    addLight(resulttable, "Billing Name", res.billing.name)
                    addDark(resulttable, "Billing Organization", res.billing.organization)
                    addLight(resulttable, "Billing Street Address", res.billing.street_address)
                    addDark(resulttable, "Billing City", res.billing.city)
                    addLight(resulttable, "Billing Region", res.billing.region)
                    addDark(resulttable, "Billing Zip Code", res.billing.zip_code)
                    addLight(resulttable, "Billing Country", res.billing.country)
                    addDark(resulttable, "Billing Phone", res.billing.phone)
                    addLight(resulttable, "Billing Fax", res.billing.fax)
                    addDark(resulttable, "Billing Email", res.billing.email)
                }
        
                addGap(resulttable)
        
                addLight(resulttable, "Nameservers", res.nameservers)
            })
        } else if (type === "email") {
            $.getJSON(`/email/${query}`, (res) => {
                $('#queryloadercircle').removeClass('loader');
                addLight(resulttable, "Deliverability Confidence Score", res.Deliverabilityconfidencescore)
                addDark(resulttable, "Email Address", res.Emailaddress)
                addLight(resulttable, "Mailbox Name", res.Mailboxname)
                addDark(resulttable, "Domain Name", res.Domainname)
                addLight(resulttable, "Top Level Domain", res.Topleveldomain)
                addDark(resulttable, "TopLevel Domain Name", res.Topleveldomainname)
                addLight(resulttable, "Date Checked", res.Datechecked)
                addDark(resulttable, "Estimated Email Age", res.Emailageestimated)
                addLight(resulttable, "Estimated Domain Age", res.Domainageestimated)
                addDark(resulttable, "Domain Expiration Date", res.Domainexpirationdate)
                addLight(resulttable, "Domain Created Date", res.Domaincreateddate)
                addDark(resulttable, "Domain Updated Date", res.Domainupdateddate)
                addLight(resulttable, "Domain Email", res.Domainemail)
                addDark(resulttable, "Domain Organization", res.Domainorganization)
                addLight(resulttable, "Domain Address 1", res.Domainaddress1)
                addDark(resulttable, "Domain Locality", res.Domainlocality)
                addLight(resulttable, "Domain Administrative Area", res.Domainadministrativearea)
                addDark(resulttable, "Domain Postal Code", res.Domainpostalcode)
                addLight(resulttable, "Domain Country", res.Domaincountry)
                addDark(resulttable, "Domain Availability", res.Domainavailability)
                addLight(resulttable, "Domain Country Code", res.Domaincountrycode)
                addDark(resulttable, "Domain Private Proxy", res.Domainprivateproxy)
                addLight(resulttable, "Privacy Flag", res.Privacyflag)
                addDark(resulttable, "MX Server", res.Mxserver)
            })
        } else if (type === "ip") {
            $.getJSON(`/ip/${query}`, (res) => {
                $('#queryloadercircle').removeClass('loader');
                addLight(resulttable, "Status", res.status)
                addDark(resulttable, "Continent", res.continent)
                addLight(resulttable, "Country", res.country)
                addDark(resulttable, "Region Name", res.regionname)
                addLight(resulttable, "City", res.city)
                addDark(resulttable, "District", res.district)
                addLight(resulttable, "Zip", res.zip)
                addDark(resulttable, "Latitude", res.lat)
                addLight(resulttable, "Longitude", res.lon)
                addDark(resulttable, "Timezone", res.timezone)
                addLight(resulttable, "Currency", res.currency)
                addDark(resulttable, "ISP", res.isp)
                addLight(resulttable, "Org", res.org)
                addDark(resulttable, "As", res.as)
                addLight(resulttable, "Asname", res.asname)
                addDark(resulttable, "Reverse", res.reverse)
                addLight(resulttable, "Mobile", res.mobile)
                addDark(resulttable, "Proxy", res.proxy)
                addLight(resulttable, "Hosting", res.hosting)
            })
        } else if (type === "phone") {
            $.getJSON(`/phone/${query}`, (res) => {
                $('#queryloadercircle').removeClass('loader');
                addLight(resulttable, "Phone Number", res.Phonenumber)
                addDark(resulttable, "Administrative Area", res.Administrativearea)
                addLight(resulttable, "Country Abbreviation", res.Countryabbreviation)
                addDark(resulttable, "Country Name", res.Countryname)
                addLight(resulttable, "Carrier", res.Carrier)
                addDark(resulttable, "Caller ID", res.Callerid)
                addLight(resulttable, "DST", res.Dst)
                addDark(resulttable, "International Phone Number", res.Internationalphonenumber)
                addLight(resulttable, "Language", res.Language)
                addDark(resulttable, "Latitude", res.Latitude)
                addLight(resulttable, "Longitude", res.Longitude)
                addDark(resulttable, "Locality", res.Locality)
                addLight(resulttable, "Phone International Prefix", res.Phoneinternationalprefix)
                addDark(resulttable, "Phone Country Dialing Code", res.Phonecountrydialingcode)
                addLight(resulttable, "Phone Nation Prefix", res.Phonenationprefix)
                addDark(resulttable, "Phone National Destination Code", res.Phonenationaldestinationcode)
                addLight(resulttable, "Phone Subscriber Number", res.Phonesubscribernumber)
                addDark(resulttable, "UTC", res.Utc)
                addLight(resulttable, "Timezone Code", res.Timezonecode)
                addDark(resulttable, "Timezone Name", res.Timezonename)
                addLight(resulttable, "Postal Code", res.Postalcode)
            })
        } else if (type === "record") {
            $.getJSON(`/record/${query}`, (res) => {
                $('#queryloadercircle').removeClass('loader');
                for(let i = 0; i < res.results.length; i++) {
                    addLight(resulttable, "Name", res.results[i].name_abbreviation)
                    addDark(resulttable, "Decision Date", res.results[i].decision_date)
                    addLight(resulttable, "Jurisdiction", res.results[i].jurisdiction.name_long)
                    addDark(resulttable, "Court", res.results[i].court.name)
                    addLight(resulttable, "Source", res.results[i].provenance.source)
                    addDark(resulttable, "URL", res.results[i].url)

                    addGap(resulttable)
                }
            })
        } else if (type === "tin") {
            $.getJSON(`/tin/${query}`, (res) => {
                $('#queryloadercircle').removeClass('loader');
                for(let i = 0; i < res.rows.length; i++) {
                    const obj = res.rows[i];
                    addLight(resulttable, "Assignment Date", obj.r)
                    addDark(resulttable, "Termination Date", obj.e)
                    addLight(resulttable, "Page", obj.pg)
                    addDark(resulttable, "Total", obj.tot)
                    addLight(resulttable, "Count", obj.cnt)
                    addDark(resulttable, "INN", obj.i)
                    addLight(resulttable, "K", obj.k)
                    addDark(resulttable, "Name", obj.n)
                    addLight(resulttable, "OGRNIP", obj.o)
                    addGap(resulttable)
                }
            })
        } else if (type === "username") {
            $.getJSON(`/username/${query}`, (res) => {
                let recentswap = false;
                $('#queryloadercircle').removeClass('loader');
                for(let i = 0; i < res.length; i++) {
                    const obj = res[i];
                    if (obj.Valid) {
                        if (recentswap) {
                            resulttable.innerHTML += `<span style='background: rgba(75, 75, 75, .25); display:flex; padding: 5px; justify-content: center'><strong>${obj.Title}:&nbsp;</strong><a>${obj.Domain}</a></span>`;
                            $(function(){ $(`a:contains(${obj.Domain})`).attr("href", obj.Domain)});
                            $(function(){ $(`a:contains(${obj.Domain})`).attr("target", "_blank")});
                            recentswap = false;
                        } else {
                            resulttable.innerHTML += `<span style='background: rgba(75, 75, 75, 1); display:flex; padding: 5px; justify-content: center'><strong>${obj.Title}:&nbsp;</strong><a>${obj.Domain}</a></span>`;
                            $(function(){ $(`a:contains(${obj.Domain})`).attr("href", obj.Domain)});
                            $(function(){ $(`a:contains(${obj.Domain})`).attr("target", "_blank")});
                            recentswap = true;
                        }
                    }
                }
            })
        } else if (type === "vin") {
            $.getJSON(`/vin/${query}`, (res) => {
                $('#queryloadercircle').removeClass('loader');
                addLight(resulttable, "VIN", res.Vin)
                addDark(resulttable, "Make", res.Make)
                addLight(resulttable, "Model", res.Model)
                addDark(resulttable, "Year", res.Year)
                addLight(resulttable, "Trim", res.Trim)
                addDark(resulttable, "Body", res.Body)
                addLight(resulttable, "Engine", res.Engine)
                addDark(resulttable, "Manufactured In", res.ManufacturedIn)
                addLight(resulttable, "Trim Level", res.TrimLevel)
                addDark(resulttable, "Steering Type", res.SteeringType)
                addLight(resulttable, "ABS", res.Abs)
                addDark(resulttable, "Tank Size", res.TankSize)
                addLight(resulttable, "Overall Height", res.OverallHeight)
                addDark(resulttable, "Overall Length", res.OverallLength)
                addLight(resulttable, "Overall Width", res.OverallWidth)
                addDark(resulttable, "Standard Seating", res.StandardSeating)
                addLight(resulttable, "Optional Seating", res.OptionalSeating)
                addDark(resulttable, "Highway Mileage", res.HighwayMileage)
                addLight(resulttable, "City Mileage", res.CityMileage)
                addDark(resulttable, "Fuel Type", res.FuelType)
                addLight(resulttable, "Drive Type", res.DriveType)
                addDark(resulttable, "Transmission", res.Transmission)
            })
        } else {
            $('#queryloadercircle').removeClass('loader');
            addLight(resulttable, "Error", "Malformed query type")
        }
    } else {
        $('#queryloadercircle').removeClass('loader');
        addLight(resulttable, "Error", "Query was not specified")
    }
}

$('#query').keypress(function (e) {                                       
    if (e.which === 13) {
        e.preventDefault();
        launch() 
    }
});


$('#launchquery').on('click', () => {
    launch()
})

window.onload = reSizeDrop(selectList);