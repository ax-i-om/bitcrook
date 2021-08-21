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

$('#discordsearch').on('click', () => {
    token = $('#token').val()
    $.getJSON('http://localhost:6174/discord/' + token, (res) => {
        $('#discordresult').val(JSON.stringify(res, null, 4))
    })
    $('#discordresult').val('') 
})

$('#vinsearch').on('click', () => {
    vin = $('#vin').val()
    $.getJSON('http://localhost:6174/vin/' + vin, (res) => {
        $('#vinresult').val(JSON.stringify(res, null, 4))
    })
    $('#vinresult').val('') 
})