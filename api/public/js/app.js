$('#ipsearch').on('click', () => {
    ip = $('#ip').val()
    $.getJSON('http://localhost:6174/ip/' + ip, (res) => {
        var obj = JSON.parse(res);
        $('#ipresult').val("Status: " + obj.status + "\n" +
        "Status: " + obj.status + "\n" +
        "Continent: " + obj.continent + "\n" +
        "Country: " + obj.country + "\n" +
        "Region Name: " + obj.regionname + "\n" +
        "City: " + obj.district + "\n" +
        "District: " + obj.zip + "\n" +
        "Zip: " + obj.lat + "\n" +
        "Latitude: " + obj.lon + "\n" +
        "Longitude: " + obj.timezone + "\n" +
        "Time Zone: " + obj.currency + "\n" +
        "Currency: " + obj.isp + "\n" +
        "ISP: " + obj.org + "\n" +
        "Org: " + obj.as + "\n" +
        "As: " + obj.asname + "\n" +
        "Asname: " + obj.reverse + "\n" +
        "Reverse: " + obj.mobile + "\n" +
        "Mobile: " + obj.proxy + "\n" +
        "Proxy: " + obj.hosting + "\n" +
        "Hosting: " + obj.query + "\n")
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