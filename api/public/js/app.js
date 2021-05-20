$('#ipsearch').on('click', () => {
    ip = $('#ip').val()
    $.getJSON('http://localhost:6174/ip/' + ip, (res) => {
        $('#ipresult').val(JSON.stringify(res, null, 4))
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

$('#backtohome').on('click', () => {
    location.href = 'http://localhost:6174'
})